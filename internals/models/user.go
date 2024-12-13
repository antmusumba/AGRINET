package models

import "time"

// User represents a user
type User struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Address   string    `json:"address"`
	Password  string    `json:"-"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Image     string    `json:"image"`
	Role      string    `json:"role"`
	Status    string    `json:"status"`
	Verified  bool      `json:"verified"`
	Votes     int       `json:"votes"`
}
