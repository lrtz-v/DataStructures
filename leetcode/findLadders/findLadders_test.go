package findLadders

import (
	"testing"
)

func TestFindLadders(t *testing.T) {
	l := findLadders("hit", "cog", []string{"hot","dot","dog","lot","log","cog"})
	if len(l) != 2 {
		t.Fatal("[*] TestFindLadders Error!")
	}
}

func TestFindLadders2(t *testing.T) {
	l := findLadders("hit", "cog", []string{"hot","dot","dog","lot","log"})
	if len(l) != 1 {
		t.Fatal("[*] TestFindLadders2 Error!")
	}
}

func TestFindLaddersEmpty(t *testing.T) {
	l := findLadders("hit", "cog", []string{})
	if len(l) != 0 {
		t.Fatal("[*] TestFindLaddersEmpty Error!")
	}
	l = findLadders("hit", "cog", []string{"asd", "zxc"})
	if len(l) != 0 {
		t.Fatal("[*] TestFindLaddersEmpty Error!")
	}
}
