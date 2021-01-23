package bshop

//Product represents a product
type Product struct {
	ID    int64
	Name  string
	Price float64
	Cost  float64
}

//Beer representes a beer
type Beer struct {
	Product

	SizeMl float64
	Vol    float64
	Brewer string
}
