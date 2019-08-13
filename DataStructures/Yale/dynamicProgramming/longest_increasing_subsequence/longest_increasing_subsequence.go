package lis

func longestIncreasingSubsequence(a []int, n int) int {
	if n == 0 {
		return 0
	}

	lisTable := make([]int, n)

	for i := 0; i < n; i++ {
		lisTable[i] = 1

		for j := 0; j < i; j++ {
			if a[j] < a[i] && lisTable[j]+1 > lisTable[i] {
				lisTable[i] = lisTable[j] + 1
			}
		}
	}

	max := 0

	for i := 0; i < n; i++ {
		if lisTable[i] > max {
			max = lisTable[i]
		}
	}

	return max
}

func longestIncreasingSubsequence2(a []int, n int) int {
	max := 0

	if n == 0 {
		return max
	}

	lisTable := make([]int, n)

	for i := 0; i < n; i++ {
		lisTable[i] = 1
		if i == 0 {
			continue
		}
		if a[i-1] < a[i] {
			lisTable[i] += lisTable[i-1]
			if lisTable[i] > max {
				max = lisTable[i]
			}
		}
	}

	return max
}
