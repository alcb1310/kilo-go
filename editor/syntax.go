package editor

import (
	"strings"

	"github.com/alcb1310/kilo-go/utils"
)

func (e *EditorConfig) editorUpdateSyntax(row *EditorRow) {
	row.hl = make([]utils.EditorHighlight, len(row.render))
	if e.syntax == nil {
		return
	}

	prevSep := true
	var inString byte = 0
	scs := e.syntax.singleLineComment
	keywords := e.syntax.keywords
	types := e.syntax.types

	i := 0
	for i < len(row.render) {
		c := row.render[i]
		var prevHL utils.EditorHighlight
		if i > 0 {
			prevHL = row.hl[i-1]
		} else {
			prevHL = utils.HL_NORMAL
		}

		if len(scs) > 0 && inString == 0 {
			if strings.HasPrefix(string(row.render[i:]), scs) {
				for j := i; j < len(row.render); j++ {
					row.hl[j] = utils.HL_COMMENT
				}
				break
			}
		}

		if e.syntax.flags&utils.HL_HIGHLIGHT_STRING == 2 {
			if inString != 0 {
				row.hl[i] = utils.HL_STRING
				if c == '\\' && i+1 < len(row.render) {
					i++
					row.hl[i] = utils.HL_STRING
					i++
					continue
				}
				if c == inString {
					inString = 0
				}
				i++
				prevSep = true
				continue
			} else {
				if c == '"' || c == '\'' {
					inString = c
					row.hl[i] = utils.HL_STRING
					i++
					continue
				}
			}
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

		if prevSep {
			j := 0
			for j = 0; j < len(keywords); j++ {
				key := keywords[j]
				if strings.HasPrefix(string(row.render[i:]), key) &&
					((i+len(key) < len(row.render) &&
						utils.IsSeparator(row.render[i+len(key)])) ||
						i+len(key) == len(row.render)) {
					for k := range key {
						row.hl[i+k] = utils.HL_KEYWORD
					}
					i += len(key) - 1
					break
				}
			}

			if j < len(keywords) {
				prevSep = false
				continue
			}

			m := 0
			for m = 0; m < len(types); m++ {
				key := types[m]
				if strings.HasPrefix(string(row.render[i:]), key) &&
					((i+len(key) < len(row.render) &&
						utils.IsSeparator(row.render[i+len(key)])) ||
						i+len(key) == len(row.render)) {
					for k := range key {
						row.hl[i+k] = utils.HL_TYPE_KEYWORD
					}
					i += len(key) - 1
					break
				}
			}

			if m < len(types) {
				prevSep = false
				continue
			}
		}

		prevSep = utils.IsSeparator(c)
		i++
	}
}

func editorSyntaxToColor(hl utils.EditorHighlight) (r uint8, g uint8, b uint8) {
	r = 255
	g = 255
	b = 255

	switch hl {
	case utils.HL_NUMBER:
		g = 0
		b = 0
	case utils.HL_MATCH:
		r = 51
		b = 0
	case utils.HL_STRING:
		g = 39
		b = 155
	case utils.HL_COMMENT:
		r = 0
	case utils.HL_KEYWORD:
		g = 239
		b = 0
	case utils.HL_TYPE_KEYWORD:
		g = 55
		b = 239
		r = 126
	}
	return
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
