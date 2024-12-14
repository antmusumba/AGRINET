# Models Package

The Models package defines the data structures used in the application. It contains definitions for important entities such as `User` and `Product`. These structures represent the core domain models and are used throughout the application.

## Structures

This package defines two primary data structures: `User` and `Product`.

### User

The `User` struct represents a user in the system. It contains fields to capture user details and their current status.

```go
type User struct {
    ID        string    `json:"id"`         // Unique identifier for the user
    Email     string    `json:"email"`      // User's email address
    Phone     string    `json:"phone"`      // User's phone number
    Address   string    `json:"address"`    // User's physical address
    Password  string    `json:"password"`   // User's hashed password
    FirstName string    `json:"firstName"`  // User's first name
    LastName  string    `json:"lastName"`   // User's last name
    CreatedAt time.Time `json:"createdAt"`  // Timestamp when user was created
    UpdatedAt time.Time `json:"updatedAt"`  // Timestamp when user was last updated
    Image     string    `json:"image"`      // URL of the user's profile image
    Role      string    `json:"role"`       // User's role in the system (e.g., admin, user)
    Status    string    `json:"status"`     // User's current status (e.g., active, inactive)
    Verified  bool      `json:"verified"`   // Indicates if the user has verified their account
    Votes     int       `json:"votes"`      // Count of votes received by the user (if applicable)
}