package utils

func CtrlKey(key byte) byte {
	return key & 0x1f
}
