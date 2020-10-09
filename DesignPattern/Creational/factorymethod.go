package creationaldesignpattern

import "fmt"

// 工厂方法模式

// GoodsType type of Goods

// Goods interface
type Goods interface {
	Set(price float32, name string)
	Buy() string
	GetName() string
}

// GoodsFactory Factory interface
type GoodsFactory interface {
	Create() Goods
}

// GoodsBase Goods
type GoodsBase struct {
	price float32
	name  string
}

// Set set name and price
func (d *GoodsBase) Set(price float32, name string) {
	d.price = price
	d.name = name
}

// Buy buy Goods
func (d *GoodsBase) Buy() string {
	return fmt.Sprintf("Get %s With %f$\n", d.name, d.price)
}

// BuyGoods create Goods method
type BuyGoods struct {
	*GoodsBase
}

// GetName create method GetName
func (b *BuyGoods) GetName() string {
	return b.name
}

// BuyGoodsFactory Factory of BuyGoods
type BuyGoodsFactory struct{}

// Create init GoodsBase
func (BuyGoodsFactory) Create() Goods {
	return &BuyGoods{
		GoodsBase: &GoodsBase{},
	}
}
