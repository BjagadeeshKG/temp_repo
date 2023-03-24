package Handlers

import (
	"LearnGoService/data"
	"fmt"
	"net/http"
)

type Products struct {
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		p.GetProducts(rw, r)
		return
	}
	if r.Method == http.MethodPost {
		p.AddProducts(rw, r)
		return
	}
	rw.WriteHeader(http.StatusMethodNotAllowed)
}
func (p *Products) AddProducts(rw http.ResponseWriter, r *http.Request) {

	prod := &data.Product{}
	err := prod.FromJson(r.Body)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	data.AddProduct(prod)

}

func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	pl := data.GetProducts()
	err := pl.ToJson(rw)
	if err != nil {
		fmt.Println("Error occured while encode")
	}

}
