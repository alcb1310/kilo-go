package editor

func (e *EditorConfig) editorReadKey() (byte, error) {
	b, err := e.reader.ReadByte()

	return b, err
}
