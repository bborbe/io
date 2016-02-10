package buffered_writer

import (
	"io"
)

const BUFFER_SIZE = 100

type bufferedWriter struct {
	done   chan bool
	buffer chan []byte
	writer io.Writer
}

func NewBufferedWriter(writer io.Writer) *bufferedWriter {
	b := new(bufferedWriter)
	b.buffer = make(chan []byte, BUFFER_SIZE)
	b.done = make(chan bool)
	b.writer = writer
	go b.work()
	return b
}

func (b *bufferedWriter) work() {
	for {
		select {
		case <-b.done:
			return
		case p := <-b.buffer:
			b.writer.Write(p)
		}
	}
}

func (b *bufferedWriter) Write(p []byte) (n int, err error) {
	b.buffer <- p
	return len(p), nil
}

func (b *bufferedWriter) Close() error {
	select {
	case b.done <- true:
	default:
	}

	for {
		select {
		case p := <-b.buffer:
			b.writer.Write(p)
		default:
			return nil
		}
	}
	return nil
}
