package creationaldesignpattern

// Drinks interface
type Drinks interface {
	Clone() Drinks
}

// DrinksManager a lot of drinks
type DrinksManager struct {
	drinkManager map[string]Drinks
}

// Set add drinks to DrinksManager
func (d *DrinksManager) Set(name string, drinks Drinks) {
	d.drinkManager[name] = drinks
}

// Get drinks from DrinksManager
func (d *DrinksManager) Get(name string) Drinks {
	return d.drinkManager[name]
}

// NewDrinksManager init DrinksManager
func NewDrinksManager() *DrinksManager {
	return &DrinksManager{
		drinkManager: make(map[string]Drinks),
	}
}
