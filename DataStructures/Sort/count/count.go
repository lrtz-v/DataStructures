package count

import "fmt"

/*
计数排序
假设原始数列的规模是n，最大和最小数差值是m
时间复杂度：O(n+m)

缺点：
1.最大数与最小数差值过大时，不适合
1.元素不是整数时，不适合
*/

func countSort(arr []int) {
	// get max val
	max := 0
	for _, v := range arr {
		if v > max {
			max = v
		}
	}
	indexArr := make([]int, max+1)
	for _, data := range arr {
		indexArr[data]++
	}
	for index, count := range indexArr {
		for j := 0; j < count; j++ {
			fmt.Println(index)
		}
	}
}

// 优化版
func countSort2(arr []int) {
	// get max val
	max := 0
	min := 0
	for _, v := range arr {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	indexArr := make([]int, max-min+1)
	for _, data := range arr {
		indexArr[data-min]++
	}
	for index, count := range indexArr {
		for j := 0; j < count; j++ {
			fmt.Println(index + min)
		}
	}
}
