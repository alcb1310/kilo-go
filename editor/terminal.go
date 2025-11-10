package editor

import (
	"github.com/alcb1310/kilo-go/utils"
)

func (e *EditorConfig) editorReadKey() (int, error) {
	b, err := e.reader.ReadByte()

	if b == utils.ESC {
		seq := make([]byte, 3)

		seq[0], err = e.reader.ReadByte()
		if err != nil {
			return utils.ESC, nil
		}
		seq[1], err = e.reader.ReadByte()
		if err != nil {
			return utils.ESC, nil
		}

		if seq[0] == '[' {
			if seq[1] >= '0' && seq[1] <= '9' {
				seq[2], err = e.reader.ReadByte()
				if err != nil {
					return utils.ESC, nil
				}

				if seq[2] == '~' {
					switch seq[1] {
					case '1':
						return utils.HOME_KEY, nil
					case '4':
						return utils.END_KEY, nil
					case '5':
						return utils.PAGE_UP, nil
					case '6':
						return utils.PAGE_DOWN, nil
					case '7':
						return utils.HOME_KEY, nil
					case '8':
						return utils.END_KEY, nil
					}
				}
			} else {
				switch seq[1] {
				case 'A':
					return utils.ARROW_UP, nil
				case 'B':
					return utils.ARROW_DOWN, nil
				case 'C':
					return utils.ARROW_RIGHT, nil
				case 'D':
					return utils.ARROW_LEFT, nil
				case 'H':
					return utils.HOME_KEY, nil
				case 'F':
					return utils.END_KEY, nil
				}
			}
		} else if seq[0] == 'O' {
			switch seq[1] {
			case 'H':
				return utils.HOME_KEY, nil
			case 'F':
				return utils.END_KEY, nil
			}
		}

		return utils.ESC, nil
	}

	return int(b), err
}
