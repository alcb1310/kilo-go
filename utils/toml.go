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

	if config.Settings.QuitTimes >= 0 {
		KILO_QUIT_TIMES = config.Settings.QuitTimes
	}
	if config.Settings.TabStop >= 0 {
		KILO_TAB_STOP = config.Settings.TabStop
	}

	var val [3]uint8
	var ok bool

	if val, ok = config.Theme["default"]; ok {
		KILO_DEFAULT_COLOR = val
	}
	if val, ok = config.Theme["number"]; ok {
		KILO_NUMBER_COLOR = val
	}
	if val, ok = config.Theme["match"]; ok {
		KILO_MATCH_COLOR = val
	}
	if val, ok = config.Theme["string"]; ok {
		KILO_STRING_COLOR = val
	}
	if val, ok = config.Theme["comment"]; ok {
		KILO_COMMENT_COLOR = val
	}
	if val, ok = config.Theme["keyword"]; ok {
		KILO_KEYWORD_COLOR = val
	}
	if val, ok = config.Theme["type"]; ok {
		KILO_TYPE_COLOR = val
	}

	return nil
}
