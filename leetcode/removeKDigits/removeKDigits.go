package removeKDigits

import (
	"strconv"
	"strings"
)

/*
	从一个整数中删除K个数字，使得新数字尽可能小

	从高位开始遍历，如果一个数字大于右侧的数字，则删除

*/

func removeKDigits(num string, k int) string {
	if len(num) == 0 || k >= len(num) {
		return ""
	} else if k == 0 {
		return num
	}

	arr := strings.Split(num, "")
	res := make([]string, len(arr))
	for i := 0; i < len(arr); i++ {
		res[i] = arr[i]
		if k == 0 {
			continue
		}

		for j := i; j > 0 && k > 0; j-- {
			if res[j-1] == "" {
				continue
			}
			val, _ := strconv.Atoi(res[i])
			lastVal, _ := strconv.Atoi(res[j-1])
			if lastVal > val {
				res[j-1] = ""
				k--
			}
		}
	}
	if k > 0 {
		res = res[0 : len(res)-k]
	}

	return strings.Join(res, "")
}
