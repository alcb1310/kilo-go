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
	fmt.Fprintf(abuf, "%c[%d;%dH", utils.ESC, e.cy+1, e.cx+1)
	fmt.Fprintf(abuf, "%c[?25h", utils.ESC)

	fmt.Fprintf(os.Stdout, "%s", abuf.Bytes())
}

func (e *EditorConfig) editorDrawRows(abuf *ab.AppendBuffer) {
	for y := range e.rows {
		if y >= e.numrows {
			if e.numrows == 0 && y == e.rows/3 {
				welcomeMessage := fmt.Sprintf("Kilo editor -- version %s", utils.KILO_VERSION)
				welcomeLen := min(len(welcomeMessage), e.cols)

				padding := (e.cols - welcomeLen) / 2
				if padding > 0 {
					fmt.Fprintf(abuf, "~")
					padding--
				}

				for range padding {
					fmt.Fprintf(abuf, " ")
				}

				fmt.Fprintf(abuf, "%s", welcomeMessage)
			} else {
				fmt.Fprintf(abuf, "~")
			}
		} else {
			fmt.Fprintf(abuf, "%s", e.row.chars)
		}

		fmt.Fprintf(abuf, "%c[K", utils.ESC)
		if y < e.rows-1 {
			fmt.Fprintf(abuf, "\r\n")
		}
	}
}
