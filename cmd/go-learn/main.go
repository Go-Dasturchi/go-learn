// Command go-learn is the entry point for the GO ZERO TO HERO
// interactive terminal course.
package main

import (
	"fmt"
	"os"

	"go-learn/internal/cli"
)

func main() {
	if err := cli.Run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, "xato:", err)
		os.Exit(1)
	}
}
