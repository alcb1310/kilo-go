package utils

func CtrlKey(key byte) int {
	return int(key & 0x1f)
}
