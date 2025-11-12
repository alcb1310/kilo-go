package editor

import (
	"bufio"
	"os"

	"github.com/alcb1310/kilo-go/utils"
)

type EditorRow struct {
	chars string
}

type EditorConfig struct {
	restoreFunc func()
	reader      *bufio.Reader
	screenrows  int
	screencols  int
	cx, cy      int
	numrows     int
	rowoffset   int
	colloffset  int
	rows        []EditorRow
}

func NewEditor(f func()) *EditorConfig {
	rows, cols, err := utils.GetWindowSize()
	if err != nil {
		utils.SafeExit(f, err)
	}

	return &EditorConfig{
		restoreFunc: f,
		reader:      bufio.NewReader(os.Stdin),
		screenrows:  rows,
		screencols:  cols,
		cx:          0,
		cy:          0,
		numrows:     0,
		rowoffset:   0,
		colloffset:  0,
		rows:        make([]EditorRow, 0),
	}
}

func (e *EditorConfig) EditorLoop() {
	defer utils.SafeExit(e.restoreFunc, nil)

	if len(os.Args) > 1 {
		e.editorOpen(os.Args[1])
	}

	for {
		e.editorRefreshScreen()
		e.editorProcessKeypress()
	}
}
