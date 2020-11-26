package tries

import (
	"testing"
)

func TestTrie(t *testing.T) {
	// t.Skip()

	tree := NewTrie()

	tree.TrieInsertKey("test")

	res := tree.TrieSearchKey("test")
	if !res {
		t.Fatal("laji")
	}

	tree.TrieDeleteKey("test")

	res = tree.TrieSearchKey("test")
	if res {
		t.Fatal("laji")
	}

	tree = nil
}
