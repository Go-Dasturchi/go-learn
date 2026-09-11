// Package golearn exposes the embedded lesson content so it ships
// inside the compiled binary and works regardless of the current
// working directory the user runs go-learn from.
package golearn

import "embed"

//go:embed all:lessons
var LessonsFS embed.FS
