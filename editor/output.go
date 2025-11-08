package editor

import (
	"fmt"
	"os"

	"github.com/alcb1310/kilo-go/utils"
)

func (e *EditorConfig) editorRefreshScreen() {
	fmt.Fprintf(os.Stdout, "%c[2J", utils.ESC)
	fmt.Fprintf(os.Stdout, "%c[H", utils.ESC)
}
