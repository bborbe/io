package file_writer

import (
	. "github.com/bborbe/assert"
	"testing"
)

func TestImplementsWriter(t *testing.T) {
	b, _ := NewFileWriter("")
	var expected *FileWriter
	err := AssertThat(b, Implements(expected))
	if err != nil {
		t.Fatal(err)
	}
}
