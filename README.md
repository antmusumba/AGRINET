# AGRINET

## Problem Statement:

Farmers, particularly those in rural and underserved regions, face significant barriers in accessing essential agricultural inputs such as high-quality seedlings, fertilizers, breeds and other farming resources. This lack of access not only hampers their ability to adopt best practices in crop production but also limits their overall productivity and economic potential. The absence of quality inputs often results in poor crop yields, increased susceptibility to pests and diseases, and diminished sustainability of farming operations. As a result, farmers struggle to meet the growing demand for food while remaining financially viable, thereby exacerbating food insecurity and undermining rural economic development.


 ## Agrinet: Bridging the Gap Between Farmers and Buyers of Rare Agricultural Products

 ## Overview:
Agrinet is an innovative digital platform designed to create a seamless connection between farmers and buyers seeking rare agricultural products, such as specialized seedlings, breeds, and essential farm inputs. By leveraging technology, Agrinet aims to enhance market access for farmers while providing buyers with high-quality and unique agricultural offerings.

 Key Features:

## User-Friendly Interface:
    Agrinet features an intuitive interface that allows both farmers and buyers to navigate the platform with ease. Users can quickly access information, post listings, and make transactions, ensuring a smooth experience for all participants.

## Product Listings:
    Farmers can create detailed listings for their rare agricultural products, including comprehensive descriptions, images, and pricing information. This feature not only showcases the uniqueness of their offerings but also assists buyers in making informed purchasing decisions.
## Buyer Profiles:
    The platform enables buyers to create profiles to enhance transparency and trust. Buyers can share their specific requirements and preferences, allowing farmers to tailor their offerings accordingly.

## Search and Filtering Options:
    Agrinet incorporates advanced search and filtering functionalities, enabling buyers to easily locate the specific products they require. This feature aids in promoting the visibility of rare agricultural items and enhances the likelihood of successful transactions.

## Secure Payment Gateway:
    To ensure a safe and reliable transaction process, Agrinet integrates a secure payment gateway that facilitates various payment options. This feature instills confidence in users and promotes a hassle-free buying experience.

    ### Installation
1. Prerequisites

    Install Go (version 1.18 or higher).


2. git clone: https://github.com/antmusumba/AgriNet.git


## Usage

Once you’ve cloned the repository, navigate into the project folder:

```bash
cd AGRINET
```

Then, fire it up with:

```bash
go run .
```

Open your browser and paste the following link to get started: `http://localhost:8080`. 

# Frontend

This project was generated using [Angular CLI](https://github.com/angular/angular-cli) version 19.0.5.

## Development server

To start a local development server, run:

```bash
ng serve
```

Once the server is running, open your browser and navigate to `http://localhost:4200/`. The application will automatically reload whenever you modify any of the source files.

## Code scaffolding

Angular CLI includes powerful code scaffolding tools. To generate a new component, run:

```bash
ng generate component component-name
```

For a complete list of available schematics (such as `components`, `directives`, or `pipes`), run:

```bash
ng generate --help
```

## Building

To build the project run:

```bash
ng build
```

This will compile your project and store the build artifacts in the `dist/` directory. By default, the production build optimizes your application for performance and speed.

## Running unit tests

To execute unit tests with the [Karma](https://karma-runner.github.io) test runner, use the following command:

```bash
ng test
```

## Running end-to-end tests

For end-to-end (e2e) testing, run:

```bash
ng e2e
```

Angular CLI does not come with an end-to-end testing framework by default. You can choose one that suits your needs.

## Additional Resources

For more information on using the Angular CLI, including detailed command references, visit the [Angular CLI Overview and Command Reference](https://angular.dev/tools/cli) page.

## TECH
# Database Package

The `database` package provides an abstraction layer for initializing and managing an SQLite database connection in a Go application. It includes functionality for creating necessary tables and handling database interactions.

## Contents

- [Installation](#installation)
- [Usage](#usage)
- [Database Schema](#database-schema)
- [Error Handling](#error-handling)
- [License](#license)

## Installation

To use the `database` package, ensure you have Go installed on your system. You can import the package into your project using Go modules.

```bash
go get github.com/mattn/go-sqlite3
```
# Handlers Package

The `handlers` package provides HTTP handler functions for user authentication and product management in a Go web application. It leverages the internal services for authentication and product management while managing HTTP requests, responses, and error handling.

## Contents

- [Installation](#installation)
- [Usage](#usage)
- [Endpoints](#endpoints)
  - [Health Check](#health-check)
  - [User Registration](#user-registration)
  - [User Login](#user-login)
  - [Create Product](#create-product)
  - [List Products](#list-products)
- [Error Handling](#error-handling)
- [License](#license)

## Installation

Ensure that you have Go installed on your system and that the package dependencies are available in your project.

To install this package, you can use:

```bash
go get path/to/your/repo/handlers
```
# Handlers Package - Response Management

The `handlers` package contains structures and methods for managing JSON responses, including success and error responses, in a Go web application. This package is designed to facilitate consistent handling of HTTP responses.

## Contents

- [Installation](#installation)
- [Usage](#usage)
- [Structures](#structures)
  - [SuccessRes](#successres)
  - [ErrorRes](#errorres)
- [Methods](#methods)
  - [WriteError](#writeerror)
  - [ReadJSON](#readjson)
  - [WriteJSON](#writejson)
- [Error Handling](#error-handling)
- [License](#license)

## Installation

To use the `handlers` package in your project, you can import it as follows:

```go
import "path/to/your/repo/handlers"
```
# Server Package

The `server` package provides an implementation of an HTTP server built in Go. It supports graceful shutdown, command-line commands for server management, and integrates routing for handling requests. 

## Contents

- [Installation](#installation)
- [Usage](#usage)
- [Structures](#structures)
  - [Server](#server)
- [Methods](#methods)
  - [NewServer](#newserver)
  - [handleShutdown](#handleshutdown)
  - [Start](#start)
- [Command-Line Commands](#command-line-commands)
- [Graceful Shutdown](#graceful-shutdown)
- [License](#license)

## Installation

To use the `server` package, you need to have Go installed on your machine. First, ensure that your Go workspace is set up, then import the package into your project:

```go
import "path/to/your/repo/server"
```
# Services Package

The `services` package provides business logic and service layer functions for managing products within the application. It interacts with repositories to perform CRUD operations and ensures data integrity by validating inputs.

## Contents

- [Installation](#installation)
- [Usage](#usage)
- [Structures](#structures)
  - [ProductService](#productservice)
- [Methods](#methods)
  - [NewProductService](#newproductservice)
  - [CreateProduct](#createproduct)
  - [GetProductByID](#getproductbyid)
  - [UpdateProduct](#updateproduct)
  - [DeleteProduct](#deleteproduct)
  - [ListProducts](#listproducts)
- [Error Handling](#error-handling)
- [License](#license)

## Installation

To use the `services` package in your project, ensure that you have Go installed on your machine. Then import the package into your project:

```go
import "path/to/your/repo/services"
```
# Services Package - AuthService

The `AuthService` in the `services` package provides authentication functionalities, including user registration and login mechanisms. It interacts with user repositories to handle operations related to user management while ensuring proper input validation and password security.

## Contents

- [Installation](#installation)
- [Usage](#usage)
- [Structures](#structures)
  - [AuthService](#authservice)
- [Methods](#methods)
  - [NewAuthService](#newauthservice)
  - [Register](#register)
  - [Login](#login)
- [Error Handling](#error-handling)
- [License](#license)

## Installation

To use the `AuthService` in your project, ensure that you have Go installed on your machine. Then import the package into your project:

```go
import "path/to/your/repo/services"
```
# Routes Package

The `routes` package is responsible for defining and configuring HTTP routes for the application. It integrates various functionalities such as user authentication and product management by linking handlers to the respective HTTP endpoints.

## Contents

- [Installation](#installation)
- [Usage](#usage)
- [Structures](#structures)
  - [Router](#router)
- [Methods](#methods)
  - [NewRouter](#newrouter)
  - [SetupRoutes](#setuproutes)
- [CORS Configuration](#cors-configuration)
- [License](#license)

## Installation

To use the `routes` package in your project, ensure you have Go installed on your machine. Then, import the package into your project:

```go
import "path/to/your/repo/routes"
```
# Repositories Package - Product Repository

The `repositories` package handles data storage and retrieval for the application, specifically focusing on product-related data management. The `productRepo` interacts with the SQL database to perform CRUD operations on products.

## Contents

- [Installation](#installation)
- [Usage](#usage)
- [Structures](#structures)
  - [ProductRepo](#productrepo)
- [Methods](#methods)
  - [NewProductRepo](#newproductrepo)
  - [CreateProd](#createprod)
  - [GetPrdByEmail](#getprdbyemail)
  - [GetProdByID](#getprodid)
  - [UpdateProd](#updateprod)
  - [DeleteProd](#deleteprod)
  - [ListProducts](#listproducts)
- [Error Handling](#error-handling)
- [License](#license)

## Installation

To use the `repositories` package in your project, ensure you have Go installed on your machine. Then, import the package into your project:

```go
import "path/to/your/repo/repositories"
```
# Repositories Package

The `repositories` package defines interfaces for data management related to users and products within the application. These interfaces outline the CRUD operations needed for managing user and product entities, abstracting the underlying database interactions.

## Contents

- [Installation](#installation)
- [Usage](#usage)
- [Interfaces](#interfaces)
  - [UserRepo](#userrepo)
  - [ProductRepo](#productrepo)
- [Error Handling](#error-handling)
- [License](#license)

## Installation

To use the `repositories` package in your project, ensure you have Go installed on your machine. Then, import the package into your project:

```go
import "path/to/your/repo/repositories"
```
# User Repository

The `userRepo` implementation provides a structured way to manage user-related data storage and retrieval in the application. This includes creating, retrieving, updating, and deleting user records in the database.

## Contents

- [Installation](#installation)
- [Usage](#usage)
- [Initialization](#initialization)
- [Methods](#methods)
  - [CreateUser](#createuser)
  - [GetUserByEmail](#getuserbyemail)
  - [GetUserByID](#getuserbyid)
  - [UpdateUser](#updateuser)
  - [DeleteUser](#deleteuser)
- [Error Handling](#error-handling)
- [Logging](#logging)
- [License](#license)

## Installation

To use the `userRepo` in your project, ensure you have Go installed on your machine. Then, import the `repositories` package into your project:

```go
import "path/to/your/repo/repositories"
```
# Models Package

The `models` package defines the data structures used throughout the application. These models represent the core entities within the system, specifically users and products. They are essential for data manipulation, storage, and retrieval.

## Contents

- [User Model](#user-model)
- [Product Model](#product-model)
- [JSON Serialization](#json-serialization)
- [Usage Examples](#usage-examples)
- [License](#license)

## User Model

The `User` struct represents a user in the application.

### Fields

- **ID**: `string`
  - Unique identifier for the user.
  
- **Email**: `string`
  - Email address of the user.

- **Phone**: `string`
  - Phone number of the user.

- **Address**: `string`
  - Residential address of the user.

- **Password**: `string`
  - Password for user authentication.

- **FirstName**: `string`
  - User's first name.

- **LastName**: `string`
  - User's last name.

- **CreatedAt**: `time.Time`
  - Timestamp when the user was created.

- **UpdatedAt**: `time.Time`
  - Timestamp when the user details were last updated.

- **Image**: `string`
  - URL or path to the user's profile image.

- **Role**: `string`
  - The role of the user (e.g., admin, user).

- **Status**: `string`
  - Current status of the user (e.g., active, inactive).

- **Verified**: `bool`
  - Indicates whether the user has verified their email.

- **Votes**: `int`
  - Number of votes associated with the user, if applicable.

## Product Model

The `Product` struct represents a product in the application.

### Fields

- **ID**: `string`
  - Unique identifier for the product.

- **UserID**: `string`
  - The ID of the user who owns the product.

- **Title**: `string`
  - Title of the product.

- **Image**: `string`
  - URL or path to the product image.

- **Description**: `string`
  - Detailed description of the product.

- **Price**: `int`
  - Price of the product in the smallest currency unit (e.g., cents).

- **CreatedAt**: `time.Time`
  - Timestamp when the product was created.

- **UpdatedAt**: `time.Time`
  - Timestamp when the product details were last updated.

- **Stock**: `int`
  - Quantity available in stock for the product.

## JSON Serialization

Both `User` and `Product` structs are annotated with JSON tags. These tags determine how the fields will be serialized into JSON format when responding to API requests.

For example, a `User` instance serialized to JSON will look like this:

```json
{
  "id": "12345",
  "email": "user@example.com",
  "phone": "1234567890",
  "address": "123 Main St",
  "password": "hashedPassword",
  "firstName": "John",
  "lastName": "Doe",
  "createdAt": "2023-10-23T12:34:56Z",
  "updatedAt": "2023-10-23T12:34:56Z",
  "image": "http://example.com/image.jpg",
  "role": "user",
  "status": "active",
  "verified": true,
  "votes": 5
}