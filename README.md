# Go-Todo Application

This project is a simple **Todo** application built in Go, following **Clean Architecture** principles. The goal of this project is to demonstrate the separation of concerns and organization of the code in a way that makes it scalable, testable, and easy to maintain.

The project is designed to serve as a learning resource for colleagues who want to understand Clean Architecture, Dependency Injection, and the structure of a well-organized Go application.

---

## Project Structure

The project is structured according to **Clean Architecture**, where different layers of the application are separated into packages with a clear responsibility. Here's an overview of the directory structure:
```
.
├── cmd
│   └── main.go                # Entry point of the application
├── config
│   └── config.go              # Loads environment variables and configuration
├── go.mod
├── go.sum
├── internal
│   ├── domain                 # Core business logic
│   │   ├── models             # Domain models (e.g., Todo)
│   │   ├── repositories       # Interfaces for data storage
│   │   └── services           # Business logic (e.g., TodoService)
│   ├── infastructure          # Infrastructure code (e.g., database connections)
│   │   └── database
│   │       └── rds            # RDS database connection setup
│   ├── repository             # Repository implementation (e.g., TodoRepository)
│   └── transport              # HTTP and gRPC transport layers
│       ├── grpc
│       └── http
│           ├── handlers       # HTTP handlers (e.g., TodoHandler)
│           └── routers        # HTTP router setup
├── pkg                        # Reusable libraries and utilities (if any)
└── test                       # Test files
```

### Key Directories:
1. **cmd**: Contains the entry point `main.go`, which starts the application.
2. **config**: Holds the configuration logic and environment variable loading.
3. **internal**: Contains the core business logic, repositories, services, and transport layers (HTTP, gRPC).
4. **pkg**: Any reusable code or libraries.
5. **test**: Test files for unit and integration tests.

---

## Core Concepts

### 1. **Clean Architecture**
Clean Architecture emphasizes separation of concerns, where:
- The **domain** layer contains the core business logic (models and services).
- The **repository** layer defines the interfaces for persistence but does not implement them.
- The **infrastructure** layer implements the repository interfaces (e.g., database connection).
- The **transport** layer is responsible for handling HTTP requests and responses, mapping them to the business logic.

This separation allows for easier testing, as dependencies can be injected and replaced without affecting the core logic.

### 2. **Dependency Injection**
The project uses **dependency injection** to inject the required dependencies into each layer. This allows components to remain decoupled and makes the system easier to extend and test.

---

## Setup

### Prerequisites

Before running the application, make sure you have the following installed on your system:
- Go 1.20+ (Install Go from [here](https://go.dev/dl/))
- PostgreSQL or another relational database (for example, use Docker to set up a database)
- `.env` file (used to store sensitive information such as database credentials)

### Installation

1. **Clone the repository**:

   ```bash
   git clone https://github.com/mehedicode-lab/go-todo.git
   cd go-todo
   ```
2. Install Dependencies

The project uses Go modules for dependency management. Run the following command to install the required dependencies:

```bash
  go mod tidy
```
3. Create .env file:

Inside the config/ directory, create an .env file with the following contents:

4. Run the application:

After the setup is complete, you can run the application using:
```sh
  go run cmd/main.go
```
The application will start the server on the specified port (default is 8080). The server will connect to the PostgreSQL database specified in the .env file.

## API Endpoints

The application exposes the following endpoints for performing CRUD operations on Todo items:

| Method | Endpoint            | Description                              |
|--------|---------------------|------------------------------------------|
| **GET**    | `/todos`            | Retrieves all Todo items                |
| **GET**    | `/todos/{id}`       | Retrieves a single Todo item by its ID  |
| **POST**   | `/todos`            | Creates a new Todo item                 |
| **PUT**    | `/todos/{id}`       | Updates an existing Todo item by its ID |
| **DELETE** | `/todos/{id}`       | Deletes a Todo item by its ID           |

---

### Example Requests

#### 1. **Create a Todo**
**Request**:
```http
POST /todos
Content-Type: application/json

{
  "title": "Learn Clean Architecture",
  "completed": false
}
