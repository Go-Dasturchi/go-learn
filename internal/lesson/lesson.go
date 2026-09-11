// Package lesson loads lesson content (theory, example, task, hints,
// starter/solution/test code) from the embedded lessons/ directory.
package lesson

import (
	"bufio"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	golearn "go-learn"
)

// Meta is one entry of lessons/manifest.json — used to build the
// ordered lesson list and the level dashboard without having to parse
// every lesson.md just to show a menu.
type Meta struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Category string `json:"category"`
	// Group and Subgroup add up to two extra menu levels below Category
	// (Category -> Group -> Subgroup -> lessons), for large practice
	// collections like Problems -> Abramyan -> Begin/Integer/Array/...
	// Left empty for regular course lessons, which stay a flat
	// Category -> lessons menu.
	Group    string `json:"group,omitempty"`
	Subgroup string `json:"subgroup,omitempty"`
}

// Lesson is the fully parsed content of a single lesson.
type Lesson struct {
	ID       string
	Title    string
	Category string
	Theory   string
	Example  string
	Task     string
	Hints    []string
	Starter  string
	Solution string
	Test     string
}

// LoadManifest reads lessons/manifest.json, which lists every lesson in
// course order together with the dashboard category it belongs to.
func LoadManifest() ([]Meta, error) {
	data, err := golearn.LessonsFS.ReadFile("lessons/manifest.json")
	if err != nil {
		return nil, fmt.Errorf("manifest: %w", err)
	}
	var metas []Meta
	if err := json.Unmarshal(data, &metas); err != nil {
		return nil, fmt.Errorf("manifest: %w", err)
	}
	return metas, nil
}

// Load reads and parses a single lesson by its directory id (e.g. "01-hello").
func Load(id string) (*Lesson, error) {
	metas, err := LoadManifest()
	if err != nil {
		return nil, err
	}
	var meta *Meta
	for i := range metas {
		if metas[i].ID == id {
			meta = &metas[i]
			break
		}
	}
	if meta == nil {
		return nil, fmt.Errorf("lesson %q not found in manifest", id)
	}

	base := "lessons/" + id

	mdData, err := golearn.LessonsFS.ReadFile(base + "/lesson.md")
	if err != nil {
		return nil, fmt.Errorf("lesson.md: %w", err)
	}
	starter, err := golearn.LessonsFS.ReadFile(base + "/starter/main.go.tmpl")
	if err != nil {
		return nil, fmt.Errorf("starter/main.go.tmpl: %w", err)
	}
	solution, err := golearn.LessonsFS.ReadFile(base + "/solution/main.go.tmpl")
	if err != nil {
		return nil, fmt.Errorf("solution/main.go.tmpl: %w", err)
	}
	test, err := golearn.LessonsFS.ReadFile(base + "/tests/main_test.go.tmpl")
	if err != nil {
		return nil, fmt.Errorf("tests/main_test.go.tmpl: %w", err)
	}

	sections := parseSections(string(mdData))

	l := &Lesson{
		ID:       meta.ID,
		Title:    meta.Title,
		Category: meta.Category,
		Theory:   strings.TrimSpace(sections["THEORY"]),
		Example:  strings.TrimSpace(extractCodeBlock(sections["EXAMPLE"])),
		Task:     strings.TrimSpace(sections["TASK"]),
		Hints:    parseHints(sections["HINTS"]),
		Starter:  string(starter),
		Solution: string(solution),
		Test:     string(test),
	}
	return l, nil
}

// parseSections splits a lesson.md body into its "## SECTION" blocks.
// Headers are matched case-insensitively so lesson authors can write
// "## Theory" or "## THEORY" interchangeably.
func parseSections(md string) map[string]string {
	sections := make(map[string]string)
	scanner := bufio.NewScanner(strings.NewReader(md))
	scanner.Buffer(make([]byte, 0, 64*1024), 1024*1024)

	current := ""
	var buf strings.Builder
	flush := func() {
		if current != "" {
			sections[current] = buf.String()
		}
		buf.Reset()
	}

	for scanner.Scan() {
		line := scanner.Text()
		trimmed := strings.TrimSpace(line)
		if strings.HasPrefix(trimmed, "## ") {
			flush()
			current = strings.ToUpper(strings.TrimSpace(strings.TrimPrefix(trimmed, "## ")))
			continue
		}
		if strings.HasPrefix(trimmed, "# ") {
			// title line, ignored — title comes from manifest.json
			continue
		}
		if current != "" {
			buf.WriteString(line)
			buf.WriteString("\n")
		}
	}
	flush()
	return sections
}

// extractCodeBlock pulls the contents of the first ```go fenced block
// out of a section; if there is no fence it returns the section as-is.
func extractCodeBlock(section string) string {
	lines := strings.Split(section, "\n")
	inFence := false
	var out []string
	found := false
	for _, line := range lines {
		t := strings.TrimSpace(line)
		if strings.HasPrefix(t, "```") {
			if inFence {
				break
			}
			inFence = true
			found = true
			continue
		}
		if inFence {
			out = append(out, line)
		}
	}
	if !found {
		return section
	}
	return strings.Join(out, "\n")
}

// parseHints reads a numbered list ("1. ...", "2. ...", "3. ...") into
// an ordered slice, tolerating blank lines between entries.
func parseHints(section string) []string {
	var hints []string
	scanner := bufio.NewScanner(strings.NewReader(section))
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		dot := strings.Index(line, ".")
		if dot <= 0 {
			continue
		}
		if _, err := strconv.Atoi(line[:dot]); err != nil {
			continue
		}
		hints = append(hints, strings.TrimSpace(line[dot+1:]))
	}
	return hints
}

// Markdown reconstructs a lesson.md-shaped document from the parsed
// sections, for display in an editor side-pane (Vim/Neovim's own
// Markdown+Go syntax highlighting renders it, no ANSI needed there).
func (l *Lesson) Markdown() string {
	var b strings.Builder
	fmt.Fprintf(&b, "# %s — %s\n\n", l.ID, l.Title)
	b.WriteString("## THEORY\n\n")
	b.WriteString(l.Theory)
	b.WriteString("\n\n")
	if l.Example != "" {
		b.WriteString("## EXAMPLE\n\n```go\n")
		b.WriteString(l.Example)
		b.WriteString("\n```\n\n")
	}
	b.WriteString("## TASK\n\n")
	b.WriteString(l.Task)
	b.WriteString("\n")
	return b.String()
}
