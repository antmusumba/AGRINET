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

func (r *userRepo) GetUserByEmail(email string) (*models.User, error) {
	user := &models.User{}
	query := `
		SELECT id, email, password, first_name, last_name,
			phone, address, image, role, status,
			verified, votes, created_at, updated_at
		FROM users WHERE email = ?`

	err := r.db.QueryRow(query, email).Scan(
		&user.ID, &user.Email, &user.Password, &user.FirstName, &user.LastName,
		&user.Phone, &user.Address, &user.Image, &user.Role, &user.Status,
		&user.Verified, &user.Votes, &user.CreatedAt, &user.UpdatedAt)

	if err == sql.ErrNoRows {
		return nil, errors.New("user not found")
	}
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *userRepo) GetUserByID(id string) (*models.User, error) {
	user := &models.User{}
	query := `
		SELECT id, email, password, first_name, last_name,
			phone, address, image, role, status,
			verified, votes, created_at, updated_at
		FROM users WHERE id = ?`

	log.Printf("Looking up user by id: %s", id)
	err := r.db.QueryRow(query, id).Scan(
		&user.ID, &user.Email, &user.Password, &user.FirstName, &user.LastName,
		&user.Phone, &user.Address, &user.Image, &user.Role, &user.Status,
		&user.Verified, &user.Votes, &user.CreatedAt, &user.UpdatedAt)

	if err == sql.ErrNoRows {
		log.Printf("User not found with id: %s", id)
		return nil, errors.New("user not found")
	}
	if err != nil {
		log.Printf("Error getting user by id: %v", err)
		return nil, err
	}

	log.Printf("Found user with id: %s", id)
	return user, nil
}

func (r *userRepo) UpdateUser(user *models.User) error {
	user.UpdatedAt = time.Now()
	query := `
		UPDATE users 
		SET email = ?, password = ?, first_name = ?, last_name = ?,
			phone = ?, address = ?, image = ?, role = ?, status = ?,
			verified = ?, votes = ?, updated_at = ?
		WHERE id = ?`

	log.Printf("Updating user with id: %s", user.ID)
	result, err := r.db.Exec(query,
		user.Email, user.Password, user.FirstName, user.LastName,
		user.Phone, user.Address, user.Image, user.Role, user.Status,
		user.Verified, user.Votes, user.UpdatedAt, user.ID)

	if err != nil {
		log.Printf("Error updating user: %v", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Error getting rows affected: %v", err)
		return err
	}
	if rowsAffected == 0 {
		log.Printf("User not found with id: %s", user.ID)
		return errors.New("user not found")
	}

	log.Printf("User updated successfully. Rows affected: %d", rowsAffected)
	return nil
}

func (r *userRepo) DeleteUser(id string) error {
	query := `DELETE FROM users WHERE id = ?`

	log.Printf("Deleting user with id: %s", id)
	result, err := r.db.Exec(query, id)
	if err != nil {
		log.Printf("Error deleting user: %v", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Error getting rows affected: %v", err)
		return err
	}
	if rowsAffected == 0 {
		log.Printf("User not found with id: %s", id)
		return errors.New("user not found")
	}

	log.Printf("User deleted successfully. Rows affected: %d", rowsAffected)
	return nil
}
