package main

import (
	"fmt"
	"os"

	"github.com/abdulahshoaib/evidence-desk/internal/app"
)

func main() {
	if err := app.Run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, "evidence-desk:", err)
		os.Exit(1)
	}
}
