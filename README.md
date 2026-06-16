# Task API

A simple RESTful API built with Go, following Clean Architecture principles. This project was created as a **learning/practice example** to understand how to structure a Go backend application using separated layers (domain, repository, usecase, handler).

## About

This is a sample Task Management API demonstrating:
- Clean Architecture pattern (domain → repository → usecase → handler → route)
- RESTful API design with versioned endpoints
- MongoDB integration using the official Go driver
- JWT-based authentication (in progress)
- Dependency injection via interfaces

> **Note:** This project is intended for learning purposes and is not production-hardened.

## Tech Stack

- **Language:** Go 1.22+
- **Web Framework:** [Gin](https://github.com/gin-gonic/gin)
- **Database:** MongoDB
- **Driver:** [mongo-driver](https://github.com/mongodb/mongo-go-driver) v1.17
- **Auth:** JWT, bcrypt for password hashing
- **Config:** godotenv
- **Containerization:** Docker Compose (for local MongoDB)

## Project Structure

```
task-api/
├── cmd/
│   └── server/
│       └── main.go              # Application entry point
├── database/
│   └── database.go              # MongoDB connection setup
├── internal/
│   ├── domain/                  # Core business entities (pure structs)
│   │   ├── task.go
│   │   └── user.go
│   ├── repository/              # Data access layer (DB queries)
│   │   ├── task_repository.go
│   │   └── user_repository.go
│   ├── usecase/                 # Business logic layer
│   │   ├── task_usecase.go
│   │   └── auth_usecase.go
│   ├── handler/                 # HTTP request/response layer
│   │   ├── task_handler.go
│   │   └── auth_handler.go
│   ├── middleware/
│   │   ├── auth.go
│   │   └── logger.go
│   └── route/
│       └── route.go             # Route registration
├── pkg/
│   └── response/
│       └── response.go          # Standardized API response format
├── migrations/
├── docker-compose.yml
├── .env
└── go.mod
```

### Architecture Flow

```
Request → route → middleware → handler → usecase → repository → database
```

Each layer has a single responsibility:
- **handler** — parses HTTP requests and returns responses (no business logic)
- **usecase** — contains business rules and decision-making
- **repository** — handles database queries only
- **domain** — plain structs shared across all layers

## Getting Started

### Prerequisites

- Go 1.22 or higher
- Docker & Docker Compose

### 1. Clone the repository

```bash
git clone <your-repo-url>
cd task-api
```

### 2. Set up environment variables

Create a `.env` file in the root directory:

```env
MONGO_URI=mongodb://admin:password@localhost:27017
MONGO_DB_NAME=task_db
PORT=8080
JWT_SECRET=your_jwt_secret_here
```

### 3. Start MongoDB with Docker Compose

```bash
docker-compose up -d
```

### 4. Install dependencies

```bash
go mod tidy
```

### 5. Run the server

```bash
go run cmd/server/main.go
```

The server should start on `http://localhost:8080`.

## API Endpoints

Base path: `/api/v1`

### Health Check

| Method | Endpoint | Description |
|--------|----------|--------------|
| GET | `/api/v1/health` | Check if the server is running |

### Auth

| Method | Endpoint | Description |
|--------|----------|--------------|
| POST | `/api/v1/auth/register` | Register a new user |
| POST | `/api/v1/auth/login` | Login (in progress) |

### Users

| Method | Endpoint | Description |
|--------|----------|--------------|
| GET | `/api/v1/users/:id` | Get user by ID |

### Tasks

| Method | Endpoint | Description |
|--------|----------|--------------|
| GET | `/api/v1/tasks` | Get all tasks |
| GET | `/api/v1/tasks/:id` | Get a task by ID |
| POST | `/api/v1/tasks` | Create a new task |
| PUT | `/api/v1/tasks/:id` | Update a task |
| DELETE | `/api/v1/tasks/:id` | Delete a task |

## Example Request

```bash
curl -X POST http://localhost:8080/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "email": "john@example.com",
    "password": "secret123"
  }'
```

## Response Format

**Success**
```json
{
  "status": "success",
  "message": "Request successful",
  "data": { },
  "meta": { "timestamp": "2026-06-16T10:30:00Z" }
}
```

**Error**
```json
{
  "status": "error",
  "message": "Invalid request body",
  "errors": "...",
  "meta": { "timestamp": "2026-06-16T10:30:00Z" }
}
```

## Roadmap

- [x] MongoDB connection
- [x] User domain, repository, usecase, handler
- [x] Task CRUD (basic)
- [ ] JWT login & auth middleware
- [ ] Pagination & filtering for task list
- [ ] Unit tests

## License

This project is for educational purposes only.
