package utils

const (
	ESC   = 0x1b
	ENTER = '\r'

	KILO_VERSION = "0.0.1"

	KILO_DEFAULT_STATUS_MESSAGE = "HELP: Ctrl-S = save | Ctrl-Q = quit | Ctrl-F = find"
	KILO_DECIMAL_SEPARATOR      = '.'
)

var (
	KILO_TAB_STOP   int = 8
	KILO_QUIT_TIMES int = 3

	KILO_DEFAULT_COLOR [3]uint8 = [3]uint8{255, 255, 255}
	KILO_NUMBER_COLOR  [3]uint8 = [3]uint8{255, 0, 0}
	KILO_MATCH_COLOR   [3]uint8 = [3]uint8{51, 255, 0}
	KILO_STRING_COLOR  [3]uint8 = [3]uint8{255, 39, 155}
	KILO_COMMENT_COLOR [3]uint8 = [3]uint8{0, 255, 255}
	KILO_KEYWORD_COLOR [3]uint8 = [3]uint8{255, 239, 0}
	KILO_TYPE_COLOR    [3]uint8 = [3]uint8{126, 239, 55}
)

const (
	BACKSPACE  = 127
	ARROW_LEFT = iota + 1000
	ARROW_RIGHT
	ARROW_UP
	ARROW_DOWN
	DEL_KEY
	HOME_KEY
	END_KEY
	PAGE_UP
	PAGE_DOWN
)

const (
	HL_HIGHLIGHT_NUMBER = (1 << 0)
	HL_HIGHLIGHT_STRING = (1 << 1)
)

type EditorHighlight int

const (
	HL_NORMAL EditorHighlight = iota
	HL_COMMENT
	HL_MLCOMMENT
	HL_KEYWORD
	HL_TYPE_KEYWORD
	HL_STRING
	HL_NUMBER
	HL_MATCH
)

type Callback func(query string, key int)
