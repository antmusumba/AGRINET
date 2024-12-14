package repositories

import "github.com/antmusumba/agrinet/internals/models"

type UserRepo interface {
	CreateUser(user *models.User) error
	GetUserByEmail(username string) (*models.User, error)
	GetUserByID(id string) (*models.User, error)
	UpdateUser(user *models.User) error
	DeleteUser(id string) error
}

type ProdRepo interface {
	CreateProd(user *models.Product) error
	GetPrdByEmail(username string) (*models.Product, error)
	GetProdByID(id string) (*models.Product, error)
	UpdateProd(user *models.Product) error
	DeleteProd(id string) error
}
