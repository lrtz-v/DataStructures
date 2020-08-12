package movingcount

func movingCount(m int, n int, k int) int {
	if (k == 0) {
		return 1
	}
	
	// init
	table := make([][]int, m)
	for i := 0; i < m; i++ {
		table[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		v1 := calculate(i)
		for j := 0; j < n; j++ {
			v2 := calculate(j)
			if (i == 0 && j == 0 || j > 0 && table[i][j - 1] == 1 || i > 0 && table[i-1][j] == 1) && v1 + v2 <= k {
				table[i][j] = 1
			}
		}
	}

	count := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if table[i][j] == 1 {
				count++
			}
		}
	}

    return count
}

func calculate(a int) int {
    sum := 0
	for a > 0 {
		sum += a % 10
		a = a / 10
	}

	return sum
}
