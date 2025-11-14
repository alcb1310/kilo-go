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
	e.editorDrawStatusBar(abuf)

	fmt.Fprintf(abuf, "%c[%d;%dH", utils.ESC, (e.cy-e.rowoffset)+1, (e.rx-e.colloffset)+1)
	fmt.Fprintf(abuf, "%c[?25h", utils.ESC)

	fmt.Fprintf(os.Stdout, "%s", abuf.Bytes())
}

func (e *EditorConfig) editorDrawRows(abuf *ab.AppendBuffer) {
	for y := range e.screenrows {
		filerow := y + e.rowoffset
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
			chars := e.rows[filerow].render
			if len(chars) < e.colloffset {
				chars = make([]byte, 0)
			} else {
				chars = chars[e.colloffset:]
			}
			if len(chars) > e.screencols {
				chars = chars[e.colloffset:e.screencols]
			}

			fmt.Fprintf(abuf, "%s", chars)
		}

		fmt.Fprintf(abuf, "%c[K", utils.ESC)
		fmt.Fprintf(abuf, "\r\n")
	}
}

func (e *EditorConfig) editorScroll() {
	e.rx = 0
	if e.cy < e.numrows {
		e.rx = editorRowCxToRx(&e.rows[e.cy], e.cx)
	}

	if e.cy < e.rowoffset {
		e.rowoffset = e.cy
	}
	if e.cy >= e.rowoffset+e.screenrows {
		e.rowoffset = e.cy - e.screenrows + 1
	}

	if e.rx < e.colloffset {
		e.colloffset = e.rx
	}
	if e.rx >= e.colloffset+e.screencols {
		e.colloffset = e.rx - e.screencols + 1
	}
}

func (e *EditorConfig) editorDrawStatusBar(abuf *ab.AppendBuffer) {
	status := e.filename
	if status == "" {
		status = "[No Name]"
	}
	width := e.screencols - len(status) - 1

	rstatus := fmt.Sprintf("column: %d row: %d/%d ", e.rx+1, e.cy+1, e.numrows)

	fmt.Fprintf(abuf, "%c[7m", utils.ESC)
	fmt.Fprintf(abuf, " %s", status)
	for k := range width {
		if k+len(rstatus) == width {
			fmt.Fprintf(abuf, "%s", rstatus)
			break
		}
		fmt.Fprintf(abuf, " ")
	}

	fmt.Fprintf(abuf, "%c[m", utils.ESC)
}
