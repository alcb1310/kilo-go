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
	e.editorRowDeleteChar(&e.rows[e.cy], e.cx)
	if e.cx > 0 {
		e.cx--
	}
}
