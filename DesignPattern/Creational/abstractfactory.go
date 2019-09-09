package creationaldesignpattern

// AppleStore sale apple
type AppleStore interface {
	BuyApple() string
}

// OrangeStore sale orange
type OrangeStore interface {
	BuyOrange() string
}

// FruitFactory create fruit store
type FruitFactory interface {
	CreateAppleStore() AppleStore
	CreateOrangeStore() OrangeStore
}

// AppleStoreA apple store A
type AppleStoreA struct{}

// BuyApple buy apple in AppleStoreA
func (a *AppleStoreA) BuyApple() string {
	return "Buy Apple in AppleStoreA"
}

// OrangeStoreA orrange store A
type OrangeStoreA struct{}

// BuyOrange Buy Orange in OrangeStoreA
func (o *OrangeStoreA) BuyOrange() string {
	return "Buy Orange in OrangeStoreA"
}

// StoreAFactory Store A
type StoreAFactory struct{}

// CreateAppleStore create apple store
func (sa *StoreAFactory) CreateAppleStore() AppleStore {
	return &AppleStoreA{}
}

// CreateOrangeStore create orange store
func (sa *StoreAFactory) CreateOrangeStore() OrangeStore {
	return &OrangeStoreA{}
}

// AppleStoreB apple store B
type AppleStoreB struct{}

// BuyApple buy apple in AppleStoreB
func (a *AppleStoreB) BuyApple() string {
	return "Buy Apple in AppleStoreB"
}

// OrangeStoreB orrange store B
type OrangeStoreB struct{}

// BuyOrange Buy Orange in OrangeStoreB
func (o *OrangeStoreB) BuyOrange() string {
	return "Buy Orange in OrangeStoreB"
}

// StoreBFactory Store B
type StoreBFactory struct{}

// CreateAppleStore create apple store
func (sa *StoreBFactory) CreateAppleStore() AppleStore {
	return &AppleStoreB{}
}

// CreateOrangeStore create orange store
func (sa *StoreBFactory) CreateOrangeStore() OrangeStore {
	return &OrangeStoreB{}
}
