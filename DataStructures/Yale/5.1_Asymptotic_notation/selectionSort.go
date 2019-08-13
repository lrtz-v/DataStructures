package asymptoticnotation

// SelectSort select sort
func SelectSort(array []int) []int {
	n := len(array)
	for i := 0; i < n-1; i++ {

		min := i + 1
		for j := i + 1; j < n; j++ {
			if array[j] < array[min] {
				min = j
			}
		}
		if array[i] > array[min] {
			array[i], array[min] = array[min], array[i]
		}
	}

	return array
}
