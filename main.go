package main

import (
	"fmt"
	"os"

	"github.com/alcb1310/kilo-go/editor"
	"github.com/alcb1310/kilo-go/linux"
)

var restoreFunc func()

func init() {
	var err error
	u := linux.NewUnixRawMode()
	restoreFunc, err = u.EnableRawMode()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\r\n", err)
		os.Exit(1)
	}
}

func main() {
	editor := editor.NewEditor(restoreFunc)
	editor.EditorLoop()
}
