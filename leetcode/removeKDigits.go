package leetcode

/*
	从一个整数中删除K个数字，使得新数字尽可能小

	从高位开始便历，如果一个数字大于右侧的数字，则删除
*/

func removeKDigits(arr []int, k int) []int {
	tmp := arr
	for i := 0; i < k; i++ {
		if len(tmp) <= 1 {
			return []int{}
		}
		for j := 0; j <= len(tmp)-2; j++ {
			if tmp[j] > tmp[j+1] {
				var res []int
				res = append(res, tmp[0:j]...)
				res = append(res, tmp[j+1:]...)
				tmp = res
				break
			}
		}
	}
	return tmp
}
