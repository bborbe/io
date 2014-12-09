package mock

import (
	"bytes"
)

type writeCloseFlusherMock struct {
	buffer *bytes.Buffer
}

func NewWriteCloseFlusherMock() *writeCloseFlusherMock {
	w := new(writeCloseFlusherMock)
	w.Reset()
	return w
}

func (w *writeCloseFlusherMock) Write(p []byte) (int, error) {
	return w.buffer.Write(p)
}

func (w *writeCloseFlusherMock) Close() error { return nil }

func (w *writeCloseFlusherMock) Flush() error { return nil }

func (b *writeCloseFlusherMock) Bytes() []byte { return b.buffer.Bytes() }

func (b *writeCloseFlusherMock) Reset() { b.buffer = bytes.NewBufferString("") }