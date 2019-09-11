package creationaldesignpattern

// Bubble drinks
type Bubble struct{}

// Drink drink Bubble
func (b *Bubble) Drink() string {
	return "cool~~"
}

// Pool beer pool
type Pool chan *Bubble

// NewObjectPool init pool
func NewObjectPool(total int) *Pool {
	p := make(Pool, total)

	for i := 0; i < total; i++ {
		p <- new(Bubble)
	}

	return &p
}
