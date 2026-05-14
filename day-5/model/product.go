package model

type Product struct {
	Id       int
	Name     string
	Category string
}

func BuildProduct(id int, name, category string) Product {
	return Product{Id: id, Name: name, Category: category}
}
