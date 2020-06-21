package divisorGame

import (
	"fmt"
	"testing"
)

func TestDivisorGame(t *testing.T) {
	N := 111
	for N > 0 {
		l := getL(N)
		fmt.Println(l)
		if len(l) <= 0 {
			break
		}
		N -= l[len(l)-1]
	}
	
	t.Fatal()
}

func getL(N int) []int {
	l := []int{}
	if N <= 1 {
		return l
	}

	for i := 1; i < N - 1; i++ {
		if N % i == 0 {
			l = append(l, i)
		}
	}
	return l
}
