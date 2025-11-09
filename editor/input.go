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
	case 'h', 'j', 'k', 'l':
		e.editorMoveCursor(b)
	}
}

func (e *EditorConfig) editorMoveCursor(key byte) {
	switch key {
	case 'h':
		e.cx--
	case 'j':
		e.cy++
	case 'k':
		e.cy--
	case 'l':
		e.cx++
	}
}
