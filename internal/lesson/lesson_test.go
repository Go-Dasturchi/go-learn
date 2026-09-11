package lesson

import (
	"strings"
	"testing"
)

func TestLoadManifestHasSixMVPLessons(t *testing.T) {
	metas, err := LoadManifest()
	if err != nil {
		t.Fatalf("LoadManifest() error = %v", err)
	}
	if len(metas) < 6 {
		t.Fatalf("Expected at least 6 lessons in manifest, got %d", len(metas))
	}
	if metas[0].ID != "01-hello" {
		t.Errorf("Expected first lesson id = 01-hello, got %s", metas[0].ID)
	}
}

func TestLoadParsesAllSections(t *testing.T) {
	l, err := Load("01-hello")
	if err != nil {
		t.Fatalf("Load(01-hello) error = %v", err)
	}
	if l.Title != "Hello World" {
		t.Errorf("Expected title = Hello World, got %s", l.Title)
	}
	if l.Theory == "" {
		t.Error("Expected non-empty Theory")
	}
	if l.Example == "" {
		t.Error("Expected non-empty Example")
	}
	if l.Task == "" {
		t.Error("Expected non-empty Task")
	}
	if len(l.Hints) < 2 {
		t.Errorf("Expected at least 2 hints, got %d", len(l.Hints))
	}
	if l.Starter == "" || l.Solution == "" || l.Test == "" {
		t.Error("Expected non-empty Starter, Solution and Test source")
	}
}

func TestLoadUnknownLessonFails(t *testing.T) {
	if _, err := Load("does-not-exist"); err == nil {
		t.Error("Expected error loading unknown lesson id, got nil")
	}
}

func TestMarkdownReconstructsReadableSections(t *testing.T) {
	l, err := Load("01-hello")
	if err != nil {
		t.Fatalf("Load(01-hello) error = %v", err)
	}
	md := l.Markdown()
	for _, want := range []string{"# 01-hello — Hello World", "## THEORY", "## EXAMPLE", "```go", "## TASK"} {
		if !strings.Contains(md, want) {
			t.Errorf("Markdown() missing %q, got:\n%s", want, md)
		}
	}
}
