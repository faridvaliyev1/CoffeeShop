package data

import "time"

type Product struct {
	ID          int     `json:"Id"`
	Name        string  `json:"Name"`
	Description string  `json:"Description"`
	Price       float32 `json:"Price"`
	SKU         string  `json:"Sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

var productList = []*Product{
	&Product{
		ID:          1,
		Name:        "Latte",
		Description: "It is delicious",
		Price:       2.45,
		SKU:         "123abd",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Name:        "Espresso",
		Description: "Short description",
		Price:       1.99,
		SKU:         "fgg123",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}

func GetProducts() []*Product {
	return productList
}
