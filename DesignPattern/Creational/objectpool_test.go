package creationaldesignpattern

import (
	"fmt"
	"testing"
)

func TestDrinkBubble(t *testing.T) {
	// var pool chan *Bubble
	pool := NewObjectPool(2)
	for i := 0; i < 2; i++ {
		select {
		case obj := <-*pool:
			s := obj.Drink()
			if s != "cool~~" {
				t.Fatal("No cool~~")
			}
			fmt.Println(s)
		default:
			return
		}
	}
}
