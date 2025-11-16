package editor

func (e *EditorConfig) editorInsertChar(c byte) {
	if e.cy >= len(e.rows) {
		e.editorAppendRow("")
	}
	e.editorRowInsertChar(&e.rows[e.cy], e.cx, c)
	e.cx++
}

func (e *EditorConfig) editorDeleteChar() {
	if e.cy >= len(e.rows) {
		return
	}
	if e.cx == 0 && e.cy == 0 {
		return
	}

	if e.cx > 0 {
		e.editorRowDeleteChar(&e.rows[e.cy], e.cx)
		e.cx--
	} else {
		e.cx = len(e.rows[e.cy-1].chars)
		e.editorRowAppendString(&e.rows[e.cy-1], e.rows[e.cy].chars)
		e.editorDelRow(e.cy)
		e.cy--
	}
}
