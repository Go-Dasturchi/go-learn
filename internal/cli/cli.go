// Package cli implements the go-learn terminal interface: the main
// menu, the per-lesson theory/example/task/Vim/check/hint loop, and
// the progress/list/reset subcommands.
package cli

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"go-learn/internal/checker"
	"go-learn/internal/lesson"
	"go-learn/internal/progress"
	"go-learn/internal/render"
	"go-learn/internal/runner"
)

var stdin = bufio.NewReader(os.Stdin)

// Run dispatches a CLI invocation to the right subcommand.
func Run(args []string) error {
	if len(args) == 0 {
		return runMenu()
	}
	switch args[0] {
	case "start":
		return runMenu()
	case "list":
		return cmdList()
	case "progress":
		return cmdProgress()
	case "reset":
		return cmdReset()
	case "lesson":
		if len(args) < 2 {
			return fmt.Errorf("usage: go-learn lesson <id>")
		}
		return cmdLesson(args[1])
	case "check":
		return cmdCheck()
	case "next":
		return cmdNext()
	case "help", "-h", "--help":
		printHelp()
		return nil
	default:
		printHelp()
		return fmt.Errorf("unknown command: %s", args[0])
	}
}

func printHelp() {
	fmt.Print(`go-learn — GO ZERO TO HERO interactive terminal course

Usage:
  go-learn              Boshlang'ich menyuni ochadi
  go-learn start        Boshlang'ich menyuni ochadi
  go-learn list         Barcha darslarni ro'yxatini ko'rsatadi
  go-learn progress     Progress dashboard'ni ko'rsatadi
  go-learn reset        Progress'ni tozalaydi
  go-learn lesson <id>  Ma'lum bir darsni to'g'ridan-to'g'ri ochadi
  go-learn check        Joriy papkadagi (dars workspace) main.go'ni tekshiradi —
                        Vim ichidagi terminal split'dan (masalan Ctrl+T) chiqmasdan
                        ishlatish uchun
  go-learn next         Birinchi tugallanmagan darsni to'g'ridan-to'g'ri ochadi
  go-learn help         Shu yordamni ko'rsatadi
`)
}

// readLine prints prompt and reads one line from stdin. If stdin is
// closed (EOF — e.g. piped input ran out, or the user pressed Ctrl-D),
// it exits the program instead of returning an empty string forever,
// which would otherwise spin the retry loop in runLessonFlow forever.
func readLine(prompt string) string {
	fmt.Print(prompt)
	line, err := stdin.ReadString('\n')
	if err != nil {
		fmt.Println()
		os.Exit(0)
	}
	return strings.TrimSpace(line)
}

func progressBar(pct float64, width int) string {
	return render.ProgressBar(pct, width)
}

func overallPercent(metas []lesson.Meta, prog *progress.Progress) float64 {
	if len(metas) == 0 {
		return 0
	}
	done := 0
	for _, m := range metas {
		if prog.IsCompleted(m.ID) {
			done++
		}
	}
	return float64(done) / float64(len(metas)) * 100
}

// categories returns the distinct categories, each with its lessons in
// manifest order. Display order puts "Problems" first (it's the
// largest, most-practiced category) regardless of where it happens to
// sit in the manifest — the manifest's own order (Fundamentals, Data
// Structures, Problems) is left untouched, since that's what course
// progression (e.g. `go-learn next`'s fallback, overall percent) walks.
func categories(metas []lesson.Meta) ([]string, map[string][]lesson.Meta) {
	var order []string
	grouped := make(map[string][]lesson.Meta)
	for _, m := range metas {
		if _, ok := grouped[m.Category]; !ok {
			order = append(order, m.Category)
		}
		grouped[m.Category] = append(grouped[m.Category], m)
	}
	return putFirst(order, "Problems"), grouped
}

// putFirst moves name to the front of order (if present), preserving
// the relative order of everything else.
func putFirst(order []string, name string) []string {
	for i, v := range order {
		if v == name {
			reordered := make([]string, 0, len(order))
			reordered = append(reordered, v)
			reordered = append(reordered, order[:i]...)
			reordered = append(reordered, order[i+1:]...)
			return reordered
		}
	}
	return order
}

func banner() {
	fmt.Println(render.Banner())
}

func runMenu() error {
	metas, err := lesson.LoadManifest()
	if err != nil {
		return err
	}
	prog, err := progress.Load()
	if err != nil {
		return err
	}
	order, grouped := categories(metas)

	for {
		fmt.Println()
		banner()
		fmt.Println()
		fmt.Printf("Progress: %s %.0f%%\n\n", progressBar(overallPercent(metas, prog), 15), overallPercent(metas, prog))
		for i, cat := range order {
			fmt.Printf("%d. %s\n", i, cat)
		}
		fmt.Println()
		choice := readLine("Choose ([Q] chiqish): ")
		if strings.EqualFold(choice, "q") || strings.EqualFold(choice, "quit") {
			return nil
		}
		idx, err := strconv.Atoi(choice)
		if err != nil || idx < 0 || idx >= len(order) {
			fmt.Println("Noto'g'ri tanlov, qaytadan urinib ko'ring.")
			continue
		}
		if err := runCategoryMenu(order[idx], grouped[order[idx]], prog); err != nil {
			return err
		}
	}
}

// runCategoryMenu drills from a chosen top-level category down through
// any Group/Subgroup levels present in its lessons (e.g. Problems ->
// Abramyan -> Begin/Integer/Array/...) before showing the final lesson
// list. Ordinary flat categories (Group/Subgroup empty on every lesson)
// skip straight to the lesson list, unchanged from before.
func runCategoryMenu(category string, metas []lesson.Meta, prog *progress.Progress) error {
	return runDrillLevel(category, metas, []func(lesson.Meta) string{
		func(m lesson.Meta) string { return m.Group },
		func(m lesson.Meta) string { return m.Subgroup },
	}, prog)
}

func runDrillLevel(title string, metas []lesson.Meta, keyFuncs []func(lesson.Meta) string, prog *progress.Progress) error {
	if len(keyFuncs) == 0 || allEmptyKey(metas, keyFuncs[0]) {
		return runLessonList(title, metas, prog)
	}
	order, grouped := groupByKey(metas, keyFuncs[0])
	for {
		fmt.Println()
		fmt.Println(title)
		fmt.Println()
		for i, key := range order {
			done, total := progressCount(grouped[key], prog)
			fmt.Printf("%d. %s (%d/%d)\n", i+1, key, done, total)
		}
		fmt.Println()
		choice := readLine("Choose ([B] orqaga, [Q] chiqish): ")
		switch strings.ToLower(choice) {
		case "b", "back":
			return nil
		case "q", "quit":
			os.Exit(0)
		}
		idx, err := strconv.Atoi(choice)
		if err != nil || idx < 1 || idx > len(order) {
			fmt.Println("Noto'g'ri tanlov, qaytadan urinib ko'ring.")
			continue
		}
		if err := runDrillLevel(order[idx-1], grouped[order[idx-1]], keyFuncs[1:], prog); err != nil {
			return err
		}
	}
}

func runLessonList(title string, metas []lesson.Meta, prog *progress.Progress) error {
	width := render.TerminalWidth()
	for {
		fmt.Println()
		fmt.Println(title)
		fmt.Println()
		for i, m := range metas {
			prefixPlain := fmt.Sprintf("⬜ %d. ", i+1)
			prefix := fmt.Sprintf("%s %d. ", render.Check(prog.IsCompleted(m.ID)), i+1)
			body := fmt.Sprintf("%s — %s", m.ID, m.Title)
			wrapped := render.WrapText(body, width-len(prefixPlain), strings.Repeat(" ", len(prefixPlain)))
			fmt.Println(prefix + wrapped)
		}
		fmt.Println()
		choice := readLine("Choose ([B] orqaga, [Q] chiqish): ")
		switch strings.ToLower(choice) {
		case "b", "back":
			return nil
		case "q", "quit":
			os.Exit(0)
		}
		idx, err := strconv.Atoi(choice)
		if err != nil || idx < 1 || idx > len(metas) {
			fmt.Println("Noto'g'ri tanlov, qaytadan urinib ko'ring.")
			continue
		}
		if err := runLessonFlow(metas[idx-1].ID, prog); err != nil {
			return err
		}
	}
}

func allEmptyKey(metas []lesson.Meta, keyFunc func(lesson.Meta) string) bool {
	for _, m := range metas {
		if keyFunc(m) != "" {
			return false
		}
	}
	return true
}

func groupByKey(metas []lesson.Meta, keyFunc func(lesson.Meta) string) ([]string, map[string][]lesson.Meta) {
	var order []string
	grouped := make(map[string][]lesson.Meta)
	for _, m := range metas {
		k := keyFunc(m)
		if _, ok := grouped[k]; !ok {
			order = append(order, k)
		}
		grouped[k] = append(grouped[k], m)
	}
	return order, grouped
}

func progressCount(metas []lesson.Meta, prog *progress.Progress) (done, total int) {
	total = len(metas)
	for _, m := range metas {
		if prog.IsCompleted(m.ID) {
			done++
		}
	}
	return
}

func cmdLesson(id string) error {
	prog, err := progress.Load()
	if err != nil {
		return err
	}
	resolved, err := resolveID(id)
	if err != nil {
		return err
	}
	return runLessonFlow(resolved, prog)
}

// resolveID lets the user type either the full id ("01-hello") or just
// its numeric prefix ("01" or "1") from `go-learn lesson <id>`.
func resolveID(id string) (string, error) {
	metas, err := lesson.LoadManifest()
	if err != nil {
		return "", err
	}
	for _, m := range metas {
		if m.ID == id {
			return m.ID, nil
		}
	}
	padded := id
	if n, err := strconv.Atoi(id); err == nil {
		padded = fmt.Sprintf("%02d", n)
	}
	for _, m := range metas {
		if strings.HasPrefix(m.ID, padded+"-") {
			return m.ID, nil
		}
	}
	return "", fmt.Errorf("lesson %q topilmadi", id)
}

// runCheckAndReport compiles and tests code against the lesson's
// hidden tests, printing the same pass/fail report used by the normal
// lesson flow. On success it marks the lesson completed. Shared by
// runLessonFlow and the standalone `check` command so both paths give
// identical feedback.
func runCheckAndReport(l *lesson.Lesson, code string, prog *progress.Progress) (bool, error) {
	fmt.Println()
	fmt.Println("Checking your solution...")
	fmt.Println()
	result, err := checker.Run(code, l.Test)
	if err != nil {
		return false, err
	}

	if !result.Compiled {
		fmt.Println(render.Fail("Code does not compile"))
		fmt.Println()
		fmt.Println(render.Dim(result.CompileOutput))
		return false, nil
	}

	fmt.Println(render.Success("Code compiles"))
	allPassed := true
	for _, t := range result.Tests {
		if t.Passed {
			fmt.Println(render.Success(t.Name + " passed"))
		} else {
			fmt.Println(render.Fail(t.Name + " failed"))
			allPassed = false
		}
	}
	if len(result.Tests) == 0 {
		allPassed = false
		fmt.Println(render.Fail("Hech qanday test topilmadi"))
	}

	if !allPassed {
		fmt.Println()
		fmt.Println(render.Fail("Tests failed"))
		fmt.Println()
		fmt.Println(render.Dim(result.RawOutput))
		return false, nil
	}

	fmt.Println()
	fmt.Println(render.Boxf("🎉 LESSON COMPLETED!"))
	if err := prog.MarkCompleted(l.ID); err != nil {
		return false, err
	}
	return true, nil
}

// cmdCheck lets the learner check their code from inside a terminal
// split opened within Vim/Neovim (e.g. via Ctrl+T), without leaving the
// editor at all. It infers the lesson id from the current directory —
// which is exactly ~/.go-learn/workspace/<lesson-id> when run from
// there — reads main.go from it, and reports pass/fail the same way
// the full lesson flow does.
func cmdCheck() error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}
	id := filepath.Base(cwd)
	resolved, err := resolveID(id)
	if err != nil {
		return fmt.Errorf("joriy papka (%s) biror darsga mos kelmadi — bu buyruqni dars workspace papkasidan (masalan ~/.go-learn/workspace/<dars-id>) ishga tushiring", id)
	}
	l, err := lesson.Load(resolved)
	if err != nil {
		return err
	}
	code, err := runner.ReadWorkspaceFile(filepath.Join(cwd, "main.go"))
	if err != nil {
		return fmt.Errorf("main.go o'qib bo'lmadi: %w", err)
	}
	prog, err := progress.Load()
	if err != nil {
		return err
	}
	passed, err := runCheckAndReport(l, code, prog)
	if err != nil {
		return err
	}
	if !passed {
		os.Exit(1)
	}

	metas, err := lesson.LoadManifest()
	if err != nil {
		return err
	}
	if next := nextLessonAfter(metas, l.ID); next != nil {
		fmt.Println()
		fmt.Printf("Next lesson:\n%s — %s\n\n", next.ID, next.Title)
		fmt.Println(render.Dim("Shu terminaldan to'g'ridan-to'g'ri o'tish uchun: go-learn next"))
	} else {
		fmt.Println()
		fmt.Println("Tabriklaymiz! Hozircha barcha darslarni tugatdingiz.")
	}
	return nil
}

// cmdNext jumps straight into the next lesson — a shortcut so the
// learner can move on from a terminal split (e.g. after `go-learn
// check` passes) without returning to the main menu.
//
// When run from inside a lesson's workspace directory (exactly what a
// Ctrl+T split's cwd is right after `go-learn check`), it continues
// sequentially from THAT lesson via nextLessonAfter — the same lesson
// `go-learn check` itself already names as "Next lesson". Scanning for
// the first incomplete lesson in the whole manifest instead (the old
// behavior) could jump backwards to an unrelated, unfinished topic
// from earlier in the course (e.g. back to a Fundamentals lesson after
// finishing a Problems/Leetcode problem) rather than actually moving
// forward from where the learner just was.
//
// Outside a lesson workspace (cwd doesn't resolve to a lesson id — the
// command run standalone, not from a Ctrl+T split), it falls back to
// resuming the course at the first not-yet-completed lesson.
func cmdNext() error {
	metas, err := lesson.LoadManifest()
	if err != nil {
		return err
	}
	prog, err := progress.Load()
	if err != nil {
		return err
	}

	if cwd, err := os.Getwd(); err == nil {
		if id, err := resolveID(filepath.Base(cwd)); err == nil {
			if next := nextLessonAfter(metas, id); next != nil {
				return runLessonFlow(next.ID, prog)
			}
			fmt.Println()
			fmt.Println("Tabriklaymiz! Hozircha barcha darslarni tugatdingiz.")
			return nil
		}
	}

	for _, m := range metas {
		if !prog.IsCompleted(m.ID) {
			return runLessonFlow(m.ID, prog)
		}
	}
	fmt.Println()
	fmt.Println("Tabriklaymiz! Hozircha barcha darslarni tugatdingiz.")
	return nil
}

func runLessonFlow(id string, prog *progress.Progress) error {
	l, err := lesson.Load(id)
	if err != nil {
		return err
	}
	metas, err := lesson.LoadManifest()
	if err != nil {
		return err
	}

	if runner.InNestedNvimTerminal() {
		// We're inside a :terminal split that a parent Neovim instance
		// opened (e.g. via Ctrl+T). The full theory/example/task text
		// would just be noise here — it's shown in the dedicated content
		// pane this updates in the parent instance — so keep this
		// terminal's own output to a short status line instead of
		// re-printing everything into it too.
		path, err := runner.EnsureWorkspaceFile(l.ID, l.Starter)
		if err != nil {
			return err
		}
		contentPath, err := runner.WriteLessonContent(l.ID, l.Markdown())
		if err != nil {
			return err
		}
		if err := runner.OpenEditorInPlace(contentPath, path); err != nil {
			return err
		}
		fmt.Println()
		fmt.Printf("Lesson %s — %s ochildi (asosiy Vim oynangizdagi dars matni/kod oynalarida).\n", l.ID, l.Title)
		fmt.Println(render.Dim("Kodni yozib bo'lgach, shu terminaldan `go-learn check` buyrug'ini ishlating."))
		return nil
	}

	fmt.Println()
	fmt.Printf("Lesson %s — %s\n\n", l.ID, l.Title)
	status := "0%"
	if prog.IsCompleted(l.ID) {
		status = "100%"
	}
	fmt.Printf("Progress: %s\n\n", status)
	fmt.Println("[ENTER] Start   [B] Back   [Q] Quit")
	choice := readLine("> ")
	switch strings.ToLower(choice) {
	case "b", "back":
		return nil
	case "q", "quit":
		os.Exit(0)
	}

	fmt.Println()
	fmt.Println(render.SectionHeader("THEORY"))
	fmt.Println(render.Markdown(l.Theory))

	if l.Example != "" {
		fmt.Println()
		fmt.Println(render.SectionHeader("EXAMPLE"))
		fmt.Println(render.GoCode(l.Example))
	}

	fmt.Println()
	fmt.Println(render.SectionHeader("TOPSHIRIQ"))
	fmt.Println(render.Markdown(l.Task))
	fmt.Println()

	fmt.Println(render.Dim("(Vim ikkita oynada ochiladi: dars matni + kod. Tugatgach `:wqa` yoki `:xa` bilan chiqing — oddiy `:wq`/`:q` faqat bitta oynani yopadi.)"))
	fmt.Println()

	attempts := 0
	for {
		readLine("Vim editorini ochish uchun ENTER bosing. > ")

		path, err := runner.EnsureWorkspaceFile(l.ID, l.Starter)
		if err != nil {
			return err
		}
		contentPath, err := runner.WriteLessonContent(l.ID, l.Markdown())
		if err != nil {
			return err
		}
		if err := runner.OpenEditor(contentPath, path); err != nil {
			return err
		}
		code, err := runner.ReadWorkspaceFile(path)
		if err != nil {
			return err
		}

		passed, err := runCheckAndReport(l, code, prog)
		if err != nil {
			return err
		}
		if !passed {
			attempts++
			if attempts >= 3 {
				if !showHintMenu(l) {
					return nil
				}
			}
			if !strings.EqualFold(readLine("Try again? [Y/n] "), "n") {
				continue
			}
			return nil
		}

		next := nextLessonAfter(metas, l.ID)
		if next == nil {
			fmt.Println()
			fmt.Println("Tabriklaymiz! Hozircha barcha darslarni tugatdingiz.")
			return nil
		}
		fmt.Println()
		fmt.Printf("Next lesson:\n%s — %s\n\n", next.ID, next.Title)
		cont := readLine("[ENTER] Continue   [Q] Quit menyuga qaytish uchun > ")
		if strings.EqualFold(cont, "q") {
			return nil
		}
		return runLessonFlow(next.ID, prog)
	}
}

func nextLessonAfter(metas []lesson.Meta, id string) *lesson.Meta {
	for i, m := range metas {
		if m.ID == id && i+1 < len(metas) {
			return &metas[i+1]
		}
	}
	return nil
}

// showHintMenu is shown after 3 failed attempts. It returns false if
// the learner chose to quit back to the menu instead of continuing.
func showHintMenu(l *lesson.Lesson) bool {
	fmt.Println()
	fmt.Println("Need a hint?")
	fmt.Println()
	fmt.Println("1. Small hint")
	fmt.Println("2. Strong hint")
	fmt.Println("3. Show solution")
	choice := readLine("> ")
	fmt.Println()
	switch choice {
	case "1":
		if len(l.Hints) > 0 {
			fmt.Println(l.Hints[0])
		} else {
			fmt.Println("Bu dars uchun hint mavjud emas.")
		}
	case "2":
		if len(l.Hints) > 1 {
			fmt.Println(l.Hints[1])
		} else {
			fmt.Println("Bu dars uchun kuchliroq hint mavjud emas.")
		}
	case "3":
		fmt.Println(render.Boxf("SOLUTION:"))
		fmt.Println(render.GoCode(l.Solution))
	}
	return true
}

func cmdList() error {
	metas, err := lesson.LoadManifest()
	if err != nil {
		return err
	}
	prog, err := progress.Load()
	if err != nil {
		return err
	}
	order, grouped := categories(metas)
	width := render.TerminalWidth()
	for _, cat := range order {
		fmt.Println()
		fmt.Println(cat)
		printLessonListIndented(grouped[cat], prog, 1, width)
	}
	return nil
}

// printLessonListIndented recursively prints lessons, indenting one
// extra level for each of Group/Subgroup present, so large nested
// collections (e.g. Problems -> Abramyan -> Begin/Integer/...) don't
// dump as one flat unstructured list.
func printLessonListIndented(metas []lesson.Meta, prog *progress.Progress, depth int, width int) {
	indent := strings.Repeat("  ", depth)
	keyFuncs := []func(lesson.Meta) string{
		func(m lesson.Meta) string { return m.Group },
		func(m lesson.Meta) string { return m.Subgroup },
	}
	if depth-1 < len(keyFuncs) && !allEmptyKey(metas, keyFuncs[depth-1]) {
		order, grouped := groupByKey(metas, keyFuncs[depth-1])
		for _, key := range order {
			fmt.Printf("%s%s\n", indent, key)
			printLessonListIndented(grouped[key], prog, depth+1, width)
		}
		return
	}
	for _, m := range metas {
		prefixPlain := fmt.Sprintf("%s⬜ ", indent)
		prefix := fmt.Sprintf("%s%s ", indent, render.Check(prog.IsCompleted(m.ID)))
		body := fmt.Sprintf("%s — %s", m.ID, m.Title)
		wrapped := render.WrapText(body, width-len(prefixPlain), strings.Repeat(" ", len(prefixPlain)))
		fmt.Println(prefix + wrapped)
	}
}

func cmdProgress() error {
	metas, err := lesson.LoadManifest()
	if err != nil {
		return err
	}
	prog, err := progress.Load()
	if err != nil {
		return err
	}
	order, grouped := categories(metas)

	fmt.Println()
	fmt.Println("GO ZERO TO HERO")
	fmt.Println()
	for _, cat := range order {
		list := grouped[cat]
		done := 0
		for _, m := range list {
			if prog.IsCompleted(m.ID) {
				done++
			}
		}
		fmt.Printf("%-18s %d/%d\n", cat, done, len(list))
	}
	fmt.Println()
	fmt.Printf("Overall progress: %.0f%%\n", overallPercent(metas, prog))

	for _, m := range metas {
		if !prog.IsCompleted(m.ID) {
			fmt.Println()
			fmt.Printf("Current:\n%s — %s\n", m.ID, m.Title)
			break
		}
	}
	return nil
}

func cmdReset() error {
	prog, err := progress.Load()
	if err != nil {
		return err
	}
	answer := readLine("Rostdan ham progress'ni tozalamoqchimisiz? [y/N] ")
	if !strings.EqualFold(answer, "y") {
		fmt.Println("Bekor qilindi.")
		return nil
	}
	if err := prog.Reset(); err != nil {
		return err
	}
	fmt.Println("Progress tozalandi.")
	return nil
}
