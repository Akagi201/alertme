package main

import (
	"io"
	"os"

	"github.com/gen2brain/beeep"
)

func main() {
	// Redirect stdin to stdout
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		io.Copy(os.Stdout, os.Stdin)
	}

	beeep.Notify("Hey!", "Task finished. You may want to take a look.", "")
}
