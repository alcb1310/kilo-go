package editor

import (
	"fmt"
	"os"

	ab "github.com/alcb1310/kilo-go/appendbuffer"
	"github.com/alcb1310/kilo-go/utils"
)

func (e *EditorConfig) editorRefreshScreen() {
	abuf := ab.New()

	fmt.Fprintf(abuf, "%c[?25l", utils.ESC)
	fmt.Fprintf(abuf, "%c[H", utils.ESC)

	e.editorDrawRows(abuf)

	fmt.Fprintf(abuf, "%c[H", utils.ESC)
	fmt.Fprintf(abuf, "%c[?25h", utils.ESC)

	fmt.Fprintf(os.Stdout, "%s", abuf.Bytes())
}

func (e *EditorConfig) editorDrawRows(abuf *ab.AppendBuffer) {
	for y := range e.rows {
		if y == e.rows/3 {
			welcomeMessage := fmt.Sprintf("Kilo editor -- version %s", utils.KILO_VERSION)
			fmt.Fprintf(abuf, "%s", welcomeMessage)
		} else {
			fmt.Fprintf(abuf, "~")
		}

		fmt.Fprintf(abuf, "%c[K", utils.ESC)
		if y < e.rows-1 {
			fmt.Fprintf(abuf, "\r\n")
		}
	}
}
