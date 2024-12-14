package repositories

import (
	"database/sql"
	"errors"
	"log"
	"time"

	"github.com/antmusumba/agrinet/internals/models"
	"github.com/google/uuid"
)

// productRepo represents the product repository
type productRepo struct {
	db *sql.DB
}

// NewProductRepo initializes a new instance of productRepo
func NewProductRepo(db *sql.DB) ProductRepo {
	err := db.Ping()
	if err != nil {
		log.Printf("Database connection error: %v", err)
	} else {
		log.Println("Database connected successfully")
	}
	return &productRepo{db: db}
}

// CreateProd creates a new product
func (r *productRepo) CreateProd(product *models.Product) error {
	product.ID = uuid.New().String()
	now := time.Now()
	product.CreatedAt = now
	product.UpdatedAt = now

	query := `
		INSERT INTO products (
			id, user_id, title, image, description, price, stock, created_at, updated_at
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`

	_, err := r.db.Exec(query,
		product.ID, product.UserID, product.Title, product.Image,
		product.Description, product.Price, product.Stock, product.CreatedAt, product.UpdatedAt)

	if err != nil {
		return err
	}

	return nil
}

// GetPrdByEmail gets a product by email
func (r *productRepo) GetPrdByEmail(userID string) (*models.Product, error) {
	product := &models.Product{}
	query := `
		SELECT id, user_id, title, image, description, price, stock, created_at, updated_at
		FROM products WHERE user_id = ?`

	err := r.db.QueryRow(query, userID).Scan(
		&product.ID, &product.UserID, &product.Title, &product.Image,
		&product.Description, &product.Price, &product.Stock, &product.CreatedAt, &product.UpdatedAt)

	if err == sql.ErrNoRows {
		return nil, errors.New("product not found")
	}
	if err != nil {
		return nil, err
	}

	return product, nil
}

// GetProdByID gets a product by id
func (r *productRepo) GetProdByID(id string) (*models.Product, error) {
	product := &models.Product{}
	query := `
		SELECT id, user_id, title, image, description, price, stock, created_at, updated_at
		FROM products WHERE id = ?`

	err := r.db.QueryRow(query, id).Scan(
		&product.ID, &product.UserID, &product.Title, &product.Image,
		&product.Description, &product.Price, &product.Stock, &product.CreatedAt, &product.UpdatedAt)

	if err == sql.ErrNoRows {
		return nil, errors.New("product not found")
	}
	if err != nil {
		return nil, err
	}

	return product, nil
}

// UpdateProd updates a product
func (r *productRepo) UpdateProd(product *models.Product) error {
	product.UpdatedAt = time.Now()
	query := `
		UPDATE products 
		SET user_id = ?, title = ?, image = ?, description = ?, price = ?, stock = ?, updated_at = ?
		WHERE id = ?`

	result, err := r.db.Exec(query,
		product.UserID, product.Title, product.Image, product.Description, product.Price,
		product.Stock, product.UpdatedAt, product.ID)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("product not found")
	}

	return nil
}

// DeleteProd deletes a product
func (r *productRepo) DeleteProd(id string) error {
	query := `DELETE FROM products WHERE id = ?`

	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("product not found")
	}

	return nil
}

// ListProducts lists all products
func (r *productRepo) ListProducts() ([]*models.Product, error) {
	query := `
		SELECT id, user_id, title, image, description, price, stock, created_at, updated_at
		FROM products`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*models.Product
	for rows.Next() {
		product := &models.Product{}
		err = rows.Scan(
			&product.ID, &product.UserID, &product.Title, &product.Image, &product.Description,
			&product.Price, &product.Stock, &product.CreatedAt, &product.UpdatedAt)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}
