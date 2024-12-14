# Repositories Package

The Repositories package contains data access layer functionalities for managing `User` and `Product` entities. This package interacts with the database to perform various operations like creating, reading, updating, and deleting records. 

## Overview

This package defines two main repository interfaces, `UserRepo` and `ProductRepo`, along with their corresponding implementations. These repositories facilitate CRUD (Create, Read, Update, Delete) operations on users and products.

## Interfaces

### UserRepo

The `UserRepo` interface provides methods for handling user-related data operations.

```go
type UserRepo interface {
    CreateUser(user *models.User) error
    GetUserByEmail(email string) (*models.User, error)
    GetUserByID(id string) (*models.User, error)
    UpdateUser(user *models.User) error
    DeleteUser(id string) error
}