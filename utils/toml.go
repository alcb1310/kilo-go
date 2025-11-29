package utils

import (
	"fmt"
	"os"
	"path"

	"github.com/BurntSushi/toml"
)

type Settings struct {
	QuitTimes int `toml:"quit_times"`
	TabStop   int `toml:"tab_stop"`
}

type TomlConfig struct {
	Settings Settings
	Theme    map[string][3]uint8
}

func LoadTOML() error {
	var config TomlConfig = TomlConfig{
		Settings: Settings{
			QuitTimes: -1,
			TabStop:   -1,
		},
	}
	dir, err := os.UserConfigDir()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\r\n", err)
		return err
	}

	if err = os.MkdirAll(path.Join(dir, "kilo"), 0o755); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\r\n", err)
		return err
	}

	filepath := path.Join(dir, "kilo", "config.toml")
	if _, err = os.Stat(filepath); os.IsNotExist(err) {
		return nil
	}

	_, err = toml.DecodeFile(filepath, &config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\r\n", err)
		return err
	}

	return nil
}
