package data

import (
	"encoding/json"
	"fmt"
	"io"
	"time"

	"github.com/go-playground/validator/v10"
)

type Product struct {
	ID          int     `json:"Id"`
	Name        string  `json:"Name" validate:"required"`
	Description string  `json:"Description"`
	Price       float32 `json:"Price"`
	SKU         string  `json:"Sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

type Products []*Product

func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *Product) Validate() error {
	validate := validator.New()
	return validate.Struct(p)
}

func GetNextID() int {
	lp := productList[len(productList)-1]
	return lp.ID + 1
}

func (p *Product) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

func AddProduct(p *Product) {
	p.ID = GetNextID()
	productList = append(productList, p)
}

func UpdateProduct(id int, p *Product) error {
	_, pos, err := FindProductById(id)

	if err != nil {
		return err
	}

	p.ID = id
	productList[pos] = p

	return nil
}

func FindProductById(id int) (*Product, int, error) {
	for i, p := range productList {
		if p.ID == id {
			return p, i, nil
		}
	}

	return nil, -1, ErrorProductNotFound
}

var ErrorProductNotFound = fmt.Errorf("Product NOt Found")

func GetProducts() Products {
	return productList
}

var productList = Products{
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
