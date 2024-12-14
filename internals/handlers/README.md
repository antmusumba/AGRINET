# Handlers Package

The Handlers package provides HTTP handlers for managing user authentication and product operations in a Go application. This package is part of a larger system, typically a web service, and interacts with the internal services that handle business logic and database operations.

## Features

- **Health Check**: A simple endpoint to check if the service is running.
- **User Registration**: Endpoint for registering new users.
- **User Login**: Endpoint for user authentication with JWT token generation.
- **Product Management**: CRUD operations for products, including creation and retrieval of all products.

## Installation

To use this package, ensure you have Go installed on your machine and add the necessary dependencies. You might already have relevant packages included in your project. If not, install them:

```bash
go get -u github.com/antmusumba/agrinet/internals/models
go get -u github.com/antmusumba/agrinet/internals/services