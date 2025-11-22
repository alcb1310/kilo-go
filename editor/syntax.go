package editor

import "github.com/alcb1310/kilo-go/utils"

func editorUpdateSyntax(row *EditorRow) {
	prevSep := true
	row.hl = make([]utils.EditorHighlight, len(row.render))

	i := 0
	for i < len(row.render) {
		c := row.render[i]
		var prevHL utils.EditorHighlight
		if i > 0 {
			prevHL = row.hl[i-1]
		} else {
			prevHL = utils.HL_NORMAL
		}

		if utils.IsDigit(c) &&
			(prevSep || prevHL == utils.HL_NUMBER) ||
			(c == utils.KILO_DECIMAL_SEPARATOR && prevHL == utils.HL_NUMBER) {
			row.hl[i] = utils.HL_NUMBER
			i++
			prevSep = false
			continue
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
