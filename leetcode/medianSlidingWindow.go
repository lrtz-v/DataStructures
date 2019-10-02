package leetcode

import (
	"sort"
)

func medianSlidingWindow(nums []int, k int) []float64 {
	knums := make([]int, k)
	copy(knums, nums[:k])
	sort.Ints(knums)
	ret := make([]float64, len(nums)-k+1)
	ret[0] = calMid(knums, k)
	for i := k; i < len(nums); i++ {
		delIdx := binarySearch(knums, nums[i-k])
		knums = append(knums[:delIdx], knums[delIdx+1:]...)
		addIdx := binarySearch(knums, nums[i])
		last := append([]int{nums[i]}, knums[addIdx:]...)
		knums = append(knums[:addIdx], last...)
		ret[i-k+1] = calMid(knums, k)
	}
	return ret
}

func binarySearch(nums []int, val int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + (high-low)/2
		if nums[mid] == val {
			return mid
		} else if nums[mid] > val {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return low
}

func calMid(nums []int, k int) float64 {
	n1, n2 := 0, 0
	if k%2 == 0 {
		n1, n2 = nums[k/2-1], nums[k/2]
	} else {
		n1, n2 = nums[k/2], nums[k/2]
	}
	return float64(n1+n2) / 2
}

func medianSlidingWindow2(nums []int, k int) []float64 {
	medianSlice := []float64{}

	getMedian := func(tmp []int) float64 {
		sort.Ints(tmp)
		l := len(tmp)
		if l%2 == 0 {
			return (float64)(tmp[l/2]+tmp[l/2-1]) / 2
		}
		return float64(tmp[l/2])
	}

	for index := 0; index <= len(nums)-k; index++ {
		tmp := make([]int, k)
		copy(tmp, nums[index:index+k])
		medianSlice = append(medianSlice, getMedian(tmp))
	}
	return medianSlice
}
