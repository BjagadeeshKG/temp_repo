package data

import (
	"encoding/json"
	"io"
	"time"
)

type Product struct {
	Name        string
	ID          int
	Description string
	Price       int
	CreatedOn   string
	Modifiedon  string
}

func (p *Product) FromJson(r io.Reader) error {
	jd := json.NewDecoder(r)
	err := jd.Decode(p)
	return err
}

type Products []*Product

var productslist Products

func (p *Products) ToJson(rw io.Writer) error {
	enc := json.NewEncoder(rw)
	err := enc.Encode(p)
	return err
}

func init() {
	productslist = make(Products, 0)
	productslist = append(productslist,
		&Product{
			Name:        "Latte",
			ID:          1,
			Description: "is nice",
			Price:       25,
			CreatedOn:   time.Now().UTC().String(),
			Modifiedon:  time.Now().UTC().String(),
		},
	)
	productslist = append(productslist,
		&Product{
			Name:        "espresso",
			ID:          2,
			Description: "is espresso nice",
			Price:       25,
			CreatedOn:   time.Now().UTC().String(),
			Modifiedon:  time.Now().UTC().String(),
		},
	)

}

func GetProducts() *Products {
	return &productslist
}
func AddProduct(prod *Product) {

	plast := productslist[len(productslist)-1]
	id := (plast.ID + 1)
	prod.ID = id
	productslist = append(productslist, prod)
}
