package mock

import (
	"io"
	"testing"

	. "github.com/bborbe/assert"
)

func TestImplementsWriter(t *testing.T) {
	writer := NewWriteCloseFlusherMock()
	var i *io.Writer
	err := AssertThat(writer, Implements(i))
	if err != nil {
		t.Fatal(err)
	}
}

func TestBytes(t *testing.T) {
	var err error
	writer := NewWriteCloseFlusherMock()
	err = AssertThat(string(writer.Bytes()), Is(""))
	if err != nil {
		t.Fatal(err)
	}
	writer.Write([]byte("hello world"))
	err = AssertThat(string(writer.Bytes()), Is("hello world"))
	if err != nil {
		t.Fatal(err)
	}
}
func TestReset(t *testing.T) {
	var err error
	writer := NewWriteCloseFlusherMock()
	writer.Write([]byte("hello world"))
	err = AssertThat(string(writer.Bytes()), Is("hello world"))
	if err != nil {
		t.Fatal(err)
	}
	writer.Reset()
	err = AssertThat(string(writer.Bytes()), Is(""))
	if err != nil {
		t.Fatal(err)
	}
}
