package utils

func CtrlKey(key byte) int {
	return int(key & 0x1f)
}

func IsCtrlKey(key int) bool {
	return key <= 0x1f || key == BACKSPACE
}
