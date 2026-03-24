# Golang Bootstrap Project

A (maybe) robust and scalable Golang backend bootstrap using **Fiber** and **GORM**, following **Clean Architecture** and **Domain-Driven Design (DDD)** principles.

## Architecture Overview

This project implements a modular Clean Architecture where each domain/feature is encapsulated within its own module. This ensures high maintainability, testability, and clear separation of concerns.

### Project Structure

```text
├── cmd/
│   ├── app/                # Main application entry point
│   ├── migrate/            # Database migration scripts
│   └── seed/               # Database seeding scripts
├── config/                 # Configuration management
├── internal/
│   ├── infrastructure/     # External tools & frameworks (DB, etc.)
│   ├── modules/            # Core business modules (DDD Layers)
│   │   └── [module-name]/
│   │       ├── contract/   # Interfaces for decoupling
│   │       ├── entity/     # Domain models
│   │       ├── handler/    # HTTP/Transport layer
│   │       ├── repository/ # Data access layer
│   │       ├── service/    # Business logic layer
│   │       └── routes/     # Module-specific routing
│   ├── shared/             # Shared contracts and utilities
│   └── token/              # JWT and Token handling
├── middleware/             # HTTP middlewares (Auth, Logger, etc.)
└── utils/                  # Common helper functions
```

## Tech Stack

- **Framework**: [Fiber v2](https://gofiber.io/) (High-performance web framework)
- **ORM**: [GORM](https://gorm.io/) (Nice-simple ORM library for Golang)
- **Database**: PostgreSQL
- **Authentication**: JWT (JSON Web Token)
- **Migrations**: golang-migrate
- **Configuration**: godotenv

## Getting Started

### Prerequisites

- Go 1.24+
- PostgreSQL
- [Optional] Podman or Docker

### Installation

1.  **Clone the repository**:
    ```bash
    git clone https://github.com/waldin-exe/golang-bootstrap.git
    cd golang-bootstrap
    ```

2.  **Setup environment variables**:
    ```bash
    cp .env.example .env
    # Edit .env with your database credentials
    ```

3.  **Install dependencies**:
    ```bash
    go mod tidy
    ```

### Running the Application

- **Development Mode**:
  ```bash
  go run cmd/app/main.go
  ```

- **Run Migrations**:
  ```bash
  go run cmd/migrate/main.go
  ```

- **Run Seeder**:
  ```bash
  go run cmd/seed/main.go
  ```

## Key Features

- **Modular DDD Structure**: Scale your project easily by adding new modules in `internal/modules`.
- **Clean Architecture**: Decoupled layers using interfaces (contracts).
- **Automated Migrations**: Easy database schema management.
- **JWT Authentication**: Secure API endpoints with built-in JWT middleware.
- **Centralized Middleware**: Global and route-specific middleware support (CORS, Logger, etc.).
- **Response Utilities**: Consistent API response format.

## License

This project is licensed under the MIT License - see the LICENSE file for details.

If you find this project helpful, please support me with your prayers. 
I hope to have the opportunity to go to Japan soon, either to work or study. Thank you :D