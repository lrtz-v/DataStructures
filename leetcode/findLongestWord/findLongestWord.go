package findLongestWord

func findLongestWord(s string, d []string) string {
	longestWord := ""
	for _, des := range d {
		if find(s, des) {
			if len(des) > len(longestWord) {
				longestWord = des
			} else if len(des) == len(longestWord) {
				if des < longestWord {
					longestWord = des
				}
			}
		}
	}
	return longestWord
}

func find(src, des string) bool {
	j := 0
	for i := 0; i < len(src) && j < len(des); i++ {
		if src[i] == des[j] {
			j++
		}
	}
	return j == len(des)
}
