package editor

import (
	"log/slog"

	"github.com/alcb1310/kilo-go/utils"
)

func (e *EditorConfig) editorProcessKeypress() {
	b, err := e.editorReadKey()
	if err != nil {
		utils.SafeExit(e.restoreFunc, err)
	}

	switch b {
	case utils.CtrlKey('q'):
		utils.SafeExit(e.restoreFunc, nil)
	case utils.ARROW_DOWN, utils.ARROW_LEFT, utils.ARROW_RIGHT, utils.ARROW_UP:
		e.editorMoveCursor(b)
	case utils.PAGE_DOWN, utils.PAGE_UP:
		times := e.screenrows
		for range times {
			if b == utils.PAGE_DOWN {
				e.editorMoveCursor(utils.ARROW_DOWN)
			} else {
				e.editorMoveCursor(utils.ARROW_UP)
			}
		}
	case utils.DEL_KEY:
		slog.Info("DEL_KEY")
	case utils.HOME_KEY:
		e.cx = 0
	case utils.END_KEY:
		e.cx = e.screencols - 1
	}
}

func (e *EditorConfig) editorMoveCursor(key int) {
	switch key {
	case utils.ARROW_LEFT:
		if e.cx > 0 {
			e.cx--
		}
	case utils.ARROW_DOWN:
		if e.cy < e.numrows {
			e.cy++
		}
	case utils.ARROW_UP:
		if e.cy != 0 {
			e.cy--
		}
	case utils.ARROW_RIGHT:
		if e.cx != e.screencols-1 {
			e.cx++
		}
	}
}
