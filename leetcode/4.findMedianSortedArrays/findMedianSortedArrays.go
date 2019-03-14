package findMedianSortedArrays

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
	shortList, longList := nums1, nums2
	if len(nums1) > len(nums2) {
		shortList, longList = longList, shortList
	}
	return 1 / 2.0
}