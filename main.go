package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/alcb1310/kilo-go/linux"
	"github.com/alcb1310/kilo-go/utils"
)

type RawMode interface {
	EnableRawMode() (func(), error)
}

type EditorConfig struct {
	restoreFunc func()
}

var editorState EditorConfig = EditorConfig{}

func init() {
	var err error
	u := linux.NewUnixRawMode()
	editorState.restoreFunc, err = u.EnableRawMode()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\r\n", err)
		os.Exit(1)
	}
}

func main() {
	defer utils.SafeExit(editorState.restoreFunc, nil)

	r := bufio.NewReader(os.Stdin)

	for {
		b, err := r.ReadByte()
		if err == io.EOF {
			break
		} else if err != nil {
			utils.SafeExit(editorState.restoreFunc, err)
		}

		if b <= 0x1f || b == 0x7f { // This will make sure we've passed a control-key combo
			fmt.Fprintf(os.Stdout, "%d\r\n", b)
		} else {
			fmt.Fprintf(os.Stdout, "%d (%c)\r\n", b, b)
		}

		if b == 'q' {
			break
		}
	}
}
