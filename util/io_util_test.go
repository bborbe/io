package util

import (
	"testing"

	. "github.com/bborbe/assert"
)

func TestIsDirectory(t *testing.T) {
	err := AssertThat(IsDirectory("/tmp"), NilValue())
	if err != nil {
		t.Fatal(err)
	}
}

func TestNormalizePath(t *testing.T) {
	dir, err := NormalizePath("/tmp")
	if err := AssertThat(err, NilValue()); err != nil {
		t.Fatal(err)
	}
	if err := AssertThat(dir, Is("/tmp")); err != nil {
		t.Fatal(err)
	}
}
