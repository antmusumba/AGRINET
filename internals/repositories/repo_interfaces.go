package repositories

import "github.com/antmusumba/agrinet/internals/models"

// UserRepo represents the user repository
type UserRepo interface {
	CreateUser(user *models.User) error
	GetUserByEmail(email string) (*models.User, error)
	GetUserByID(id string) (*models.User, error)
	UpdateUser(user *models.User) error
	DeleteUser(id string) error
}

// ProductRepo represents the product repository
type ProductRepo interface {
	CreateProd(product *models.Product) error
	GetPrdByEmail(email string) (*models.Product, error)
	GetProdByID(id string) (*models.Product, error)
	UpdateProd(product *models.Product) error
	DeleteProd(id string) error
	ListProducts() ([]*models.Product, error)
}
