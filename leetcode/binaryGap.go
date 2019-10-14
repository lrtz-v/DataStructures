package leetcode

func binaryGap(N int) int {

	max := 0
	startIndedx := -1
	for i := 0; ; i++ {
		t := N % 2

		if t == 1 {
			if startIndedx == -1 {
				startIndedx = i
			} else {
				l := i - startIndedx
				if l > max {
					max = l
				}
				startIndedx = i
			}
		}

		if N < 2 {
			break
		}
		N = N / 2
	}

	return max
}
