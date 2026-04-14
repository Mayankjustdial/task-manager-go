# Task Management API (Go + DDD)

A production-ready Task Management REST API built using Golang, following Domain-Driven Design (DDD), Test-Driven Development (TDD), and clean architecture principles.

---

## 🚀 Features

* Create, update, retrieve, and delete tasks
* Domain-driven design (DDD architecture)
* In-memory data persistence
* Input validation (title and due date)
* Pagination and filtering support
* Unit and integration tests
* Docker support

---

## 📦 Prerequisites

Ensure the following are installed:

* Go (>= 1.22)
* Git
* Docker (optional)

Check installation:

```bash
go version
docker --version
```

---

## 📥 Clone the Repository

```bash
git clone <your-github-repo-url>
cd task-manager
```

---

## 📦 Install Dependencies

```bash
go mod tidy
```

This installs:

* Gin (HTTP framework)
* UUID generator
* Testify (testing library)

---

## ▶️ Run the Application

```bash
go run cmd/api/main.go
```

Server runs at:

```
http://localhost:8080
```

---

## 🐳 Run with Docker

### Build image

```bash
docker build -t task-api .
```

### Run container

```bash
docker run -p 8080:8080 task-api
```

---

## 📬 API Endpoints

### Create Task

POST /tasks

### Get Task

GET /tasks/{id}

### Update Task

PUT /tasks/{id}

### Delete Task

DELETE /tasks/{id}

### List Tasks

GET /tasks?limit=10&offset=0&status=PENDING

---

## 🧪 Running Tests

```bash
go test ./...
```

---

## 📌 Example Request

### Create Task

```json
{
  "title": "Learn Golang",
  "description": "DDD project",
  "due_date": "2026-04-20T10:00:00Z"
}
```

---

## 🏗️ Project Structure

```
task-manager/
├── cmd/api            # Entry point
├── internal/
│   ├── domain         # Business logic
│   ├── repository     # Data layer
│   ├── service        # Use cases
│   ├── handler        # HTTP layer
│   └── dto            # Request/response models
├── tests/             # integration tests
```

---

## 💡 Design Decisions

* Domain layer is independent of frameworks
* DTO prevents leaking internal domain logic
* Repository pattern enables easy database replacement
* In-memory store used for simplicity

---
