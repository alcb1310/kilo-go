package syntax

import (
	"fmt"
	"os"
	"path"

	"github.com/BurntSushi/toml"
	"github.com/alcb1310/kilo-go/utils"
)

type EditorSyntax struct {
	Filetype              string   `toml:"filetype"`
	Filematch             []string `toml:"extensions"`
	Flags                 uint
	SingleLineComment     string `toml:"slc"`
	MultiLineCommentStart string `toml:"mlcs"`
	MultiLineCommentEnd   string `toml:"mlce"`
	Keywords              []string
	Types                 []string
	NumberFlag            bool `toml:"number"`
	StringFlag            bool `toml:"string"`
}

var HLDB = []EditorSyntax{}

func LoadSyntax() error {
	dir, err := os.UserConfigDir()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\r\n", err)
		return err
	}

	tomlDir := path.Join(path.Join(dir, "kilo"), "highlight")
	_, err = os.Stat(tomlDir)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}

		fmt.Fprintf(os.Stderr, "Error: %s\r\n", err)
		return err
	}

	files, err := os.ReadDir(tomlDir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\r\n", err)
		return err
	}

	for _, file := range files {
		s := EditorSyntax{}
		if file.IsDir() {
			continue
		}

		_, err := toml.DecodeFile(path.Join(tomlDir, file.Name()), &s)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s\r\n", err)
			return err
		}

		if s.NumberFlag {
			s.Flags |= utils.HL_HIGHLIGHT_NUMBER
		}

		if s.StringFlag {
			s.Flags |= utils.HL_HIGHLIGHT_STRING
		}

		HLDB = append(HLDB, s)
	}

	return nil
}
