package file_writer

import (
	"bufio"
	"github.com/bborbe/io"
	"os"
)

type FileWriter interface {
	io.WriteCloseFlusher
}

type fileWriter struct {
	writer io.WriteFlusher
	file   *os.File
}

func NewFileWriter(fileName string) (*fileWriter, error) {
	file, err := os.Create(fileName)
	if err != nil {
		return nil, err
	}

	f := new(fileWriter)
	f.file = file
	f.writer = bufio.NewWriter(file)
	return f, nil
}

func (f *fileWriter) Write(p []byte) (int, error) {
	return f.writer.Write(p)
}

func (f *fileWriter) Close() error {
	f.Flush()
	return f.file.Close()
}

func (f *fileWriter) Flush() error {
	return f.writer.Flush()
}