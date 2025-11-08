package editor

import (
	"fmt"
	"os"

	"github.com/alcb1310/kilo-go/utils"
)

func (e *EditorConfig) editorRefreshScreen() {
	fmt.Fprintf(os.Stdout, "%c[2J", utils.ESC)
	fmt.Fprintf(os.Stdout, "%c[H", utils.ESC)

	e.editorDrawRows()

	fmt.Fprintf(os.Stdout, "%c[H", utils.ESC)
}

func (e *EditorConfig) editorDrawRows() {
	for y := range e.rows {
		fmt.Fprintf(os.Stdout, "~")

		if y < e.rows-1 {
			fmt.Fprintf(os.Stdout, "\r\n")
		}
	}
}
