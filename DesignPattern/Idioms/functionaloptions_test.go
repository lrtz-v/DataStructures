package idioms

import (
	"testing"
)

func TestOpenFile(t *testing.T) {
	err := Openfile("/tmp/empty.txt")
	if err != nil {
		panic(err)
	}

	err = Openfile("/tmp/file.txt", UID(1000), Contents("Lorem Ipsum Dolor Amet"))
	if err != nil {
		panic(err)
	}
}
