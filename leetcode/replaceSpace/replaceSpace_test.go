package replaceSpace

import "testing"

func TestReplaceSpace(t *testing.T) {

	s := replaceSpace("We are happy.")
	if s != "We%20are%20happy." {
		t.Fatalf("[*] replaceSpace Error! Got %s\n", s)
	}

}
