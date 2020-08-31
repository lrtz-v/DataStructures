package printNumbers

import "math"

func printNumbers(n int) []int {
	size := int(math.Pow10(n) - 1)
	l := make([]int, size)
	for i := 1; i <= size; i++ {
		l[i-1] = i
	}
	return l
}
