package editor

import (
	"fmt"
	"log/slog"

	"github.com/alcb1310/kilo-go/utils"
)

var quit_times = utils.KILO_QUIT_TIMES

func (e *EditorConfig) editorProcessKeypress() {
	b, err := e.editorReadKey()
	if err != nil {
		utils.SafeExit(e.restoreFunc, err)
	}

	switch b {
	case utils.ENTER:
		e.editorInsertNewline()

	case utils.CtrlKey('q'):
		if e.isDirty && quit_times > 0 {
			e.editorSetStatusMessage(fmt.Sprintf("WARNING! File has unsaved changes. Ctrl-Q %d more times to quit", quit_times))
			quit_times--
			return
		}
		utils.SafeExit(e.restoreFunc, nil)
	case utils.CtrlKey('s'):
		e.editorSave()
	case utils.CtrlKey('f'):
		e.editorFind()
		e.editorSetStatusMessage(utils.KILO_DEFAULT_STATUS_MESSAGE)

	case utils.ARROW_DOWN, utils.ARROW_LEFT, utils.ARROW_RIGHT, utils.ARROW_UP:
		e.editorMoveCursor(b)
	case utils.PAGE_DOWN:
		e.cy = min(e.rowoffset+e.screenrows+1, e.numrows)
		times := e.screenrows
		for range times {
			e.editorMoveCursor(utils.ARROW_DOWN)
		}
	case utils.PAGE_UP:
		e.cy = e.rowoffset
		times := e.screenrows
		for range times {
			e.editorMoveCursor(utils.ARROW_UP)
		}
	case utils.DEL_KEY, utils.BACKSPACE:
		if b == utils.DEL_KEY {
			e.editorMoveCursor(utils.ARROW_RIGHT)
		}
		e.editorDeleteChar()
	case utils.HOME_KEY:
		e.cx = 0
	case utils.END_KEY:
		if e.cy < e.numrows {
			e.cx = len(e.rows[e.cy].chars)
		}

	case utils.ESC:
		// for now we will ignore when the user press the escape key
		break
	default:
		e.editorInsertChar(byte(b))
	}

	quit_times = utils.KILO_QUIT_TIMES
}

func (e *EditorConfig) editorMoveCursor(key int) {
	var row *EditorRow = nil
	if e.cy < e.numrows {
		row = &e.rows[e.cy]
	}

	switch key {
	case utils.ARROW_LEFT:
		if e.cx != 0 {
			e.cx--
		} else if e.cy > 0 {
			e.cy--
			e.cx = len(e.rows[e.cy].chars)
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
		if row != nil && e.cx < len(row.chars) {
			e.cx++
		} else if row != nil && e.cx == len(row.chars) {
			e.cy++
			e.cx = 0
		}
	}

	row = nil
	if e.cy < e.numrows {
		row = &e.rows[e.cy]
	}

	if row != nil && e.cx > len(row.chars) {
		e.cx = len(row.chars)
	}
}

func (e *EditorConfig) editorPrompt(prompt string) string {
	var buf string
	for {
		e.editorSetStatusMessage(prompt + buf)
		e.editorRefreshScreen()

		b, err := e.editorReadKey()
		if err != nil {
			utils.SafeExit(e.restoreFunc, err)
		}

		switch b {
		case utils.BACKSPACE, utils.DEL_KEY:
			if len(buf) > 0 {
				buf = buf[:len(buf)-1]
			}
		case utils.ESC:
			slog.Info("editorPrompt, ESC")
			e.editorSetStatusMessage(utils.KILO_DEFAULT_STATUS_MESSAGE)
			return ""
		case utils.ENTER:
			e.editorSetStatusMessage(utils.KILO_DEFAULT_STATUS_MESSAGE)
			return buf
		default:
			if !utils.IsCtrlKey(b) || b < 128 {
				buf += string(rune(b))
			}
		}
	}
}
