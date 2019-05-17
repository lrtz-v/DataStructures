package leetcode

// There are two sorted arrays nums1 and nums2 of size m and n respectively.
// Find the median of the two sorted arrays. The overall run time complexity should be O(log (m+n)).
// You may assume nums1 and nums2 cannot be both empty.
// Example 1:
// nums1 = [1, 3]
// nums2 = [2]
// The median is 2.0
// Example 2:
// nums1 = [1, 2]
// nums2 = [3, 4]
// The median is (2 + 3)/2 = 2.5

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var tmpArr []int
	tmpArr = append(tmpArr, nums1...)
	tmpArr = append(tmpArr, nums2...)

	// max, min
	max := 0
	min := 0
	for _, v := range tmpArr {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}

	indexArr := make([]int, max-min+1)
	var resArr []int
	for _, data := range tmpArr {
		indexArr[data-min]++
	}

	for index, count := range indexArr {
		for j := 0; j < count; j++ {
			resArr = append(resArr, index+min)
		}
	}

	if len(resArr)%2 == 0 {
		m := len(resArr)/2 - 1
		return float64(resArr[m]+resArr[m+1]) / 2
	}
	m := (len(resArr)+1)/2 - 1
	return float64(resArr[m])
}
