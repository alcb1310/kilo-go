package appendbuffer

type AppendBuffer struct {
	buf []byte
}

func New() *AppendBuffer {
	return &AppendBuffer{}
}

func (ab *AppendBuffer) Write(p []byte) (int, error) {
	ab.buf = append(ab.buf, p...)
	return len(p), nil
}

func (ab *AppendBuffer) Bytes() []byte {
	return ab.buf
}
