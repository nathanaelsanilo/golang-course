package main

import "fmt"

type Product struct {
	ID    int
	Name  string
	Price float64
	Stock float64
}

func BuildProduct() Product {
	return Product{
		ID:    1,
		Name:  "",
		Price: 0,
		Stock: 0,
	}
}

func (p Product) IsAvailable() bool {
	return p.Stock > 0
}

func (p *Product) ApplyDiscount(percent float64) {
	p.Price -= (p.Price * percent)
}

func (p Product) String() string {
	return fmt.Sprintf("[%d] %s - Rp%.0f", p.ID, p.Name, p.Price)
}
