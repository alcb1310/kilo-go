package linux

import (
	"fmt"
	"os"

	"golang.org/x/sys/unix"
)

func EnableRawMode() (func(), error) {
	termios, err := unix.IoctlGetTermios(unix.Stdin, unix.TCGETS)
	if err != nil {
		return nil, fmt.Errorf("EnableRawMode: error getting terminal flags: %w", err)
	}

	original := *termios

	termios.Lflag &^= unix.ECHO | unix.ICANON | unix.ISIG

	if err = unix.IoctlSetTermios(unix.Stdin, unix.TCSETS, termios); err != nil {
		return nil, fmt.Errorf("EnableRawMode: error setting terminal flags: %w", err)
	}

	return func() {
		if err = unix.IoctlSetTermios(unix.Stdin, unix.TCSETS, &original); err != nil {
			fmt.Fprintf(os.Stderr, "EnableRawMode: error restoring terminal flags: %s\n", err)
			os.Exit(1)
		}
	}, nil
}
