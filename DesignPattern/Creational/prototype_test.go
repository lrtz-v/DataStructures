package creationaldesignpattern

import (
	"testing"
)

var drinkManager *DrinksManager

// Water drinks
type Water struct {
	name string
}

func (w *Water) Clone() Drinks {
	anotherOne := *w
	return &anotherOne
}

// Budweiser drinks
type Budweiser struct {
	name string
}

func (b *Budweiser) Clone() Drinks {
	anotherOne := *b
	return &anotherOne
}

func TestClone(t *testing.T) {
	t1 := drinkManager.Get("Water")

	t2 := t1.Clone()

	if t1 == t2 {
		t.Fatal("error! get clone not working")
	}
}

func TestCloneFromManager(t *testing.T) {
	c := drinkManager.Get("Water").Clone()

	t1 := c.(*Water)
	if t1.name != "Water" {
		t.Fatal("error")
	}

}

func init() {
	drinkManager = NewDrinksManager()
	drinkManager.Set(
		"Water", &Water{name: "Water"},
	)

}
