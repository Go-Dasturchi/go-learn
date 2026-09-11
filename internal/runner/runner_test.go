package runner

import (
	"os"
	"path/filepath"
	"testing"
)

func withTempHome(t *testing.T) {
	t.Helper()
	t.Setenv("HOME", t.TempDir())
}

func TestEnsureWorkspaceFileSeedsOnceAndPreservesEdits(t *testing.T) {
	withTempHome(t)

	path, err := EnsureWorkspaceFile("01-hello", "package main // starter")
	if err != nil {
		t.Fatalf("EnsureWorkspaceFile() error = %v", err)
	}
	got, _ := ReadWorkspaceFile(path)
	if got != "package main // starter" {
		t.Errorf("expected starter content on first visit, got %q", got)
	}

	if err := os.WriteFile(path, []byte("package main // learner wrote this"), 0o644); err != nil {
		t.Fatalf("WriteFile() error = %v", err)
	}

	path2, err := EnsureWorkspaceFile("01-hello", "package main // starter")
	if err != nil {
		t.Fatalf("EnsureWorkspaceFile() second call error = %v", err)
	}
	got2, _ := ReadWorkspaceFile(path2)
	if got2 != "package main // learner wrote this" {
		t.Errorf("expected learner's edit to survive a second visit, got %q", got2)
	}
}

func TestWriteLessonContentAlwaysOverwrites(t *testing.T) {
	withTempHome(t)

	path, err := WriteLessonContent("01-hello", "# first version")
	if err != nil {
		t.Fatalf("WriteLessonContent() error = %v", err)
	}
	if filepath.Base(path) != "LESSON.md" {
		t.Errorf("expected file named LESSON.md, got %s", path)
	}

	path2, err := WriteLessonContent("01-hello", "# second version")
	if err != nil {
		t.Fatalf("WriteLessonContent() second call error = %v", err)
	}
	data, _ := os.ReadFile(path2)
	if string(data) != "# second version" {
		t.Errorf("expected content to be refreshed on every call, got %q", string(data))
	}
}
