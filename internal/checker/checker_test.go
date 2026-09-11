package checker

import "testing"

const goodCode = `package main

import "fmt"

func main() {
	fmt.Println("Hello, Go!")
}
`

const badCode = `package main

import "fmt"

func main() {
	fmt.Println("wrong output")
}
`

const test = `package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func captureStdout(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	f()
	w.Close()
	os.Stdout = old
	var buf bytes.Buffer
	io.Copy(&buf, r)
	return buf.String()
}

func TestHelloOutput(t *testing.T) {
	got := strings.TrimSpace(captureStdout(main))
	want := "Hello, Go!"
	if got != want {
		t.Errorf("Expected:\n%s\n\nGot:\n%s", want, got)
	}
}
`

func TestRunPassingCode(t *testing.T) {
	result, err := Run(goodCode, test)
	if err != nil {
		t.Fatalf("Run() error = %v", err)
	}
	if !result.Compiled {
		t.Fatalf("Expected code to compile, output: %s", result.CompileOutput)
	}
	if !result.Passed() {
		t.Errorf("Expected Passed() = true, got false. Raw output:\n%s", result.RawOutput)
	}
}

func TestRunFailingCode(t *testing.T) {
	result, err := Run(badCode, test)
	if err != nil {
		t.Fatalf("Run() error = %v", err)
	}
	if !result.Compiled {
		t.Fatalf("Expected code to compile, output: %s", result.CompileOutput)
	}
	if result.Passed() {
		t.Error("Expected Passed() = false for code that doesn't print anything")
	}
}

func TestRunUncompilableCode(t *testing.T) {
	result, err := Run("this is not go code", test)
	if err != nil {
		t.Fatalf("Run() error = %v", err)
	}
	if result.Compiled {
		t.Error("Expected Compiled = false for invalid Go source")
	}
}
