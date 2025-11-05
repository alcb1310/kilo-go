package linux

import (
	"fmt"

	"golang.org/x/sys/unix"
)

func EnableRawMode() error {
	termios, err := unix.IoctlGetTermios(unix.Stdin, unix.TCGETS)
	if err != nil {
		return fmt.Errorf("EnableRawMode: error getting terminal flags: %w", err)
	}

	termios.Lflag &^= unix.ECHO

	if err = unix.IoctlSetTermios(unix.Stdin, unix.TCSETS, termios); err != nil {
		return fmt.Errorf("EnableRawMode: error setting terminal flags: %w", err)
	}

	return nil
}
