package utils

import (
	"os"

	"golang.org/x/term"
)

func GetWindowSize() (int, int, error) {
	return term.GetSize(int(os.Stdout.Fd()))
}
