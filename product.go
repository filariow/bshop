package bshop

//Product represents a product
type Product struct {
	ID    uint64
	Name  string
	Price float64
	Cost  float64
}

//Beer representes a beer
type Beer struct {
	Product

	Size   float64
	Vol    float64
	Brewer string
}
