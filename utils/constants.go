package utils

const (
	ESC   = 0x1b
	ENTER = '\r'

	KILO_VERSION = "0.0.1"

	KILO_TAB_STOP               = 8
	KILO_QUIT_TIMES             = 3
	KILO_DEFAULT_STATUS_MESSAGE = "HELP: Ctrl-S = save | Ctrl-Q = quit | Ctrl-F = find"
	KILO_DECIMAL_SEPARATOR      = '.'
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

type EditorHighlight int

const (
	HL_NORMAL EditorHighlight = iota
	HL_NUMBER
	HL_MATCH
)

type Callback func(query string, key int)
