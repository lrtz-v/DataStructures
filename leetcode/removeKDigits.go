package leetcode

import (
	"strings"
)

/*
	从一个整数中删除K个数字，使得新数字尽可能小

	从高位开始便历，如果一个数字大于右侧的数字，则删除
*/

func removeKDigits(num string, k int) string {
	if len(num) == 0 || (len(num) == 1 && k > 0) {
		return ""
	} else if k == 0 {
		return num
	}

	arr := strings.Split(num, "")
	res := make([]string, len(arr))

	for i := 0; i < len(arr); i++ {
		if k > 0 {
			for j := 0; j < len(res); j++ {

			}
		}
		res[i] = arr[i]
	}

	return strings.Join(res, "")
}
