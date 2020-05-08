package maximalsquare

import (
	"testing"
)

func TestMaximalSquare(t *testing.T) {
	matrix := [][]int {
		{1, 0, 1, 0, 0},
		{1, 0, 1, 1, 1},
		{1, 1, 1, 1, 1},
		{1, 0, 0, 1, 0},
	}
	if area := maximalSquare(matrix); area != 4 {
		t.Fatalf("[*]ERROR Expected 4, Got %d\n", area)
	}
}

func TestMaximalSquare2(t *testing.T) {
	matrix := [][]int {
		{1, 1, 1, 1},
		{1, 1, 1, 1},
		{1, 1, 1, 1},
	}
	if area := maximalSquare(matrix); area != 9 {
		t.Fatalf("[*]ERROR Expected 9, Got %d\n", area)
	}
}

func TestMaximalSquare3(t *testing.T) {
	matrix := [][]int {
		{1 , 0 , 1 , 0 , 0 , 1 , 1 , 1 , 0},
		{1 , 1 , 1 , 0 , 0 , 0 , 0 , 0 , 1},
		{0 , 0 , 1 , 1 , 0 , 0 , 0 , 1 , 1},
		{0 , 1 , 1 , 0 , 0 , 1 , 0 , 0 , 1},
		{1 , 1 , 0 , 1 , 1 , 0 , 0 , 1 , 0},
		{0 , 1 , 1 , 1 , 1 , 1 , 1 , 0 , 1},
		{1 , 0 , 1 , 1 , 1 , 0 , 0 , 1 , 0},
		{1 , 1 , 1 , 0 , 1 , 0 , 0 , 0 , 1},
		{0 , 1 , 1 , 1 , 1 , 0 , 0 , 1 , 0},
		{1 , 0 , 0 , 1 , 1 , 1 , 0 , 0 , 0},
	}
	if area := maximalSquare(matrix); area != 4 {
		t.Fatalf("[*]ERROR Expected 4, Got %d\n", area)
	}
}
