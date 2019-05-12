package leetcode

type unit struct{}

type bitmap struct {
	words  []uint64
	length int
}

func newBitmap() *bitmap {
	return &bitmap{}
}

func (bitmap *bitmap) has(num int) bool {
	word, bit := num/64, uint(num%64)
	return word < len(bitmap.words) && (bitmap.words[word]&(1<<bit)) != 0
}

func (bitmap *bitmap) add(num int) {
	word, bit := num/64, uint(num%64)
	for word >= len(bitmap.words) {
		bitmap.words = append(bitmap.words, 0)
	}
	// 判断num是否已经存在bitmap中
	if bitmap.words[word]&(1<<bit) == 0 {
		bitmap.words[word] |= 1 << bit
		bitmap.length++
	}
}

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

func lengthOfLongestSubstring2(s string) int {
	length := len(s)
	if length == 0 {
		return 0
	}

	maxLen := 1
	startIndex := 0
	endIndex := 0
	tmpMaxLen := 0
	bitmapPtr := newBitmap()
	for {
		tmpVal := int(s[endIndex : endIndex+1][0])
		if bitmapPtr.has(tmpVal) {
			if tmpMaxLen > maxLen {
				maxLen = tmpMaxLen
			}
			startIndex++
			endIndex = startIndex
			bitmapPtr = newBitmap()
			tmpMaxLen = 0
			continue
		} else {
			bitmapPtr.add(tmpVal)
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

func lengthOfLongestSubstring3(s string) int {
	length := len(s)
	if length == 0 {
		return 0
	}

	maxLen := 0
	startIndex := 0
	endIndex := 0
	tmpLen := 0

	for ; endIndex < length; endIndex++ {
		for i := startIndex; i < endIndex; i++ {
			tmpLen = endIndex - startIndex
			if s[i] == s[endIndex] {
				if tmpLen > maxLen {
					maxLen = tmpLen
				}
				startIndex = i + 1
				if length-startIndex <= maxLen {
					return maxLen
				}
				break
			}
		}
	}
	return endIndex - startIndex
}

func lengthOfLongestSubstring4(s string) int {
	m := 0
	iS := 0
	iE := 0

	for ; iE < len(s); iE++ {
		for j := iS; j < iE; j++ {
			if s[iE] == s[j] {
				if iE-iS > m {
					m = iE - iS
				}

				iS = j + 1

				if len(s)-iS <= m {
					return m
				}
				break
			}
		}
	}

	return iE - iS
}
