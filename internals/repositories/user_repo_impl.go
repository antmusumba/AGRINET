package repositories

import (
	"database/sql"
	"errors"
	"log"
	"time"

	"github.com/antmusumba/agrinet/internals/models"
	"github.com/google/uuid"
)

type userRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) UserRepo {
	// Test database connection
	err := db.Ping()
	if err != nil {
		log.Printf("Database connection error: %v", err)
	} else {
		log.Println("Database connected successfully")
	}
	return &userRepo{db: db}
}

func (r *userRepo) CreateUser(user *models.User) error {
	user.ID = uuid.New().String()
	now := time.Now()
	user.CreatedAt = now
	user.UpdatedAt = now

	query := `
		INSERT INTO users (
			id, email, password, first_name, last_name,
			phone, address, image, role, status,
			verified, votes, created_at, updated_at
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	_, err := r.db.Exec(query,
		user.ID, user.Email, user.Password, user.FirstName, user.LastName,
		user.Phone, user.Address, user.Image, user.Role, user.Status,
		user.Verified, user.Votes, user.CreatedAt, user.UpdatedAt)

	if err != nil {
		return err
	}

	return nil
}
