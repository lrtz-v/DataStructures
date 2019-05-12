package leetcode

import (
	"sort"
)

/*
找出无需数组排序后的最大相邻差
1. 使用计数排序，最大相邻差为0连续出现大的最大长度
2. 使用桶排序，遍历所有桶，统计每个桶的最大值和桶右侧非空桶的最小值的差
*/

func getMaxSortedDistance(arr []float64) float64 {
	if len(arr) == 0 {
		return 0
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
	// get max distance
	var maxDistance float64
	var maxLeft float64
	for i := 0; i < len(bucketList); i++ {
		bucket := bucketList[i]
		if len(bucket) == 0 {
			continue
		}
		if bucket[0]-maxLeft > maxDistance {
			maxDistance = bucket[0] - maxLeft
		}
		maxLeft = bucket[len(bucket)-1]
	}

	return maxDistance
}
