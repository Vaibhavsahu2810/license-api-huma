# License API – Powered by Huma

A RESTful API in Go for managing software licenses, built using the [Huma](https://github.com/danielgtaylor/huma) OpenAPI 3.1-compliant framework. This project demonstrates modern, code-first API development with auto-generated documentation and strong typing.

## Tech Stack

- **Go** 1.21+
- **[Huma](https://github.com/danielgtaylor/huma)** – A fast, modern OpenAPI 3.1 framework
- **Gin** – Lightweight web framework for Go
- **Swagger UI** – Integrated interactive API docs

## Features

- Code-first OpenAPI 3.1 generation (also compatible with 3.0.3)
- Gin-compatible using `humagin` adapter
- Built-in Swagger UI available at `http://localhost:8080/docs`
- Dynamic server configuration via `.env` (host, port, etc.)
- CRUD endpoints: Create and Update licenses
- Type-safe request/response models

## Getting Started

### 1. Clone the repository

```bash
git clone https://github.com/Vaibhavsahu2810/huma-license-api.git
cd huma-license-api
```

### Install dependencies

go mod tidy

### Add a .env file (optional)

HOST=localhost
PORT=8080

### Run the server

go run main.go

### Access Swagger UI

Open your browser and go to:
http://localhost:8080/docs
