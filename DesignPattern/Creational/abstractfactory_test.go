package creationaldesignpattern

import (
	"testing"
)

func TestStoreAFactory(t *testing.T) {

	var factory FruitFactory
	var res string
	factory = &StoreAFactory{}

	appleStoreA := factory.CreateAppleStore()
	res = appleStoreA.BuyApple()
	if res != "Buy Apple in AppleStoreA" {
		t.Fatal("Buy Apple in AppleStoreA Error")
	}
	orangeStoreA := factory.CreateOrangeStore()
	res = orangeStoreA.BuyOrange()
	if res != "Buy Orange in OrangeStoreA" {
		t.Fatal("Buy Orange in OrangeStoreA Error")
	}

}

func TestStoreBFactory(t *testing.T) {

	var factory FruitFactory
	var res string
	factory = &StoreAFactory{}

	factory = &StoreBFactory{}
	appleStoreB := factory.CreateAppleStore()
	res = appleStoreB.BuyApple()
	if res != "Buy Apple in AppleStoreB" {
		t.Fatal("Buy Apple in AppleStoreB Error")
	}
	orangeStoreB := factory.CreateOrangeStore()
	res = orangeStoreB.BuyOrange()
	if res != "Buy Orange in OrangeStoreB" {
		t.Fatal("Buy Orange in OrangeStoreB Error")
	}

}
