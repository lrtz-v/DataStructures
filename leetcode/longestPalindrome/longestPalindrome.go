package longestPalindrome

/*
给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。

示例 1：
输入: "babad"
输出: "bab"
注意: "aba" 也是一个有效答案。

示例 2：
输入: "cbbd"
输出: "bb"

*/

// 暴力破解
func longestPalindrome(s string) string {

	indexMap := make(map[byte][]int)
	for i := 0; i < len(s); i++ {
		if v, ok := indexMap[s[i]]; ok {
			v = append(v, i)
			indexMap[s[i]] = v
		} else {
			indexMap[s[i]] = []int{i}
		}
	}

	longestSubString := ""
	for i := 0; i < len(s); i++ {
		if _, ok := indexMap[s[i]]; !ok {
			continue
		}
		j := findPalindrome(s, i, indexMap[s[i]])
		if j >= i && (j-i+1) > len(longestSubString) {
			longestSubString = s[i:j+1]
		}
	}

	return longestSubString
}

func findPalindrome(s string, i int, indexSlice []int) int {
	for j := len(indexSlice)-1; j > 0; j-- {
		if i >= indexSlice[j] {
			return 0
		}
		if isPalindrome(i, indexSlice[j], s) {
			return indexSlice[j]
		}
	}
	return 0
}

func isPalindrome(i, j int, s string) bool {
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}
