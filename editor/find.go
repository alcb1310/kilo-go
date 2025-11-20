package editor

import (
	"strings"

	"github.com/alcb1310/kilo-go/utils"
)

// lastMatch is the index of the row that the last match was on
// or -1 if there was no match
var lastMatch int = -1

// direction will store the direction of the search:
// 1 for forward, -1 for backward
var direction int = 1

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
		lastMatch = -1
		direction = 1
		return
	}

	switch key {
	case utils.ARROW_DOWN, utils.ARROW_RIGHT:
		direction = 1
	case utils.ARROW_UP, utils.ARROW_LEFT:
		direction = -1
	default:
		lastMatch = -1
		direction = 1
	}

	if lastMatch == -1 {
		direction = 1
	}
	current := lastMatch
	for range e.numrows {
		current += direction
		switch current {
		case -1:
			current = e.numrows - 1
		case e.numrows:
			current = 0
		}

		row := &e.rows[current]
		if strings.Contains(row.chars, query) {
			lastMatch = current
			e.cy = current
			e.cx = strings.Index(row.chars, query)
			e.rowoffset = e.numrows
			return
		}
	}
}
