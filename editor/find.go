package editor

import (
	"strings"

	"github.com/alcb1310/kilo-go/utils"
)

func (e *EditorConfig) editorFind() {
	cx := e.cx
	cy := e.cy
	colloffset := e.colloffset
	rowoffset := e.rowoffset

	query := e.editorPrompt("Search: ", e.editorFindCallback)
	if query == "" {
		e.cx = cx
		e.cy = cy
		e.colloffset = colloffset
		e.rowoffset = rowoffset
	}

}

func (e *EditorConfig) editorFindCallback(query string, key int) {
	if key == utils.ENTER || key == utils.ESC {
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
