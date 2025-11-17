package editor

import (
	"bufio"
	"os"

	"github.com/alcb1310/kilo-go/utils"
)

type EditorRow struct {
	chars  string
	render []byte
}

type EditorConfig struct {
	restoreFunc   func()
	reader        *bufio.Reader
	screenrows    int
	screencols    int
	cx, cy        int
	rx            int
	numrows       int
	rowoffset     int
	colloffset    int
	filename      string
	statusMessage string
	rows          []EditorRow
	isDirty       bool
}

func NewEditor(f func()) *EditorConfig {
	rows, cols, err := utils.GetWindowSize()
	if err != nil {
		utils.SafeExit(f, err)
	}

	return &EditorConfig{
		restoreFunc:   f,
		reader:        bufio.NewReader(os.Stdin),
		screenrows:    rows - 2,
		screencols:    cols,
		cx:            0,
		cy:            0,
		numrows:       0,
		rowoffset:     0,
		colloffset:    0,
		rx:            0,
		filename:      "",
		statusMessage: "",
		rows:          make([]EditorRow, 0),
		isDirty:       false,
	}
}

func (e *EditorConfig) EditorLoop() {
	defer utils.SafeExit(e.restoreFunc, nil)

	if len(os.Args) > 1 {
		e.editorOpen(os.Args[1])
	}

	e.editorSetStatusMessage(utils.KILO_DEFAULT_STATUS_MESSAGE)

	for {
		e.editorRefreshScreen()
		e.editorProcessKeypress()
	}
}
