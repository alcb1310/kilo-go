package editor

import (
	"bufio"
	"fmt"
	"log/slog"
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
	e.filename = filename

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		e.editorAppendRow(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "editorOpen, error scanning file: %v\r\n", err)
		utils.SafeExit(e.restoreFunc, err)
		os.Exit(1)
	}
}

func (e *EditorConfig) editorSave() {
	data := make([]byte, 0)
	for _, row := range e.rows {
		data = append(data, row.chars...)
		data = append(data, '\r', '\n')
	}
	err := os.WriteFile(e.filename, data, 0644)
	if err != nil {
		slog.Error("editorSave, error saving file", "error", err)
		e.editorSetStatusMessage(fmt.Sprintf("Error saving file: %v", err))
		return
	}

	e.editorSetStatusMessage("File saved")
}
