package repositories

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/antmusumba/agrinet/internals/models"
	"github.com/stretchr/testify/assert"
)

func TestCreateProd(t *testing.T) {
	// Create a mock database and mock object
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock db: %v", err)
	}
	defer db.Close()

	// Create an instance of productRepo
	productRepo := NewProductRepo(db)

	// Create a test product
	product := &models.Product{
		UserID:      "user-123",
		Title:       "Test Product",
		Image:       "test.jpg",
		Description: "Test product description",
		Price:       100.0,
		Stock:       10,
	}

	// Setup the expected SQL query and mock behavior
	mock.ExpectExec("INSERT INTO products").
		WithArgs(sqlmock.AnyArg(), product.UserID, product.Title, product.Image,
			product.Description, product.Price, product.Stock, sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(1, 1))

	// Call the CreateProd method
	err = productRepo.CreateProd(product)

	// Assert that there were no errors
	assert.NoError(t, err)

	// Ensure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("there were unmet expectations: %v", err)
	}
}

func TestGetProdByID(t *testing.T) {
	// Create a mock database and mock object
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock db: %v", err)
	}
	defer db.Close()

	// Create an instance of productRepo
	productRepo := NewProductRepo(db)

	// Create a test product
	product := &models.Product{
		ID:          "prod-123",
		UserID:      "user-123",
		Title:       "Test Product",
		Image:       "test.jpg",
		Description: "Test product description",
		Price:       100.0,
		Stock:       10,
	}

	// Setup the expected SQL query and mock behavior
	mock.ExpectQuery("SELECT id, user_id, title, image, description, price, stock, created_at, updated_at").
		WithArgs(product.ID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "user_id", "title", "image", "description", "price", "stock"}).
			AddRow(product.ID, product.UserID, product.Title, product.Image, product.Description, product.Price, product.Stock))

	// Call the GetProdByID method
	result, err := productRepo.GetProdByID(product.ID)

	// Assert that there were no errors and the result matches the expected product
	assert.NoError(t, err)
	assert.Equal(t, product.ID, result.ID)
	assert.Equal(t, product.Title, result.Title)

	// Ensure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("there were unmet expectations: %v", err)
	}
}

func TestUpdateProd(t *testing.T) {
	// Create a mock database and mock object
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock db: %v", err)
	}
	defer db.Close()

	// Create an instance of productRepo
	productRepo := NewProductRepo(db)

	// Create a test product
	product := &models.Product{
		ID:          "prod-123",
		UserID:      "user-123",
		Title:       "Updated Product",
		Image:       "updated.jpg",
		Description: "Updated product description",
		Price:       200.0,
		Stock:       5,
	}

	// Setup the expected SQL query and mock behavior
	mock.ExpectExec("UPDATE products").
		WithArgs(product.UserID, product.Title, product.Image, product.Description, product.Price,
			product.Stock, product.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	// Call the UpdateProd method
	err = productRepo.UpdateProd(product)

	// Assert that there were no errors
	assert.NoError(t, err)

	// Ensure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("there were unmet expectations: %v", err)
	}
}
