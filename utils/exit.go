package utils

import (
	"fmt"
	"os"
)

// SafeExit is a function that allows us to safely exit the program
//
// # It will call the provided function and exit with the provided error
// if no error is provided, it will exit with 0
//
// @param f - The function to call
// @param err - The error to exit with
func SafeExit(f func(), err error) {
	fmt.Fprintf(os.Stdout, "%c[2J", ESC)
	fmt.Fprintf(os.Stdout, "%c[H", ESC)

	if f != nil {
		f()
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\r\n", err)
		os.Exit(1)
	}

	os.Exit(0)
}
