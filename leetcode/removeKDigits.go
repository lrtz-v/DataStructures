package leetcode

import (
	"fmt"
	"strconv"
	"strings"
)

/*
	从一个整数中删除K个数字，使得新数字尽可能小

	从高位开始遍历，如果一个数字大于右侧的数字，则删除
*/

func removeKDigits(num string, k int) string {
	if len(num) == 0 || (len(num) == 1 && k > 0) {
		return ""
	} else if k == 0 {
		return num
	}

	arr := strings.Split(num, "")
	res := make([]string, len(arr))
	index := 0
	for i := 0; i < len(arr); i++ {
		res[index] = arr[i]
		if index == 0 || k == 0 {
			index++
			continue
		}

		for j := i; j > 0 && k > 0; j-- {
			val, _ := strconv.Atoi(res[j])
			lastVal, _ := strconv.Atoi(res[j-1])
			if lastVal > val {
				res[j-1], res[j] = res[j], ""
				index--
				k--
			}
		}
		//
		//fmt.Printf("val: %d, lastVal: %d\n", val, lastVal)
		//if lastVal > val {
		//	res[i-1] = ""
		//	res[i] = arr[i]
		//	k--
		//} else {
		//	res[i] = arr[i]
		//}
		fmt.Println(res)
	}

	return strings.Join(res, "")
}

func removeKDigits2(num string, k int) string {
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
