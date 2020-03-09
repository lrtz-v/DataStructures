package skiplist

import (
	"fmt"
	"testing"
)

func TestInsert(t *testing.T) {
	zsl := zslInit()
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for _, v := range data {
		zsl.insert(1.0, v)
	}
	x := zsl.header
	for {
		if x.level[0].forward != nil {
			fmt.Println(x.level[0].forward.obj)
			x = x.level[0].forward
		} else {
			break
		}

	}
}
