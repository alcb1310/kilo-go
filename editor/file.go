package editor

func (e *EditorConfig) editorOpen() {
	line := "Hello, world!"
	e.row.chars = line
	e.numrows = 1
}
