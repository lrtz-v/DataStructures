package sort

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

// Bucket 桶排序
type Bucket struct {
	Template
}

// Sort 桶排序
func (b Bucket) Sort(arr []int64) []int64 {
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
	bucketList := make([][]int64, bucketNum)
	for i := 0; i < bucketNum; i++ {
		bucketList[i] = []int64{}
	}

	// put elements to bucket
	for i := 0; i < len(arr); i++ {
		num := int((arr[i] - min) * int64(bucketNum-1) / d)
		// index := num.(float64)
		bucketList[num] = append(bucketList[num], arr[i])
	}

	// bucket sort
	for i := 0; i < len(bucketList); i++ {
		bucketList[i] = Count{}.Sort(bucketList[i])
	}

	// sorted arr
	sortedArr := make([]int64, len(arr))
	index := 0
	for _, bucket := range bucketList {
		for _, val := range bucket {
			sortedArr[index] = val
			index++
		}
	}

	return sortedArr
}
