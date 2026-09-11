package progress

import (
	"testing"
)

func withTempHome(t *testing.T) {
	t.Helper()
	t.Setenv("HOME", t.TempDir())
}

func TestLoadWithNoFileReturnsEmpty(t *testing.T) {
	withTempHome(t)

	p, err := Load()
	if err != nil {
		t.Fatalf("Load() error = %v", err)
	}
	if len(p.Completed) != 0 {
		t.Errorf("Expected empty Completed, got %v", p.Completed)
	}
}

func TestMarkCompletedPersists(t *testing.T) {
	withTempHome(t)

	p, err := Load()
	if err != nil {
		t.Fatalf("Load() error = %v", err)
	}
	if err := p.MarkCompleted("01-hello"); err != nil {
		t.Fatalf("MarkCompleted() error = %v", err)
	}

	reloaded, err := Load()
	if err != nil {
		t.Fatalf("Load() error = %v", err)
	}
	if !reloaded.IsCompleted("01-hello") {
		t.Errorf("Expected 01-hello to be completed after reload, got %v", reloaded.Completed)
	}
}

func TestMarkCompletedIsIdempotent(t *testing.T) {
	withTempHome(t)

	p, _ := Load()
	p.MarkCompleted("01-hello")
	p.MarkCompleted("01-hello")

	if len(p.Completed) != 1 {
		t.Errorf("Expected 1 completed lesson after marking twice, got %d: %v", len(p.Completed), p.Completed)
	}
}

func TestReset(t *testing.T) {
	withTempHome(t)

	p, _ := Load()
	p.MarkCompleted("01-hello")
	p.MarkCompleted("02-variables")

	if err := p.Reset(); err != nil {
		t.Fatalf("Reset() error = %v", err)
	}
	if len(p.Completed) != 0 {
		t.Errorf("Expected empty Completed after Reset, got %v", p.Completed)
	}

	reloaded, _ := Load()
	if len(reloaded.Completed) != 0 {
		t.Errorf("Expected empty Completed after reload, got %v", reloaded.Completed)
	}
}
