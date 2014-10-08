package io

import "io"

type writer struct {
	content []byte
}

type WriterContent interface {
	io.Writer
	Content() []byte
}

func NewWriter() *writer {
	w := new(writer)
	w.content = make([]byte, 0)
	return w
}

func (w *writer) Write(p []byte) (int, error) {
	w.content = append(w.content, p...)
	return len(p), nil
}

func (w *writer) Content() []byte {
	return w.content
}