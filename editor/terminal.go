package editor

import (
	"github.com/alcb1310/kilo-go/utils"
)

func (e *EditorConfig) editorReadKey() (byte, error) {
	b, err := e.reader.ReadByte()

	if b == utils.ESC {
		seq := make([]byte, 2)

		seq[0], err = e.reader.ReadByte()
		if err != nil {
			return utils.ESC, nil
		}
		seq[1], err = e.reader.ReadByte()
		if err != nil {
			return utils.ESC, nil
		}

		if seq[0] == '[' {
			switch seq[1] {
			case 'A':
				return 'k', nil
			case 'B':
				return 'j', nil
			case 'C':
				return 'l', nil
			case 'D':
				return 'h', nil
			}
		}

		return utils.ESC, nil
	}

	return b, err
}
