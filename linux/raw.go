package linux

import (
	"fmt"

	"github.com/alcb1310/kilo-go/utils"
	"golang.org/x/sys/unix"
)

type UnixRawMode struct{}

func NewUnixRawMode() *UnixRawMode {
	return &UnixRawMode{}
}

func (r *UnixRawMode) EnableRawMode() (func(), error) {
	termios, err := unix.IoctlGetTermios(unix.Stdin, unix.TCGETS)
	if err != nil {
		return nil, fmt.Errorf("EnableRawMode: error getting terminal flags: %w", err)
	}

	original := *termios

	termios.Lflag &^= unix.ECHO | unix.ICANON | unix.IEXTEN | unix.ISIG
	termios.Iflag &^= unix.BRKINT | unix.ICRNL | unix.INPCK | unix.ISTRIP | unix.IXON
	termios.Oflag &^= unix.OPOST
	termios.Cflag |= unix.CS8

	if err = unix.IoctlSetTermios(unix.Stdin, unix.TCSETS, termios); err != nil {
		return nil, fmt.Errorf("EnableRawMode: error setting terminal flags: %w", err)
	}

	return func() {
		if err = unix.IoctlSetTermios(unix.Stdin, unix.TCSETS, &original); err != nil {
			utils.SafeExit(nil, fmt.Errorf("EnableRawMode: error restoring terminal flags: %w", err))
		}
	}, nil
}
