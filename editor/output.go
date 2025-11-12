package editor

import (
	"fmt"
	"os"

	ab "github.com/alcb1310/kilo-go/appendbuffer"
	"github.com/alcb1310/kilo-go/utils"
)

func (e *EditorConfig) editorRefreshScreen() {
	e.editorScroll()
	abuf := ab.New()

	fmt.Fprintf(abuf, "%c[?25l", utils.ESC)
	fmt.Fprintf(abuf, "%c[H", utils.ESC)

	e.editorDrawRows(abuf)
	fmt.Fprintf(abuf, "%c[%d;%dH", utils.ESC, (e.cy-e.rowoffset)+1, (e.cx-e.colloffset)+1)
	fmt.Fprintf(abuf, "%c[?25h", utils.ESC)

	fmt.Fprintf(os.Stdout, "%s", abuf.Bytes())
}

func (e *EditorConfig) editorDrawRows(abuf *ab.AppendBuffer) {
	for y := range e.screenrows {
		filerow := y + e.rowoffset
		fmt.Fprintf(abuf, "%c[K", utils.ESC)
		if filerow >= e.numrows {
			if e.numrows == 0 && y == e.screenrows/3 {
				welcomeMessage := fmt.Sprintf("Kilo editor -- version %s", utils.KILO_VERSION)
				welcomeLen := min(len(welcomeMessage), e.screencols)

				padding := (e.screencols - welcomeLen) / 2
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
			chars := e.rows[filerow].chars
			if len(chars) < e.colloffset {
				chars = ""
			} else {
				chars = chars[e.colloffset:]
			}
			if len(chars) > e.screencols {
				chars = chars[e.colloffset:e.screencols]
			}

			fmt.Fprintf(abuf, "%s", chars)
		}

		if y < e.screenrows-1 {
			fmt.Fprintf(abuf, "\r\n")
		}
	}
}

func (e *EditorConfig) editorScroll() {
	if e.cy < e.rowoffset {
		e.rowoffset = e.cy
	}
	if e.cy >= e.rowoffset+e.screenrows {
		e.rowoffset = e.cy - e.screenrows + 1
	}

	if e.cx < e.colloffset {
		e.colloffset = e.cx
	}
	if e.cx >= e.colloffset+e.screencols {
		e.colloffset = e.cx - e.screencols + 1
	}
}
