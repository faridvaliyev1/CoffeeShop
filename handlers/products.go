package handlers

import (
	"log"
	"net/http"

	"github.com/faridvaliyev1/v2/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.GetProducts(rw, r)
		return
	}

	if r.Method == http.MethodPost {
		p.AddProduct(rw, r)
		return
	}

	//handle an update
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()
	err := lp.ToJSON(rw)

	if err != nil {
		http.Error(rw, "Unable to marshall json", http.StatusInternalServerError)
	}
}

func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
	prod := &data.Product{}

	err := prod.FromJSON(r.Body)

	if err != nil {
		http.Error(rw, "Unable to convert json", http.StatusBadRequest)
	}

	data.AddProduct(prod)

}
