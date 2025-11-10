package editor

import "github.com/alcb1310/kilo-go/utils"

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
	}
}

func (e *EditorConfig) editorMoveCursor(key byte) {
	switch key {
	case utils.ARROW_LEFT:
		e.cx--
	case utils.ARROW_DOWN:
		e.cy++
	case utils.ARROW_UP:
		e.cy--
	case utils.ARROW_RIGHT:
		e.cx++
	}
}
