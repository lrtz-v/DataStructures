package letterCasePermutation

import "strconv"

var numberMapping = map[byte][]byte{
	'2': {'a', 'b', 'c'},
	'3': {'d', 'e', 'f'},
	'4': {'g', 'h', 'i'},
	'5': {'j', 'k', 'l'},
	'6': {'m', 'n', 'o'},
	'7': {'p', 'q', 'r', 's'},
	'8': {'t', 'u', 'v'},
	'9': {'w', 'x', 'y', 'z'},
}

func letterCombinations(digits string) []string {
	length := len(digits)
	ans := []string{}

	if length == 0 {
		return ans
	}

	if !isDigits(digits) {
		return ans
	}

	var backtrace func(digits string, str string)

	backtrace = func(digits string, str string) {
		if len(digits) == 0 {
			ans = append(ans, str)
			return
		}

		bytes, ok := numberMapping[digits[0]]
		if !ok {
			backtrace(digits[1:], str)
		}

		for _, v := range bytes {
			backtrace(digits[1:], str + string(v))
		}
	}

	backtrace(digits, "")

	return ans
}

func isDigits(digits string) bool {
	_, err := strconv.Atoi(digits)
	return err == nil
}
