package editor

import (
	"github.com/alcb1310/kilo-go/utils"
)

func (e *EditorConfig) editorAppendRow(s string) {
	row := EditorRow{
		chars:  s,
		render: make([]byte, 0),
	}
	e.cx = 0
	e.editorUpdateRow(&row)
	e.rows = append(e.rows, row)
	e.numrows++
	e.isDirty = true
}

func (e *EditorConfig) editorUpdateRow(row *EditorRow) {
	for j := 0; j < len(row.chars); j++ {
		if row.chars[j] == '\t' {
			for range utils.KILO_TAB_STOP {
				row.render = append(row.render, ' ')
			}
		} else {
			row.render = append(row.render, row.chars[j])
		}
	}
	row.idx = e.numrows

	e.editorUpdateSyntax(row)
}

func editorRowCxToRx(row *EditorRow, cx int) int {
	rx := 0
	for j := range cx {
		if row.chars[j] == '\t' {
			rx += utils.KILO_TAB_STOP
		} else {
			rx++
		}
	}
	return rx
}

func (e *EditorConfig) editorRowInsertChar(row *EditorRow, at int, c byte) {
	row.render = make([]byte, 0)
	row.chars = row.chars[:at] + string(c) + row.chars[at:]
	e.editorUpdateRow(row)
	e.isDirty = true
}

func (e *EditorConfig) editorRowDeleteChar(row *EditorRow, at int) {
	if at < 0 || at >= len(row.chars) {
		return
	}
	row.render = make([]byte, 0)
	row.chars = row.chars[:at] + row.chars[at+1:]

	for j := at; j <= e.numrows-1; j++ {
		e.rows[j].idx--
	}
	e.editorUpdateRow(row)
	e.isDirty = true
}

func (e *EditorConfig) editorDelRow(at int) {
	if at < 0 || at >= len(e.rows) {
		return
	}
	e.rows = append(e.rows[:at], e.rows[at+1:]...)
	e.numrows--
	e.isDirty = true
}

func (e *EditorConfig) editorRowAppendString(row *EditorRow, s string) {
	row.chars += s
	row.render = make([]byte, 0)
	e.editorUpdateRow(row)
	e.isDirty = true
}

func (e *EditorConfig) editorInsertRow(at int, s string) {
	e.rows = append(e.rows[:at], append([]EditorRow{{chars: s}}, e.rows[at:]...)...)
	for j := at + 1; j <= e.numrows; j++ {
		e.rows[j].idx++
	}

	e.rows[at].idx = at

	e.rows[at].hl = nil
	e.rows[at].hlOpenComment = false
	e.numrows++
	e.isDirty = true
}
