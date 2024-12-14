# Database Package

This package provides a simple interface for initializing and managing an SQLite database in a Go application. It focuses on establishing a secure connection to the database, enabling foreign key support, and creating necessary tables for user and product data.

## Features

- **SQLite Connection**: Establishes a connection to an SQLite database using a provided data source name.
- **Foreign Keys Support**: Enables foreign key constraints for referential integrity.
- **Table Creation**:
  - Creates a `users` table to store user-related information.
  - Creates a `products` table to store product-related information, linking it to the users table via foreign key.

## Installation

To use this package, ensure you have Go installed and the SQLite3 driver. You can install the driver using the following command:

```bash
go get -u github.com/mattn/go-sqlite3