package creationaldesignpattern

// DrinkType type of drink
type DrinkType int

const (
	// WineType wine
	WineType DrinkType = 1 << iota
	// CokeType coke
	CokeType
	// BeerType beer
	BeerType
)

// Drink interface
type Drink interface {
	Drink() string
}

//Wine is one of API implement
type Wine struct{}

// Drink do drinking
func (*Wine) Drink() string {
	return "Wine~~"
}

//Coke is one of API implement
type Coke struct{}

// Drink do drinking
func (*Coke) Drink() string {
	return "Coke~~"
}

//Beer is one of API implement
type Beer struct{}

// Drink do drinking
func (*Beer) Drink() string {
	return "Beer~~"
}

// NewDrink new drink
func NewDrink(drinkType DrinkType) Drink {
	switch drinkType {
	case WineType:
		return &Wine{}
	case CokeType:
		return &Coke{}
	case BeerType:
		return &Beer{}
	default:
		return &Coke{}
	}
}
