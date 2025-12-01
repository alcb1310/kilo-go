package syntax

import "github.com/alcb1310/kilo-go/utils"

var C_HL_EXTENSIONS []string = []string{".c", ".h", ".cpp"}

var C_SYNTAX = EditorSyntax{
	Filetype:              "c",
	Filematch:             C_HL_EXTENSIONS,
	Flags:                 utils.HL_HIGHLIGHT_NUMBER | utils.HL_HIGHLIGHT_STRING,
	SingleLineComment:     "//",
	MultiLineCommentStart: "/*",
	MultiLineCommentEnd:   "*/",
	Keywords: []string{"switch", "if", "while", "for", "break", "continue", "return", "else",
		"struct", "union", "typedef", "static", "enum", "class", "case", "#include", "#define", "#ifndef", "#ifdef", "#endif", "#else"},
	Types: []string{"int", "long", "double", "float", "char", "unsigned", "signed",
		"void"},
}
