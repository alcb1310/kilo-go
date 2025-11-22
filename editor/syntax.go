package editor

import (
	"strings"

	"github.com/alcb1310/kilo-go/utils"
)

func (e *EditorConfig) editorUpdateSyntax(row *EditorRow) {
	prevSep := true
	row.hl = make([]utils.EditorHighlight, len(row.render))

	if e.syntax == nil {
		return
	}

	i := 0
	for i < len(row.render) {
		c := row.render[i]
		var prevHL utils.EditorHighlight
		if i > 0 {
			prevHL = row.hl[i-1]
		} else {
			prevHL = utils.HL_NORMAL
		}

		if e.syntax.flags&utils.HL_HIGHLIGHT_NUMBER == 1 {
			if utils.IsDigit(c) &&
				(prevSep || prevHL == utils.HL_NUMBER) ||
				(c == utils.KILO_DECIMAL_SEPARATOR && prevHL == utils.HL_NUMBER) {
				row.hl[i] = utils.HL_NUMBER
				i++
				prevSep = false
				continue
			}
		}

		prevSep = utils.IsSeparator(c)
		i++
	}
}

func editorSyntaxToColor(hl utils.EditorHighlight) (r uint8, g uint8, b uint8) {
	switch hl {
	case utils.HL_NORMAL:
		return 255, 255, 255
	case utils.HL_NUMBER:
		return 255, 0, 0
	case utils.HL_MATCH:
		return 51, 255, 0
	default:
		return 255, 255, 255
	}
}

func (e *EditorConfig) editorSelectSyntaxHighlight() {
	e.syntax = nil
	if e.filename == "" {
		return
	}

	lastIndex := strings.LastIndex(e.filename, ".")
	if lastIndex == -1 {
		return
	}
	ext := e.filename[lastIndex:]

	for i, s := range HLDB {
		isExt := s.filematch[i][0] == '.'

		if (isExt && ext == s.filematch[i]) ||
			(!isExt && strings.Contains(ext, s.filematch[i])) {
			e.syntax = &s

			for _, row := range e.rows {
				e.editorUpdateSyntax(&row)
			}

			return
		}
	}
}
