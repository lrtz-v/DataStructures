package solveNQueens

import (
	"testing"
)


func TestSolveNQueens(t *testing.T) {
	res := solveNQueens(4)
	if len(res) != 2 {
		t.Fatal(res)
	}
}

func TestSolveNQueens2(t *testing.T) {
	solveNQueens(8)
}
