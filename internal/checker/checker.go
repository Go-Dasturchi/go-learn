// Package checker compiles and tests the learner's solution in an
// isolated temporary module, without ever exposing the lesson's
// solution/main.go to them.
package checker

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// TestResult is the outcome of a single Go test function.
type TestResult struct {
	Name   string
	Passed bool
}

// Result is the full outcome of checking one lesson submission.
type Result struct {
	Compiled      bool
	CompileOutput string
	Tests         []TestResult
	RawOutput     string
}

// Passed reports whether the code compiled and every test passed.
func (r Result) Passed() bool {
	if !r.Compiled || len(r.Tests) == 0 {
		return false
	}
	for _, t := range r.Tests {
		if !t.Passed {
			return false
		}
	}
	return true
}

// Run copies the learner's code and the lesson's hidden test file into
// a throwaway module, compiles it, and runs `go test -v`.
func Run(code, test string) (Result, error) {
	dir, err := os.MkdirTemp("", "go-learn-check-*")
	if err != nil {
		return Result{}, err
	}
	defer os.RemoveAll(dir)

	if err := os.WriteFile(filepath.Join(dir, "main.go"), []byte(code), 0o644); err != nil {
		return Result{}, err
	}
	if err := os.WriteFile(filepath.Join(dir, "main_test.go"), []byte(test), 0o644); err != nil {
		return Result{}, err
	}
	if err := os.WriteFile(filepath.Join(dir, "go.mod"), []byte("module checkrun\n\ngo 1.21\n"), 0o644); err != nil {
		return Result{}, err
	}

	var res Result

	buildOut, buildErr := runGo(dir, "build", "-o", os.DevNull, ".")
	if buildErr != nil {
		res.Compiled = false
		res.CompileOutput = buildOut
		return res, nil
	}
	res.Compiled = true

	testOut, _ := runGo(dir, "test", "-v", ".")
	res.RawOutput = testOut
	res.Tests = parseTestOutput(testOut)
	return res, nil
}

func runGo(dir string, args ...string) (string, error) {
	cmd := exec.Command("go", args...)
	cmd.Dir = dir
	out, err := cmd.CombinedOutput()
	return string(out), err
}

// parseTestOutput scans `go test -v` output for "--- PASS: Name" and
// "--- FAIL: Name" lines to build a per-test checklist.
func parseTestOutput(output string) []TestResult {
	var results []TestResult
	for _, line := range strings.Split(output, "\n") {
		line = strings.TrimSpace(line)
		switch {
		case strings.HasPrefix(line, "--- PASS: "):
			results = append(results, TestResult{Name: extractTestName(line), Passed: true})
		case strings.HasPrefix(line, "--- FAIL: "):
			results = append(results, TestResult{Name: extractTestName(line), Passed: false})
		}
	}
	return results
}

func extractTestName(line string) string {
	rest := strings.TrimPrefix(line, "--- PASS: ")
	rest = strings.TrimPrefix(rest, "--- FAIL: ")
	if idx := strings.Index(rest, " "); idx != -1 {
		rest = rest[:idx]
	}
	return rest
}
