package editor

import "github.com/alcb1310/kilo-go/utils"

func (e *EditorConfig) editorAppendRow(s string) {
	row := EditorRow{
		chars:  s,
		render: make([]byte, len(s)),
	}
	e.cx = 0
	e.editorUpdateRow(&row)
	e.rows = append(e.rows, row)
	e.numrows++
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
	row.render = make([]byte, len(row.chars)+1)
	row.chars = row.chars[:at] + string(c) + row.chars[at:]
	e.editorUpdateRow(row)
}
