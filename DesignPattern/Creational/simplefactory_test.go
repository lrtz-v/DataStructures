package creationaldesignpattern

import (
	"testing"
)

func TestNewDrink(t *testing.T) {
	drink := NewDrink(CokeType)
	s := drink.Drink()
	if s != "Coke~~" {
		t.Fatal("TestNewDrink test fail")
	}
}

func TestDefaultDrink(t *testing.T) {
	drink := NewDrink(0)
	s := drink.Drink()
	if s != "Coke~~" {
		t.Fatal("TestDefaultDrink test fail")
	}
}
