package leetcode

import "fmt"

func findSubsequences(nums []int) [][]int {

	var subs [][]int

	for _, num := range nums {
		var tmp [][]int
		tmp = append(tmp, []int{num})

		for i := 0; i < len(subs); i++ {
			if subs[i][len(subs[i])-1] <= num {
				tmp = append(tmp, subs[i])
				subs[i] = append(subs[i], num)
			}
		}
		subs = append(subs, tmp...)
	}
	fmt.Println(subs)

	var result [][]int

	for i := 0; i < len(subs); i++ {
		if len(subs[i]) >= 2 {
			result = append(result, subs[i])
		}
	}

	return result

}
