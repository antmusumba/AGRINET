package services

import (
	"errors"

	"github.com/antmusumba/agrinet/internals/models"
	"github.com/antmusumba/agrinet/internals/repositories"
)

// ProductService represents the product service
type ProductService struct {
	repo repositories.ProductRepo
}

// NewProductService initializes the ProductService
func NewProductService(repo repositories.ProductRepo) *ProductService {
	return &ProductService{repo: repo}
}

// CreateProduct creates a new product
func (s *ProductService) CreateProduct(product *models.Product) error {
	if product.Title == "" {
		return errors.New("product title cannot be empty")
	}
	if product.Price <= 0 {
		return errors.New("product price must be greater than zero")
	}
	if product.Stock < 0 {
		return errors.New("product stock cannot be negative")
	}
	return s.repo.CreateProd(product)
}

// GetProductByID gets a product by id
func (s *ProductService) GetProductByID(id string) (*models.Product, error) {
	if id == "" {
		return nil, errors.New("product ID cannot be empty")
	}
	return s.repo.GetProdByID(id)
}

// UpdateProduct updates a product
func (s *ProductService) UpdateProduct(product *models.Product) error {
	if product.ID == "" {
		return errors.New("product ID cannot be empty")
	}
	return s.repo.UpdateProd(product)
}

// DeleteProduct deletes a product
func (s *ProductService) DeleteProduct(id string) error {
	if id == "" {
		return errors.New("product ID cannot be empty")
	}
	return s.repo.DeleteProd(id)
}

// ListProducts lists all products
func (s *ProductService) ListProducts() ([]*models.Product, error) {
	products, err := s.repo.ListProducts()
	if err != nil {
		return nil, err
	}
	if len(products) == 0 {
		return nil, errors.New("no products found")
	}
	return products, nil
}
