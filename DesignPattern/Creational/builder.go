package creationaldesignpattern

// Beer interface
type Beer interface {
	Drink()
}

// Drink struct
type Drink struct {
	beer Beer
}

// InitDrink init Drink
func InitDrink(beer Beer) *Drink {
	return &Drink{
		beer: beer,
	}
}

// Drink beer
func (d *Drink) Drink() {
	d.beer.Drink()
}

// MUSU beer
type MUSU struct {
	result string
}

// Drink method
func (m *MUSU) Drink() {
	m.result = "Drink MUSU"
}

// GetResult return drink
func (m *MUSU) GetResult() string {
	return m.result
}

// QingDao beer
type QingDao struct {
	result string
}

// Drink beer
func (q *QingDao) Drink() {
	q.result = "Drink QingDao"
}

// GetResult return drink
func (q *QingDao) GetResult() string {
	return q.result
}
