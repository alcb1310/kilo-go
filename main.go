package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/alcb1310/kilo-go/linux"
)

func main() {
	restoreFunc, err := linux.EnableRawMode()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}
	defer restoreFunc()

	r := bufio.NewReader(os.Stdin)

	for {
		b, err := r.ReadByte()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "Error: reading key from Stdin: %s\n", err)
			os.Exit(1)
		}

		if b <= 0x1f || b == 0x7f { // This will make sure we've passed a control-key combo
			fmt.Fprintf(os.Stdout, "%d\n", b)
		} else {
			fmt.Fprintf(os.Stdout, "%d (%c)\n", b, b)
		}

		if b == 'q' {
			break
		}
	}
}
