package models

import (
	"errors"
	"fmt"
)

type Product struct {
	productId   int64
	productName string
	productCode string
	releaseDate string
	price       float32
	description string
	starRating  float32
	imageUrl    string
}

var (
	Products   []*Product
	nextProdID int64 = 1
)

func GetProducts() []*Product {
	return Products
}

func AddProduct(p Product) (Product, error) {
	if p.productId != 0 {
		return Product{}, errors.New("new Product must not include id or it must be set to zero")
	}
	p.productId = nextProdID
	nextProdID++
	Products = append(Products, &p)
	return p, nil
}

func GetProductByID(id int64) (Product, error) {
	for _, p := range Products {
		if p.productId == id {
			return *p, nil
		}
	}

	return Product{}, fmt.Errorf("Product with ID '%v' not found", id)
}

func UpdateProduct(p Product) (Product, error) {
	for i, candidate := range Products {
		if candidate.productId == p.productId {
			Products[i] = &p
			return p, nil
		}
	}

	return Product{}, fmt.Errorf("Product with ID '%v' not found", p.productId)
}

func RemoveProductById(id int64) error {
	for i, p := range Products {
		if int(p.productId) == id {
			Products = append(Products[:i], Products[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("Product with ID '%v' not found", id)
}
