package linux

import (
	"fmt"
	"os"

	"github.com/alcb1310/kilo-go/utils"
	"golang.org/x/term"
)

func EnableRawMode() (func(), error) {
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))

	return func() {
		if err = term.Restore(int(os.Stdin.Fd()), oldState); err != nil {
			utils.SafeExit(nil, fmt.Errorf("EnableRawMode: error restoring terminal flags: %w", err))
		}
	}, err
}
