package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/alcb1310/kilo-go/editor"
	"github.com/alcb1310/kilo-go/linux"
	"github.com/alcb1310/kilo-go/utils"
)

var restoreFunc func()

func init() {
	var f *os.File
	var err error
	userTempDir, _ := os.UserConfigDir()
	if f, err = utils.CreateLoggerFile(userTempDir); err != nil {
		utils.SafeExit(nil, err)
	}

	handlerOptions := &slog.HandlerOptions{}
	handlerOptions.Level = slog.LevelDebug

	loggerHandler := slog.NewTextHandler(f, handlerOptions)
	slog.SetDefault(slog.New(loggerHandler))

	u := linux.NewUnixRawMode()
	restoreFunc, err = u.EnableRawMode()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\r\n", err)
		os.Exit(1)
	}
}

func main() {
	editor := editor.NewEditor(restoreFunc)
	editor.EditorLoop()
}
