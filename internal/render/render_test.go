package render

import (
	"regexp"
	"strings"
	"testing"
)

var ansiRe = regexp.MustCompile("\x1b\\[[0-9;]*m")

func stripANSI(s string) string {
	return ansiRe.ReplaceAllString(s, "")
}

func TestGoCodePreservesSourceByteForByte(t *testing.T) {
	src := `package main

import "fmt"

func Toifa(age int) string {
	// bu izoh
	if age >= 18 {
		return "Voyaga yetgan"
	}
	return "Voyaga yetmagan"
}

func main() {
	fmt.Println(Toifa(20))
}
`
	got := stripANSI(GoCode(src))
	if got != src {
		t.Errorf("Stripped-color output does not match source.\n--- want ---\n%q\n--- got ---\n%q", src, got)
	}
}

func TestGoCodeColorsKeywordsAndStrings(t *testing.T) {
	out := GoCode(`func main() { return "hi" }`)
	if !strings.Contains(out, magenta+"func"+reset) {
		t.Errorf("expected 'func' to be colored as keyword, got: %q", out)
	}
	if !strings.Contains(out, green+`"hi"`+reset) {
		t.Errorf("expected string literal to be colored, got: %q", out)
	}
}

func TestMarkdownStripsBackticksAndHighlightsFence(t *testing.T) {
	md := "Bu `name` o'zgaruvchi.\n\n```go\nfmt.Println(1)\n```\n"
	out := Markdown(md)
	if strings.Contains(out, "`") {
		t.Errorf("expected no raw backticks in rendered output, got: %q", out)
	}
	if strings.Contains(out, "```") {
		t.Errorf("expected no raw code fences in rendered output, got: %q", out)
	}
	if !strings.Contains(stripANSI(out), "fmt.Println(1)") {
		t.Errorf("expected code block content preserved, got: %q", stripANSI(out))
	}
}

func TestWrapTextNeverBreaksAWord(t *testing.T) {
	text := "if-01-sonni-oshirish-yoki-kamaytirish — Sonni oshirish yoki kamaytirish butun songa"
	wrapped := WrapText(text, 40, "      ")
	words := strings.Fields(text)
	for _, w := range words {
		if !strings.Contains(wrapped, w) {
			t.Errorf("expected word %q to survive intact in wrapped output, got:\n%s", w, wrapped)
		}
	}
	for _, line := range strings.Split(wrapped, "\n") {
		if len([]rune(line)) > 46 { // 40 + up to 6 for the hanging indent
			t.Errorf("line exceeds expected width: %q", line)
		}
	}
}

func TestWrapTextIndentsContinuationLines(t *testing.T) {
	wrapped := WrapText("one two three four five six seven eight", 15, "    ")
	lines := strings.Split(wrapped, "\n")
	if len(lines) < 2 {
		t.Fatalf("expected wrapping to produce multiple lines, got: %q", wrapped)
	}
	for _, line := range lines[1:] {
		if !strings.HasPrefix(line, "    ") {
			t.Errorf("expected continuation line to start with hanging indent, got: %q", line)
		}
	}
}

func TestWrapTextShortTextUnchanged(t *testing.T) {
	got := WrapText("short text", 80, "  ")
	if got != "short text" {
		t.Errorf("expected short text to pass through unwrapped, got: %q", got)
	}
}
