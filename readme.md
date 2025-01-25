
# A Simple Key-Value in Memory Storage Database

![Go](https://img.shields.io/badge/Go-1.21-blue)
![License](https://img.shields.io/badge/License-MIT-green)
[![Postman](https://img.shields.io/badge/Tested%20with-Postman-orange)](https://www.postman.com/)

A lightweight, high-performance, in-memory key-value store built with Go. This project provides a simple REST API to perform basic operations like `Set`, `Get`, and `Delete` on key-value pairs. It's perfect for learning Go, building small-scale applications, or as a starting point for more complex projects.

---


---

## Why This Project?

- **Learn Go**: A great way to understand Go's concurrency model, HTTP server, and package organization.
- **Lightweight**: No external dependencies—just pure Go.
- **Extensible**: Easily extendable to add features like persistence, transactions, or advanced querying. (Hopefully will be able to implement that in future.)
- **Production-Ready**: Built with best practices like concurrency safety and clean code structure.

---


## Features

- **In-Memory Storage**: Fast and efficient key-value storage using Go's native `map` and `sync.RWMutex` for concurrency safety.
- **REST API**: Exposes a simple HTTP API to interact with the database.
- **Concurrency Safe**: Built with `sync.RWMutex` to handle concurrent read/write operations.
- **Easy to Use**: Simple and intuitive API endpoints for `Set`, `Get`, and `Delete` operations.
- **Lightweight**: No external dependencies—just pure Go.

---

## Getting Started

### Prerequisites

- Go 1.21 or higher installed on your machine.
- Basic knowledge of HTTP and REST APIs.
- (Optional) Postman or any HTTP client for testing.

---

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/nomannaq/simple-database-golang
   ```

2. Navigate to the project directory:
   ```bash
   cd my-database-project
   ```

3. Build and run the project:
   ```bash
   go run cmd/my-database/main.go
   ```

4. The server will start on `http://localhost:8000`.

---

## API Endpoints

### 1. Set a Key-Value Pair

- **Method**: `POST`
- **URL**: `/set`
- **Request Body** (JSON):
  ```json
  {
    "key": "name",
    "value": "John Doe"
  }
  ```
- **Response**:
  ```
  Key 'name' set successfully
  ```

### 2. Get a Value by Key

- **Method**: `GET`
- **URL**: `/get?key=<key>`
- **Example**: `GET /get?key=name`
- **Response** (JSON):
  ```json
  {
    "name": "John Doe"
  }
  ```

### 3. Delete a Key-Value Pair

- **Method**: `DELETE`
- **URL**: `/delete?key=<key>`
- **Example**: `DELETE /delete?key=name`
- **Response**:
  ```
  Key 'name' deleted successfully
  ```

---

## Example Usage

### Using `curl`

1. Set a key-value pair:
   ```bash
   curl -X POST -H "Content-Type: application/json" -d '{"key": "name", "value": "John Doe"}' http://localhost:8000/set
   ```

2. Get a value by key:
   ```bash
   curl http://localhost:8000/get?key=name
   ```

3. Delete a key-value pair:
   ```bash
   curl -X DELETE http://localhost:8000/delete?key=name
   ```

### Using Postman

1. Import the provided Postman collection (link to collection).
2. Use the pre-configured requests to test the API.

---

## Project Structure

```
my-database-project/
├── cmd/
│   └── main/
│       └── main.go            # Entry point of the application
├── internal/
│   ├── database/
│   │   └── database.go        # Core database logic
│   └── handlers/
│       └── handlers.go        # HTTP handlers for the API
├── go.mod                     # Go module file
├── go.sum                     # Dependency checksum file
└── README.md                  # This file
```


## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---

## Acknowledgments

- Built with ❤️ using Go.
- Inspired by simple key-value stores like Redis.
- Special thanks to the Go community for amazing resources and tools.

---
