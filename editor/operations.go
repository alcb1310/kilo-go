package editor

func (e *EditorConfig) editorInsertChar(c byte) {
	if e.cy >= len(e.rows) {
		e.editorAppendRow("")
	}
	e.editorRowInsertChar(&e.rows[e.cy], e.cx, c)
	e.cx++
}
