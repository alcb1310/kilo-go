package utils

import "strings"

func IsDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func IsSeparator(c byte) bool {
	return IsSpace(c) || strings.ContainsRune(":,.()+-/*=~%<>[];{}", rune(c)) || rune(c) == 0
}

func IsSpace(c byte) bool {
	return c == ' ' || c == '\t' || c == '\n' || c == '\r'
}
