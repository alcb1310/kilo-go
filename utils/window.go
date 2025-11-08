package utils

import (
	"fmt"
	"os"

	"golang.org/x/sys/unix"
)

func GetWindowSize() (rows int, cols int, err error) {
	ws, err := unix.IoctlGetWinsize(unix.Stdin, unix.TIOCGWINSZ)
	if err != nil {
		fmt.Fprintf(os.Stderr, "getWindowSize: Error getting window size: %v\r\n", err)
		return
	}

	rows = int(ws.Row)
	cols = int(ws.Col)

	return
}
