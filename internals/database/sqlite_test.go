package database

import (
	"testing"

	"github.com/antmusumba/agrinet/internals/models"
	"github.com/antmusumba/agrinet/internals/repositories"
)

func TestInsertProduct(t *testing.T) {
	db, err := InitDB(":memory:")
	if err != nil {
		t.Fatalf("Failed to initialize DB: %v", err)
	}

	// Seed user
	_, err = db.Exec(`INSERT INTO users (id, email, password) VALUES ('test-user-id', 'test@example.com', 'hashed-password')`)
	if err != nil {
		t.Fatalf("Failed to seed user: %v", err)
	}

	// Attempt to insert product
	product := &models.Product{
		UserID:      "test-user-id",
		Title:       "Test Product",
		Description: "Test Description",
		Price:       100,
		Stock:       10,
	}
	repo := repositories.NewProductRepo(db)
	err = repo.CreateProd(product)
	if err != nil {
		t.Fatalf("Failed to insert product: %v", err)
	}
}
