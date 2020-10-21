package calculate

import (
	"strconv"
	"strings"
)

func calculate(s string) int {
	if len(s) == 0 {
		return 0
	}

	str := strings.TrimSpace(s)
	if len(str) == 0 {
		return 0
	}

	nums := []int{}
	ops := []byte{}

	for i := 0; i < len(str); {
		if str[i] == ' ' {
			i++
			continue
		} else if str[i] >= '0' && str[i] <= '9' {
			num, index := getNextNum(str, i)
			nums = append(nums, num)
			i = index
		} else if isOperator(str[i]) {
			op := str[i]
			i++
			if op == '+' || op == '-' {
				ops = append(ops, op)
			} else if op == '*' || op == '/' {
				nextNum, index := getNextNum(str, i)
				i = index

				if op == '*' {
					nums[len(nums)-1] *= nextNum
				} else if op == '/' {
					nums[len(nums)-1] /= nextNum
				}
			}
		}
	}

	ans := nums[0]
	nums = nums[1:]

	for {
		if len(ops) <= 0 {
			break
		}
		op := ops[0]
		ops = ops[1:]

		if op == '+' {
			ans += nums[0]
		} else {
			ans -= nums[0]
		}
		nums = nums[1:]
	}

	return ans
}

func getNextNum(s string, i int) (int, int) {
	for ;i < len(s); i++ {
		if s[i] != ' ' {
			break
		} else {
			continue
		}
	}

	j := i
	for ; j < len(s); j++ {
		if '0' <= s[j] && s[j] <= '9' {
			continue
		} else {
			break
		}
	}
	num, _ := strconv.Atoi(s[i:j])
	return num, j
}

func isOperator(b byte) bool {
	return b == '+' || b == '-' || b == '*' || b == '/'
}