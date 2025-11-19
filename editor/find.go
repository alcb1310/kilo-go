package editor

import (
	"strings"
)

func (e *EditorConfig) editorFind() {
	query := e.editorPrompt("Search: ", nil)
	if query == "" {
		return
	}

	for i := range e.numrows {
		row := &e.rows[i]
		if strings.Contains(row.chars, query) {
			e.cy = i
			e.cx = strings.Index(row.chars, query)
			return
		}
	}
}
