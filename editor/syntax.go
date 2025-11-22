package editor

import "github.com/alcb1310/kilo-go/utils"

func editorUpdateSyntax(row *EditorRow) {
	row.hl = make([]utils.EditorHighlight, len(row.render))

	for i := range len(row.render) {
		if utils.IsDigit(row.render[i]) {
			row.hl[i] = utils.HL_NUMBER
		}
	}
}

func editorSyntaxToColor(hl utils.EditorHighlight) (r uint8, g uint8, b uint8) {
	switch hl {
	case utils.HL_NORMAL:
		return 255, 255, 255
	case utils.HL_NUMBER:
		return 255, 0, 0
	case utils.HL_MATCH:
		return 0, 0, 255
	default:
		return 255, 255, 255
	}
}
