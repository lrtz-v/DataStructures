package creationaldesignpattern

// DrinkType type of drink
type DrinkType int

const (
	// WineType wine
	WineType DrinkType = 1 << iota
	// CokeType coke
	CokeType
	// SodaType beer
	SodaType
)

// Drinking interface
type Drinking interface {
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

//Soda is one of API implement
type Soda struct{}

// Drink do drinking
func (*Soda) Drink() string {
	return "Soda~~"
}

// NewDrink new drink
func NewDrink(drinkType DrinkType) Drinking {
	switch drinkType {
	case WineType:
		return &Wine{}
	case CokeType:
		return &Coke{}
	case SodaType:
		return &Soda{}
	default:
		return &Coke{}
	}
}
