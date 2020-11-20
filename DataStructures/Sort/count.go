package sort

// Count 计数排序
type Count struct {
	Template
}

// Sort 计数排序
func (c Count) Sort(l []int64) []int64 {
	size := len(l)
	if size == 0 {
		return l
	}
	max := l[0]
	for i := 1; i < size; i++ {
		if max < l[i] {
			max = l[i]
		}
	}

	arr := make([]int64, max+1)
	for i := 0; i < size; i++ {
		arr[int(l[i])]++
	}

	ans := []int64{}
	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 {
			continue
		}
		for j := int64(0); j < arr[i]; j++ {
			ans = append(ans, int64(i))
		}
	}
	return ans
}
