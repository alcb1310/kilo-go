package editor

import (
	"bufio"
	"os"

	"github.com/alcb1310/kilo-go/utils"
)

type EditorConfig struct {
	restoreFunc func()
	reader      *bufio.Reader
}

func NewEditor(f func()) *EditorConfig {
	return &EditorConfig{
		restoreFunc: f,
		reader:      bufio.NewReader(os.Stdin),
	}
}

func (e *EditorConfig) EditorLoop() {
	defer utils.SafeExit(e.restoreFunc, nil)

	for {
		e.editorProcessKeypress()
	}
}
