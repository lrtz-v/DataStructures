package intrusivelinklist

/*
	介入式链表✅
*/

// Intrusive Intrusive
type Intrusive interface {
	Next() Intrusive
	Prev() Intrusive
	AddNext(Intrusive)
	AddPrev(Intrusive)
}

// List provides the implementation of intrusive linked lists
type List struct {
	prev Intrusive
	next Intrusive
}

// Next Next
func (l *List) Next() Intrusive {
	return l.next
}

// Prev Prev
func (l *List) Prev() Intrusive {
	return l.prev
}

// AddNext AddNext
func (l *List) AddNext(i Intrusive) {
	l.next = i
}

// AddPrev AddPrev
func (l *List) AddPrev(i Intrusive) {
	l.prev = i
}

// Reset Reset
func (l *List) Reset() {
	l.prev = nil
	l.next = nil
}

// Empty Empty
func (l *List) Empty() bool {
	return l.prev == nil
}

// Front returns the first element of list l or nil.
func (l *List) Front() Intrusive {
	return l.prev
}

// Back returns the last element of list l or nil.
func (l *List) Back() Intrusive {
	return l.next
}

// PushFront inserts the element e at the front of list l.
func (l *List) PushFront(e Intrusive) {
	e.AddPrev(nil)
	e.AddNext(l.prev)

	if l.prev != nil {
		l.prev.AddPrev(e)
	} else {
		l.next = e
	}
	l.prev = e
}

// PushBack inserts the element e at the back of list l.
func (l *List) PushBack(e Intrusive) {
	e.AddNext(nil)
	e.AddPrev(l.next)

	if l.next != nil {
		l.next.AddNext(e)
	} else {
		l.prev = e
	}
	l.next = e
}

// InsertAfter inserts e after b.
func (l *List) InsertAfter(e, b Intrusive) {
	a := b.Next()
	e.AddNext(a)
	e.AddPrev(b)
	b.AddNext(e)

	if a != nil {
		a.AddPrev(e)
	} else {
		l.next = e
	}
}

// InsertBefore inserts e before a.
func (l *List) InsertBefore(e, a Intrusive) {
	b := a.Prev()
	e.AddNext(a)
	e.AddPrev(b)
	a.AddPrev(e)

	if b != nil {
		b.AddNext(e)
	} else {
		l.prev = e
	}
}

// Remove removes e from l.
func (l *List) Remove(e Intrusive) {
	prev := e.Prev()
	next := e.Next()

	if prev != nil {
		prev.AddNext(next)
	} else {
		l.prev = next
	}

	if next != nil {
		next.AddPrev(prev)
	} else {
		l.next = prev
	}
}
