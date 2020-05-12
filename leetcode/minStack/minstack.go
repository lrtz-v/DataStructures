package minstack

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

// MinStack stack
type MinStack struct {
	m []int
	q []int
}

// Constructor initialize
func Constructor() MinStack {
	return MinStack{}
}

// Push put element into stack
func (this *MinStack) Push(x int) {
	this.q = append(this.q, x)

	if len(this.m) == 0 {
		this.m = append(this.m, x)
	} else {
		min := this.m[len(this.m)-1]
		if x < min {
			this.m = append(this.m, x)
		} else {
			this.m = append(this.m, min)
		}
	}
}

// Pop delete top element
func (this *MinStack) Pop() {
	this.q = this.q[0 : len(this.q)-1]
	this.m = this.m[0 : len(this.m)-1]
}

// Top return top element
func (this *MinStack) Top() int {
	return this.q[len(this.q)-1]
}

// GetMin return min element
func (this *MinStack) GetMin() int {
	return this.m[len(this.m)-1]
}
