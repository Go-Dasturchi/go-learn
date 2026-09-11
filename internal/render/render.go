// Package render turns raw lesson.md text and Go source snippets into
// colorized terminal output: ANSI colors for section chrome, and real
// lexical syntax highlighting for Go code (via go/scanner — no
// external dependency needed).
package render

import (
	"fmt"
	"go/scanner"
	"go/token"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

// ANSI SGR codes. Plain 16-color codes are used (rather than 256-color
// or truecolor) so this renders correctly in the widest range of
// terminals.
const (
	reset   = "\x1b[0m"
	bold    = "\x1b[1m"
	dim     = "\x1b[2m"
	cyan    = "\x1b[36m"
	green   = "\x1b[32m"
	red     = "\x1b[31m"
	yellow  = "\x1b[33m"
	magenta = "\x1b[35m"
	gray    = "\x1b[90m"
)

func c(color, text string) string { return color + text + reset }

// Banner renders the top-of-screen GO ZERO TO HERO box.
func Banner() string {
	return c(bold+cyan, "╔════════════════════════════════════════════╗") + "\n" +
		c(bold+cyan, "║") + c(bold, "              GO ZERO TO HERO               ") + c(bold+cyan, "║") + "\n" +
		c(bold+cyan, "║") + "          Interactive Go Course             " + c(bold+cyan, "║") + "\n" +
		c(bold+cyan, "╚════════════════════════════════════════════╝")
}

// ProgressBar renders a filled/empty block bar, e.g. "███░░░░ 42%".
func ProgressBar(pct float64, width int) string {
	filled := int(pct / 100 * float64(width))
	if filled > width {
		filled = width
	}
	if filled < 0 {
		filled = 0
	}
	return c(green, strings.Repeat("█", filled)) + c(gray, strings.Repeat("░", width-filled))
}

// SectionHeader renders a lesson section title ("THEORY", "EXAMPLE", ...)
// with an underline, matching the look already used by the CLI.
func SectionHeader(title string) string {
	return c(bold+yellow, title) + "\n" + c(gray, strings.Repeat("─", 40))
}

// Success/Fail wrap a single line with a colored status glyph.
func Success(msg string) string { return c(green, "✓ "+msg) }
func Fail(msg string) string    { return c(red, "✗ "+msg) }

// Dim renders diagnostic text (compiler/test output) in a muted color
// so it reads clearly as "output", not as part of the lesson prose.
func Dim(text string) string { return c(gray, text) }

// Check renders a checklist marker: ✅ for done, ⬜ for not done.
func Check(done bool) string {
	if done {
		return "✅"
	}
	return "⬜"
}

// TerminalWidth returns the current terminal's column count, falling
// back to 80 if it can't be determined (not a terminal, `stty`
// missing, etc). Unlike a Vim buffer, a raw terminal never word-wraps
// program output on its own — it hard-wraps mid-word at the exact
// column — so long menu lines must be wrapped by the program itself.
func TerminalWidth() int {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		return 80
	}
	parts := strings.Fields(string(out))
	if len(parts) != 2 {
		return 80
	}
	cols, err := strconv.Atoi(parts[1])
	if err != nil || cols <= 20 {
		return 80
	}
	return cols
}

// WrapText word-wraps text to fit within width columns, indenting every
// line after the first with indent (a hanging indent, so wrapped text
// reads as a continuation rather than starting back at column 0).
func WrapText(text string, width int, indent string) string {
	if width <= len(indent)+10 {
		width = 80
	}
	words := strings.Fields(text)
	if len(words) == 0 {
		return text
	}
	var lines []string
	var cur strings.Builder
	curLen := 0
	for _, w := range words {
		wl := len([]rune(w))
		addSep := curLen > 0
		sepLen := 0
		if addSep {
			sepLen = 1
		}
		if curLen+sepLen+wl > width && curLen > 0 {
			lines = append(lines, cur.String())
			cur.Reset()
			curLen = 0
			addSep = false
		}
		if addSep {
			cur.WriteByte(' ')
			curLen++
		}
		cur.WriteString(w)
		curLen += wl
	}
	if curLen > 0 {
		lines = append(lines, cur.String())
	}
	for i := 1; i < len(lines); i++ {
		lines[i] = indent + lines[i]
	}
	return strings.Join(lines, "\n")
}

// Markdown renders lesson.md body text: fenced ```go code blocks get
// real syntax highlighting inside a left-bordered panel, and inline
// `code` spans get colorized without their backticks. Everything else
// passes through unchanged.
func Markdown(text string) string {
	lines := strings.Split(text, "\n")
	var out strings.Builder
	i := 0
	for i < len(lines) {
		trimmed := strings.TrimSpace(lines[i])
		if strings.HasPrefix(trimmed, "```") {
			i++
			start := i
			for i < len(lines) && !strings.HasPrefix(strings.TrimSpace(lines[i]), "```") {
				i++
			}
			code := strings.Join(lines[start:i], "\n")
			out.WriteString(codeBlock(code))
			out.WriteString("\n")
			i++ // skip closing fence
			continue
		}
		out.WriteString(inline(lines[i]))
		out.WriteString("\n")
		i++
	}
	return strings.TrimRight(out.String(), "\n")
}

func codeBlock(code string) string {
	highlighted := GoCode(code)
	var out strings.Builder
	border := c(gray, "  │ ")
	for _, l := range strings.Split(highlighted, "\n") {
		out.WriteString(border)
		out.WriteString(l)
		out.WriteString("\n")
	}
	return strings.TrimRight(out.String(), "\n")
}

// inline replaces `code` spans in a single line of prose with
// colorized text (backticks stripped).
func inline(line string) string {
	var out strings.Builder
	var cur strings.Builder
	inCode := false
	for _, r := range line {
		if r == '`' {
			if inCode {
				out.WriteString(c(bold+cyan, cur.String()))
			} else {
				out.WriteString(cur.String())
			}
			cur.Reset()
			inCode = !inCode
			continue
		}
		cur.WriteRune(r)
	}
	out.WriteString(cur.String())
	return out.String()
}

// GoCode lexically tokenizes Go source with go/scanner and re-emits it
// with ANSI colors per token kind, preserving the original whitespace
// and formatting byte-for-byte.
func GoCode(src string) string {
	fset := token.NewFileSet()
	file := fset.AddFile("", fset.Base(), len(src))

	var s scanner.Scanner
	s.Init(file, []byte(src), nil, scanner.ScanComments)

	srcBytes := []byte(src)
	var out strings.Builder
	lastOffset := 0

	for {
		pos, tok, lit := s.Scan()
		if tok == token.EOF {
			break
		}
		offset := fset.Position(pos).Offset

		// Implicit semicolons (Go's automatic-semicolon-insertion) don't
		// correspond to real source bytes — skip them; the newline that
		// triggered them is copied verbatim by the next token's gap-fill.
		if tok == token.SEMICOLON && lit == "\n" {
			continue
		}

		if offset > lastOffset && offset <= len(srcBytes) {
			out.Write(srcBytes[lastOffset:offset])
		}

		text := lit
		if text == "" {
			text = tok.String()
		}

		switch {
		case tok.IsKeyword():
			out.WriteString(c(magenta, text))
		case tok == token.STRING || tok == token.CHAR:
			out.WriteString(c(green, text))
		case tok == token.COMMENT:
			out.WriteString(c(gray, text))
		case tok == token.INT || tok == token.FLOAT:
			out.WriteString(c(yellow, text))
		case tok == token.IDENT && (text == "true" || text == "false" || text == "nil" || text == "iota"):
			out.WriteString(c(yellow, text))
		default:
			out.WriteString(text)
		}
		lastOffset = offset + len(text)
	}
	if lastOffset < len(srcBytes) {
		out.Write(srcBytes[lastOffset:])
	}
	return out.String()
}

// Boxf renders a bold cyan "SOLUTION:"-style label followed by a
// horizontal rule, for use above a revealed solution.
func Boxf(format string, args ...interface{}) string {
	return c(bold+cyan, fmt.Sprintf(format, args...)) + "\n" + c(gray, strings.Repeat("─", 40))
}
