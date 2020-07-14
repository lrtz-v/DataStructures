package findNumberIn2DArray

func findNumberIn2DArray(matrix [][]int, target int) bool {
	n := len(matrix)
	if n == 0 {
		return false
	}
	m := len(matrix[0])
	if m == 0 {
		return false
	}

	for i := 0; i < n; i++ {
		if matrix[i][m-1] < target {
			continue
		}

		index := find(matrix[i], target)
		if index == -1 {
			continue
		}
		if matrix[i][index] == target {
			return true
		}

	}

	return false
}

func find(nums []int, target int) int {
	start := 0
	end := len(nums)-1

	for start <= end {
		mid := start + (end - start) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			start = mid + 1
		}  else if nums[mid] > target {
			end = mid - 1
		}
	}
	return -1
}

func findNumberIn2DArray2(matrix [][]int, target int) bool {
	n := len(matrix)
	if n == 0 {
		return false
	}
	m := len(matrix[0])
	if m == 0 {
		return false
	}

	row := 0
	column := m - 1
	for row < n && column >= 0 {
		num := matrix[row][column];
		if num == target {
			return true
		} else if num > target {
			column--
		} else {
			row++
		}
	}

	return false
}
