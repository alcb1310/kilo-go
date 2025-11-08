package editor

import (
	"bufio"
	"os"

	"github.com/alcb1310/kilo-go/utils"
)

type EditorConfig struct {
	restoreFunc func()
	reader      *bufio.Reader
	rows, cols  int
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
	}
}

func (e *EditorConfig) EditorLoop() {
	defer utils.SafeExit(e.restoreFunc, nil)

	for {
		e.editorRefreshScreen()
		e.editorProcessKeypress()
	}
}
