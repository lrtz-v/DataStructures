package list

type List[T any] interface {
	// return the length of the list
	Size() int
	// return the value Ei
	Get(i int) T
	// set the value of Ei equal to e
	Set(i int, e T)
	// set Ei = e, displacing elemtents if index >= i
	AddAtIndex(i int, e T)
	// remove the value Ei, displacing elemtents if index >= i
	Remove(e T)
}
