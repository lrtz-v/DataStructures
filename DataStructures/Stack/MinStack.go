package Stack

// MinStack . 最小栈
type MinStack struct {
	mainStack []int
	minStack  []int
}

// Push put element into stack
func (m *MinStack) Push(data int) {
	m.mainStack = append(m.mainStack, data)
	if len(m.minStack) == 0 {
		m.minStack = append(m.minStack, data)
	} else {
		min := m.minStack[len(m.minStack)-1]
		if data <= min {
			m.minStack = append(m.minStack, data)
		}
	}
}

// Pop pop element
func (m *MinStack) Pop() int {
	popData := m.mainStack[len(m.mainStack)-1]
	m.mainStack = m.mainStack[0 : len(m.mainStack)-1]
	min := m.minStack[len(m.minStack)-1]
	if popData == min {
		m.minStack = m.minStack[0 : len(m.minStack)-1]
	}
	return popData
}

// GetMin return min element, no Pop
func (m *MinStack) GetMin() int {
	return m.minStack[len(m.minStack)-1]
}

// NewMinStack init MinStack
func NewMinStack() *MinStack {
	return &MinStack{}
}
