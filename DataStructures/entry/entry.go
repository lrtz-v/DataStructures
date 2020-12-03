package entry

import (
	"github.com/dchest/siphash"
)

type (
	// Key key
	Key string

	// Value value
	Value int
)

// Entry 节点
type Entry struct {
	Key   Key
	Value Value
}

// Hash return hash code
func (k Key) Hash(seed uint64) uintptr {
	return uintptr(siphash.Hash(seed, seed, []byte(k)))
}
