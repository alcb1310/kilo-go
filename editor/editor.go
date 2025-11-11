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
	rows, cols  int
	cx, cy      int
	numrows     int
	row         EditorRow
}

func NewEditor(f func()) *EditorConfig {
	rows, cols, err := utils.GetWindowSize()
	if err != nil {
		utils.SafeExit(f, err)
	}

	return &EditorConfig{
		restoreFunc: f,
		reader:      bufio.NewReader(os.Stdin),
		rows:        rows,
		cols:        cols,
		cx:          0,
		cy:          0,
		numrows:     0,
	}
}

func (e *EditorConfig) EditorLoop() {
	defer utils.SafeExit(e.restoreFunc, nil)
	e.editorOpen()

	for {
		e.editorRefreshScreen()
		e.editorProcessKeypress()
	}
}
