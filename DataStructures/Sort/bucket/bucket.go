package bucket

import (
	"sort"
)

/*
桶排序
桶代表一个区间范围，可以承载一个或多个元素
1.确认桶的数量：元素数量，最后一个桶只包含最大值
2.确认区间跨度： (最大值 - 最小值) / (桶的数量 - 1)
3.把元素放入桶中
4.对每个桶内元素排序

假设有n个元素，n个桶
时间复杂度度为O(n) - o(nlogn)

缺点：
需要元素分布较为均匀，否则元素会集中分在某几个桶中
*/

func bucketSort(arr []float64) []float64 {
	if len(arr) == 0 {
		return arr
	}
	// get max && min
	max := arr[0]
	min := arr[0]
	for _, v := range arr {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	d := max - min

	// init buckets
	bucketNum := len(arr)
	bucketList := make([][]float64, bucketNum)
	for i := 0; i < bucketNum; i++ {
		bucketList[i] = []float64{}
	}

	// put elements to bucket
	for i := 0; i < len(arr); i++ {
		num := int((arr[i] - min) * float64(bucketNum-1) / d)
		// index := num.(float64)
		bucketList[num] = append(bucketList[num], arr[i])
	}

	// bucket sort
	for i := 0; i < len(bucketList); i++ {
		a := sort.Float64Slice(bucketList[i])
		sort.Sort(a)
	}

	// sorted arr
	sortedArr := make([]float64, len(arr))
	index := 0
	for _, bucket := range bucketList {
		for _, val := range bucket {
			sortedArr[index] = val
			index++
		}
	}

	return sortedArr
}
