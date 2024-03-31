package products

import "github.com/google/uuid"

type Product struct {
	Id          uuid.UUID
	ProductName string
	Price       int
}

func New(productName string, price int) Product {
	return Product{
		Id:          uuid.New(),
		ProductName: productName,
		Price:       price,
	}
}
