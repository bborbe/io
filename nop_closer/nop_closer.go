package nop_closer

import "io"

type nopCloser struct {
	io.Reader
}

func (nopCloser) Close() error {
	return nil
}

func New(r io.Reader) io.ReadCloser {
	return nopCloser{r}
}
