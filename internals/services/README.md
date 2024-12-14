# Services Package

The Services package provides business logic and performs operations related to user authentication and product management. It interacts with the repositories to handle data-related tasks while encapsulating application-specific rules and validation.

## AuthService

The `AuthService` struct manages user authentication, including user registration and login functionalities.

### AuthService Structure

```go
type AuthService struct {
    repo repositories.UserRepo
}