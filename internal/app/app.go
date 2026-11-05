package app

import (
	"fmt"
	"io"
	"os"
)

// Version is the current build version.
const Version = "0.1.0"

// Run executes the application with the given arguments, writing output to
// stdout.
func Run(args []string) error {
	return RunWithOutput(args, os.Stdout)
}

// RunWithOutput executes the application, writing output to w.
func RunWithOutput(args []string, w io.Writer) error {
	_ = args
	_, err := fmt.Fprintln(w, "evidence-desk", Version)
	return err
}
