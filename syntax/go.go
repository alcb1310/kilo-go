package syntax

import "github.com/alcb1310/kilo-go/utils"

var GO_HL_EXTENSIONS = []string{".go"}
var GO_HL_KEYWORDS = []string{"package", "import", "func", "type", "var", "const", "if", "else",
	"switch", "case", "default", "for", "range", "goto", "continue", "select", "return", "break",
}

var GO_HL_TYPES = []string{"bool", "byte", "error", "float32", "float64", "int", "int16", "int32",
	"int64", "int8", "rune", "string", "uint", "uint16", "uint32", "uint64", "uint8",
}

var GO_SYNTAX = EditorSyntax{
	Filetype:              "go",
	Filematch:             GO_HL_EXTENSIONS,
	Flags:                 utils.HL_HIGHLIGHT_NUMBER | utils.HL_HIGHLIGHT_STRING,
	SingleLineComment:     "//",
	MultiLineCommentStart: "/*",
	MultiLineCommentEnd:   "*/",
	Keywords:              GO_HL_KEYWORDS,
	Types:                 GO_HL_TYPES,
}
