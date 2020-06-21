package isSubsequence

// func isSubsequence(s string, t string) bool {
// 	indexMap := make(map[byte][]int, 26)
// 	for i := 0; i < len(t); i++ {
// 		if m, ok := indexMap[t[i]]; ok {
// 			indexMap[t[i]] = append(m, i)
// 		}else {
// 			indexMap[t[i]] = []int{i}
// 		}
// 	}

// 	offset := -1
// 	for i := 0; i < len(s); i++ {
// 		o := find(indexMap[s[i]], offset)
// 		if o <= offset {
// 			return false
// 		}
// 		offset = o
// 	}

// 	return true
// }

// func find(index []int, offset int) int {
// 	if index == nil {
// 		return offset
// 	} else if len(index) == 1 {
// 		return index[0]
// 	} else if index[len(index)-1] <= offset {
// 		return offset
// 	}

// 	start := 0
// 	end := len(index) - 1
// 	for start <= end {
// 		mid := start + (end - start) / 2
// 		if index[mid] == offset {
// 			return index[mid + 1]
// 		} else if index[mid] > offset {
// 			if mid == 0 || index[mid - 1] <= offset {
// 				return index[mid]
// 			}
// 			end = mid - 1
//         } else {
// 			start = mid + 1
// 		}
// 	}

// 	return -1
// }

// DP
func isSubsequence(s string, t string) bool {

	t = " " + t
	n := len(t)

	dp := make(map[int][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 26)
	}

	for ch := 0; ch < 26; ch++ {
		p := -1
		for i := n-1; i >= 0; i-- {
			dp[i][ch] = p
			if t[i] == byte('a' + ch) {
				p = i
			}
		}
	}

	i := 0
	for _, ch := range s {
		i = dp[i][ch - 'a']
		if (i == -1) {
			return false
		}
	}
	return true
}
