# Routes Package

The Routes package is responsible for handling the HTTP routing for the application. It defines the routes that connect incoming HTTP requests to the corresponding handler functions in the application, allowing for various operations, such as user authentication and product management.

## Overview

This package utilizes the `mux` router from the Gorilla Toolkit to manage routes and implements CORS (Cross-Origin Resource Sharing) to facilitate requests from different origins. It organizes the routes for the application's API, enabling external clients to interact with the service effectively.

## Router

### Router Structure

The `Router` struct holds the main routing setup, including a mux router and a handler for managing the application logic.

```go
type Router struct {
    muxRouter *mux.Router
    handler   *handlers.Handler
}