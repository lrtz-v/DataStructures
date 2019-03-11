package longestSubstring

type Unit struct{}

func repeat(str *string) bool {
	tmpList := make(map[int32]bool)
	for _, val := range *str {
		if _, ok := tmpList[val]; ok {
			return true
		}
		tmpList[val] = true
	}
	return false
}

func lengthOfLongestSubstring(s string) int {
	length := len(s)
	maxIndex := length
	tmpLen := length
	var tmpStr string
	tmpHash := make(map[string]bool)
	for {
		startIndex := 0
		endIndex := startIndex + tmpLen
		for {
			tmpStr = s[startIndex:endIndex]
			if _, ok := tmpHash[tmpStr]; !ok {
				if !repeat(&tmpStr) {
					return len(tmpStr)
				}
			}
			tmpHash[tmpStr] = false

			startIndex++
			endIndex++

			if endIndex > maxIndex {
				break
			}
		}

		tmpLen--
		if tmpLen == 1 {
			return 1
		}
	}
}

func lengthOfLongestSubstring3(s string) int {
	length := len(s)
	if length == 0 {
		return 0
	}

	maxLen := 1
	startIndex := 0
	endIndex := 0
	tmpMaxLen := 0
	tmpHash := make(map[rune]Unit)
	for {
		tmpStr := rune(s[endIndex : endIndex+1][0])
		if _, ok := tmpHash[tmpStr]; ok {
			if tmpMaxLen > maxLen {
				maxLen = tmpMaxLen
			}
			startIndex++
			endIndex = startIndex
			tmpHash = nil
			tmpHash = make(map[int32]Unit)
			tmpMaxLen = 0
			continue
		} else {
			tmpHash[tmpStr] = Unit{}
			tmpMaxLen++
		}

		endIndex++

		if endIndex+1 > length {
			if tmpMaxLen > maxLen {
				maxLen = tmpMaxLen
			}
			break
		}

		if length-startIndex <= maxLen {
			break
		}
	}
	return maxLen
}
