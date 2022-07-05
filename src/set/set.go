package set

type UnSortedSet[T any] interface {
	// return the length of set
	Size() int
	// add the element e to the set if not already present; Return true if x was added to the set and false otherwise
	Add(e T) bool
	// Find an element y in the set such that e equals y and remove y; Return y, or nil if no such element exists
	Remove(e T) T
	// Find an element y in the set such that y equals e. Return y, or nil if no such element exists
	Find(e T) T
}

type SortedSet[T comparable] interface {
	Size() int
	Add(e T) bool
	Remove(e T) T
	// Find the smallest element y in the set such that y >= e; Return y or nil if no such element exists
	Find(e T) T
}
