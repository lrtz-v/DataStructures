package exist

import "testing"

func TestExist(t *testing.T) {
	data := make([][]byte, 3)
	data[0] = []byte{'a', 'b', 'c', 'e'}
	data[1] = []byte{'s', 'f', 'c', 's'}
	data[2] = []byte{'a', 'd', 'e', 'e'}
	if !(exist(data, "bfce")) {
		t.Fatal("TestExist Failed!")
	}
}
