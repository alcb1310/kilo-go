package editor

import (
	"fmt"
	"os"

	ab "github.com/alcb1310/kilo-go/appendbuffer"
	"github.com/alcb1310/kilo-go/utils"
)

func (e *EditorConfig) editorRefreshScreen() {
	abuf := ab.New()

	fmt.Fprintf(abuf, "%c[2J", utils.ESC)
	fmt.Fprintf(abuf, "%c[H", utils.ESC)

	e.editorDrawRows(abuf)

	fmt.Fprintf(abuf, "%c[H", utils.ESC)

	fmt.Fprintf(os.Stdout, "%s", abuf.Bytes())
}

func (e *EditorConfig) editorDrawRows(abuf *ab.AppendBuffer) {
	for y := range e.rows {
		fmt.Fprintf(abuf, "~")

		if y < e.rows-1 {
			fmt.Fprintf(abuf, "\r\n")
		}
	}
}
