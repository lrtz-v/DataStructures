package sort

import "fmt"

//Template template
type Template struct {
}

// Sort ...
func (s Template) Sort(l []int64) {
	return
}

// Len return length
func (s Template) Len(l []int64) int {
	return len(l)
}

// Less check values
func (s Template) Less(l []int64, a, b int) bool {
	return l[a] < l[b]
}

// Exch exchanges in list
func (s Template) Exch(l []int64, i, j int) {
	l[i], l[j] = l[j], l[i]
}

// Show print the list
func (s Template) Show(l []int64) {
	fmt.Println(l)
}

// IsSorted check wether the list is already sorted
func (s Template) IsSorted(l []int64) bool {
	for i := 1; i < len(l); i++ {
		if l[i] < l[i-1] {
			return false
		}
	}
	return true
}
