package letterCasePermutation

import (
	"strconv"
)

func letterCasePermutation(S string) []string {
	ans := []string{}

	if len(S) == 0 {
		return ans
	}

	var backtrace func(str string, i int)
	backtrace = func(str string, i int) {
		if i == len(str) {
			ans = append(ans, str)
			return
		}

		// change nothing
		backtrace(str, i+1)
		
		// chang letter
		if (str[i] >= 'A' && str[i] <= 'Z') || (str[i] >= 'a' && str[i] <= 'z') {
			str = letterChange(str, i)
			backtrace(str, i+1)
		}
	}
	backtrace(S, 0)

	return ans
}

func isDigits(digits string) bool {
	_, err := strconv.Atoi(digits)
	return err == nil
}

func letterChange(str string, i int) string {
	tmp := []byte(str)
	tmp[i] ^= 32
	return string(tmp)
}
