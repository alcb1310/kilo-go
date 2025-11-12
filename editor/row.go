package editor

import "github.com/alcb1310/kilo-go/utils"

func (e *EditorConfig) editorAppendRow(s string) {
	row := EditorRow{
		chars:  s,
		render: make([]byte, 0),
	}
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
