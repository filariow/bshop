package bshop

type Bill struct {
	Products []Product
	Discount Discount
}

type Discount interface {
	//CalculateOn calculates the discount on the total
	CalculateOn(total float64) float64
}

func (b *Bill) Total() float64 {
	t := float64(0.)
	for _, p := range b.Products {
		t += p.Price
	}

	if b.Discount != nil {
		t -= b.Discount.CalculateOn(t)
	}
	return t
}
