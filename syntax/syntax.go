package syntax

type EditorSyntax struct {
	Filetype              string
	Filematch             []string
	Flags                 uint
	SingleLineComment     string
	MultiLineCommentStart string
	MultiLineCommentEnd   string
	Keywords              []string
	Types                 []string
}

var HLDB = []EditorSyntax{
	GO_SYNTAX,
	C_SYNTAX,
}
