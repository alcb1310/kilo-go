package editor

import (
	"bufio"
	"fmt"
	"os"

	"github.com/alcb1310/kilo-go/utils"
)

func (e *EditorConfig) editorOpen(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "editorOpen, error opening file: %v\r\n", err)
		utils.SafeExit(e.restoreFunc, err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		row := EditorRow{chars: scanner.Text()}
		e.rows = append(e.rows, row)
		e.numrows++
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "editorOpen, error scanning file: %v\r\n", err)
		utils.SafeExit(e.restoreFunc, err)
		os.Exit(1)
	}
}
