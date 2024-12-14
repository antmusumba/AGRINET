# Server Package

The Server package is responsible for initializing and running the HTTP server for the application, handling incoming requests and providing an interface for graceful shutdown and command management.

## Overview

This package creates and configures an HTTP server using Go's `net/http` package, and integrates with the routes package to define the application's API endpoints. It supports graceful shutdown, ensuring that ongoing requests are completed before the server stops.

## Server

### Server Structure

The `Server` struct manages the HTTP server instance and its associated functionalities.

```go
type Server struct {
    server *http.Server
}