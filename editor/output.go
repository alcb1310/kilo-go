package editor

import (
	"fmt"
	"os"
)

const (
	ESC = 0x1b
)

func (e *EditorConfig) editorRefreshScreen() {
	fmt.Fprintf(os.Stdout, "%c[2J", ESC)
}
