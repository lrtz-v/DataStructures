package creationaldesignpattern

import (
	"testing"
)

func TestDrinkFactory(t *testing.T) {

	var factory GoodsFactory
	factory = BuyGoodsFactory{}

	op := factory.Create()
	op.Set(1, "Coke")
	op.Buy()
	name := op.GetName()
	if name != "Coke" {
		t.Fatal("Wrong name")
	}
}
