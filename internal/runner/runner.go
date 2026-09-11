// Package runner manages the learner's editable workspace copy of each
// lesson's starter code and opens it in Vim.
package runner

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

// WorkspaceDir returns ~/.go-learn/workspace/<lessonID>, creating it if
// it doesn't exist.
func WorkspaceDir(lessonID string) (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	dir := filepath.Join(home, ".go-learn", "workspace", lessonID)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return "", err
	}
	return dir, nil
}

// EnsureWorkspaceFile makes sure main.go exists in the lesson's
// workspace, seeding it from the starter template on first visit only —
// re-visiting a lesson never overwrites code the learner already wrote.
func EnsureWorkspaceFile(lessonID, starter string) (string, error) {
	dir, err := WorkspaceDir(lessonID)
	if err != nil {
		return "", err
	}
	path := filepath.Join(dir, "main.go")
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err := os.WriteFile(path, []byte(starter), 0o644); err != nil {
			return "", err
		}
	} else if err != nil {
		return "", err
	}
	return path, nil
}

// ReadWorkspaceFile returns the learner's current code for a lesson.
func ReadWorkspaceFile(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// WriteLessonContent (re)writes the lesson's theory/example/task as a
// Markdown file in the workspace, so it can sit in a side pane next to
// the code being edited. Unlike main.go, this is fully regenerated on
// every visit — it always reflects the current lesson content.
func WriteLessonContent(lessonID, markdown string) (string, error) {
	dir, err := WorkspaceDir(lessonID)
	if err != nil {
		return "", err
	}
	path := filepath.Join(dir, "LESSON.md")
	if err := os.WriteFile(path, []byte(markdown), 0o644); err != nil {
		return "", err
	}
	return path, nil
}

func resolveEditor() string {
	editor := os.Getenv("VISUAL")
	if editor == "" {
		editor = os.Getenv("EDITOR")
	}
	if editor == "" {
		editor = "vim"
	}
	return editor
}

// InNestedNvimTerminal reports whether go-learn is running inside a
// :terminal buffer that a Neovim instance opened (Neovim sets $NVIM to
// that instance's RPC server address for every :terminal it spawns —
// this is how "go-learn check"/"go-learn next" run from a Ctrl+T split
// know a parent editor is already open).
func InNestedNvimTerminal() bool {
	return os.Getenv("NVIM") != ""
}

// OpenEditorInPlace opens contentPath and codePath in the parent
// Neovim instance (found via $NVIM) without spawning a nested editor
// process and — critically — without ever creating a new tab. It
// reuses the same two windows across every call in a session:
//   - First call in a tab: finds a non-:terminal window to repurpose
//     as the content pane (or opens one above the terminal if every
//     window in the tab is a terminal), marks it and a vsplit next to
//     it with window-local golearn_content/golearn_code variables.
//   - Every later call (a different lesson): finds those two marked
//     windows again and just swaps their buffers — the terminal
//     window (where this command itself is running) is never touched.
//
// This mirrors what OpenEditor's "-O" flow already sets up on the very
// first (non-nested) launch — that flow marks its two windows the same
// way, so a later "go-learn next" from a Ctrl+T split finds and reuses
// them instead of creating anything new.
//
// Uses --remote-send (not --remote-tab, which opens one tab per file
// rather than a split, and not a blocking wait — Neovim has none;
// "E5600: Wait commands not yet implemented in Nvim" — confirmed) so
// this returns as soon as the buffers are swapped; it does not wait
// for the learner to finish editing. Use InNestedNvimTerminal to
// decide when this applies.
func OpenEditorInPlace(contentPath, codePath string) error {
	addr := os.Getenv("NVIM")
	keys := fmt.Sprintf(
		// <C-\><C-n> — not <Esc> — because the target window is a
		// :terminal buffer running the very shell go-learn was invoked
		// from: it's in terminal-job mode, where plain <Esc> does
		// nothing and any keys sent next would go straight to the shell
		// inside it instead of being read as Ex commands.
		// Script-local (s:) variables only work inside a :source'd
		// script, not in commands typed/sent ad hoc like this ("E461:
		// Illegal variable name") — so scratch state uses uniquely
		// prefixed globals (g:golearn_*) instead.
		"<C-\\><C-n>"+
			":let g:golearn_cw = 0 | let g:golearn_kw = 0<CR>"+
			":for g:golearn_w in range(1, winnr('$')) | "+
			// The buftype!='terminal' guard is cheap insurance against
			// ever editing over the very terminal this command is
			// running in, in case a window that once held marked
			// content/code somehow ends up hosting a terminal buffer —
			// the normal Ctrl+T mapping always opens its terminal in a
			// fresh window, so this shouldn't trigger in practice.
			"if getwinvar(g:golearn_w, 'golearn_content', 0) && getwinvar(g:golearn_w, '&buftype') !=# 'terminal' | let g:golearn_cw = g:golearn_w | endif | "+
			"if getwinvar(g:golearn_w, 'golearn_code', 0) && getwinvar(g:golearn_w, '&buftype') !=# 'terminal' | let g:golearn_kw = g:golearn_w | endif | "+
			"endfor<CR>"+
			":if g:golearn_cw > 0 && g:golearn_kw > 0<CR>"+
			":execute g:golearn_cw . 'wincmd w'<CR>"+
			":edit %s<CR>"+
			":setlocal nomodifiable<CR>"+
			":execute g:golearn_kw . 'wincmd w'<CR>"+
			":edit %s<CR>"+
			":else<CR>"+
			":let g:golearn_target = 0<CR>"+
			":for g:golearn_w in range(1, winnr('$')) | "+
			"if getwinvar(g:golearn_w, '&buftype') !=# 'terminal' | let g:golearn_target = g:golearn_w | break | endif | "+
			"endfor<CR>"+
			":if g:golearn_target > 0<CR>"+
			":execute g:golearn_target . 'wincmd w'<CR>"+
			":else<CR>"+
			":topleft new<CR>"+
			":endif<CR>"+
			":edit %s<CR>"+
			":let w:golearn_content = 1<CR>"+
			":setlocal nomodifiable<CR>"+
			":rightbelow vsplit %s<CR>"+
			":let w:golearn_code = 1<CR>"+
			":endif<CR>"+
			":wincmd =<CR>"+
			":for g:golearn_w in range(1, winnr('$')) | "+
			"if getwinvar(g:golearn_w, 'golearn_code', 0) | execute g:golearn_w . 'wincmd w' | endif | "+
			"endfor<CR>"+
			// Close the terminal(s) in this tab last — this is the very
			// window/job the command sending all of this is running
			// inside of. bwipeout! force-kills the still-running shell
			// job, which also tears down this "nvim --remote-send"
			// client process itself; that's fine because --remote-send
			// only needs to *deliver* these keys to the server, not stay
			// alive while the server executes them, so the close still
			// completes normally on the server side. Iterate windows
			// highest-to-lowest since closing shifts window numbers.
			":for g:golearn_tw in range(winnr('$'), 1, -1) | "+
			"if getwinvar(g:golearn_tw, '&buftype') ==# 'terminal' | execute g:golearn_tw . 'bwipeout!' | endif | "+
			"endfor<CR>",
		contentPath, codePath, contentPath, codePath,
	)
	cmd := exec.Command("nvim", "--server", addr, "--remote-send", keys)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("asosiy Vim'ga ulanib bo'lmadi: %w", err)
	}
	return nil
}

// OpenEditor opens contentPath and codePath side by side in a vertical
// split (Vim/Neovim's -O flag) and blocks until the user exits
// (":qa" / closing both windows). Editors other than vim/nvim don't
// reliably support -O, so for those we fall back to opening just the
// code file. It honors $EDITOR/$VISUAL if set, falling back to vim.
func OpenEditor(contentPath, codePath string) error {
	editor := resolveEditor()
	base := filepath.Base(editor)

	var cmd *exec.Cmd
	if base == "vim" || base == "nvim" {
		// -O opens both files in vertical splits (content on the left,
		// code on the right), leaving focus on the first file (content).
		// "let w:golearn_content/golearn_code = 1" marks each window so
		// that a later "go-learn check"/"next" run from a Ctrl+T
		// terminal split (see OpenEditorInPlace) can find and reuse
		// these exact two windows instead of creating new ones.
		// "setlocal nomodifiable" — applied to the content window before
		// moving on — makes it read-only: if focus ever accidentally
		// lands back on it (a stray window-switch, a misclick), keys
		// typed there are rejected instead of silently corrupting the
		// lesson text. "wincmd =" forces an even 50/50 width split
		// regardless of anything (a plugin, a restored session, terminal
		// quirks) that might otherwise skew it. "wincmd l" then moves
		// focus one window to the right, so the learner lands directly
		// in their code, with the lesson content visible alongside it.
		cmd = exec.Command(editor, "-O", contentPath, codePath,
			"-c", "let w:golearn_content = 1",
			"-c", "setlocal nomodifiable",
			"-c", "wincmd =",
			"-c", "wincmd l",
			"-c", "let w:golearn_code = 1")
	} else {
		cmd = exec.Command(editor, codePath)
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("editor exited with error: %w", err)
	}
	return nil
}
