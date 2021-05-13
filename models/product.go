package models

import (
	"database/sql"
	"errors"
	"fmt"
)

type Product struct {
	ProductId   int64
	ProductName string
	ProductCode string
	ReleaseDate string
	Price       float32
	Description string
	StarRating  float32
	ImageUrl    string
}

var (
	Products   []*Product
	nextProdID int64 = 1
	ProductDB  *sql.DB
)

func GetProducts() []*Product {
	return Products
}

func AddProduct(p Product) (Product, error) {
	if p.ProductId != 0 {
		return Product{}, errors.New("new Product must not include id or it must be set to zero")
	}
	p.ProductId = nextProdID
	nextProdID++
	//Products = append(Products, &p)
	// perform a db.Query insert
	insert, err := ProductDB.Query(fmt.Sprintf("CALL `products`.`productsAddOrEdit`(%d, %s, %s, %s, %s, %f, %f,%s)",
		p.ProductId, p.ProductName, p.ProductCode, p.ReleaseDate, p.Description, p.Price, p.StarRating, p.ImageUrl))

	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}
	// be careful deferring Queries if you are using transactions
	defer insert.Close()
	return p, nil
}

func GetProductByID(id int64) (Product, error) {
	for _, p := range Products {
		if p.ProductId == id {
			return *p, nil
		}
	}

	return Product{}, fmt.Errorf("Product with ID '%v' not found", id)
}

func UpdateProduct(p Product) (Product, error) {
	for i, candidate := range Products {
		if candidate.ProductId == p.ProductId {
			Products[i] = &p
			return p, nil
		}
	}

	return Product{}, fmt.Errorf("Product with ID '%v' not found", p.ProductId)
}

func RemoveProductById(id int64) error {
	for i, p := range Products {
		if int64(p.ProductId) == id {
			Products = append(Products[:i], Products[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Product with ID '%v' not found", id)
}
