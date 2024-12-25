# README.md

## User Management API

This Go application provides a simple HTTP server to manage users, including functions for creating and listing users.

### Features

- HTTP server running on port 8080
- Two endpoints:
  - `GET /`: Lists all users
  - `POST /users`: Creates a new user

### User Structure

Users have the following fields:
- `Name`: Required field for user's name
- `Description`: User's description
- `Age`: User's age

### Installation

1. Clone the repository.
2. Navigate to the project directory.
3. Run the server using:
   ```bash
   go run main.go
   ```

### Endpoints

**GET /**

- Responds with a list of all users in the format:
  ```
  {Name: "name", Description: "description", Age: "age"}
  ```

**POST /users**

- Creates a new user. Expects a JSON object in the body with `name`, `description`, and `age`.
- Returns a `400 Bad Request` if `name` is missing.

### Concurrency

- Handles concurrent requests using a mutex lock ensuring that user data is accessed safely.

### Dependencies

- Uses Go's standard library packages:
  - `net/http` for HTTP server
  - `encoding/json` for JSON parsing
  - `fmt` for formatted I/O
  - `sync` for thread-safe operations

### Example Usage

**To create a user:**

```bash
curl -X POST http://localhost:8080/users -d '{"name":"John Doe", "description":"A cool guy", "age":"30"}' -H 'Content-Type: application/json'
```

**To list users:**

```bash
curl http://localhost:8080/
```
