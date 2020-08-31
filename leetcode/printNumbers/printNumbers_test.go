package printNumbers

import (
	"testing"
)

func TestPrintNumbers(t *testing.T) {
	list := printNumbers(1)
	if len(list) != 9 {
		t.Fatal("[*] TestPrintNumbers Failed when n = 1\n")
	}

	list = printNumbers(2)
	if len(list) != 99 {
		t.Fatal("[*] TestPrintNumbers Failed when n = 2\n")
	}
}
