# 🏦 Eticket Service

A backend system for Public Transportation E-Ticketing that operates 24/7 with offline-first capability. Supports prepaid card payments, tariff calculation based on entry and exit terminals, and automatic/manual data synchronization.

---

## 🚀 Tech Stack

- **Golang** 1.23+
- **Fiber** as the HTTP framework
- **Goose** (for database migrations)
- **PostgreSQL** + sqlx + squirrel for database access
- **Docker Compose** for local development
- **Gomock** and **Testify** for unit testing

---

## 📁 Project Structure

```
loan-service/
├── cmd/                # Application entry point (main.go)
├── config/             # Configuration loader
├── constant/           # Global constants (e.g., loan states)
├── handler/            # HTTP handlers (Fiber)
├── migrations/         # migration files
├── model/
│   ├── db/             # Database models
│   └── payload/        # HTTP request/response payloads
├── repository/         # Data access layer using sqlx + squirrel
├── server/             # server setup
├── service/            # Business logic layer
├── Dockerfile.go.1.23
├── config.yaml
├── docker-compose.yml
└── README.md
```

---

## ⚙️ Makefile Commands

This project includes a `Makefile` for simplified development. Key commands include:

### 🧰 Dependency Installer
Automatically installs required CLI tools if not present:
```bash
make ensure-reflex     # Hot reload Go server during development
make ensure-goose      # Database migration tool
make ensure-swagger    # Swagger documentation generator
make ensure-mockgen    # GoMock interface generator
```

### 🔍 Testing & Mocking
```bash
make tidy              # Run go mod tidy
make mock              # Generate mocks from interfaces
make unit-test         # Run unit tests with coverage
make test              # Run mocks + tests together
```

### 🚀 Running the Server
```bash
make run-http          # Run HTTP server with hot reload
make run-worker        # Run Kafka worker with hot reload
make run-all           # Run both HTTP and worker
```

### 📄 API Documentation (Swagger)
```bash
make api-docs
```
Generates Swagger documentation in the `docs/` folder. Access via:

```
http://localhost:8081/swagger/index.html#/
```

### 🔧 Database Migration
```bash
make migrate-up        # Run all up migrations
make migrate-down      # Rollback last migration
```

To create a new migration:
```bash
goose -dir ./migrations create add_column_name sql
```

---

## 🐳 Dockerized Setup

This project is **preconfigured** in both `docker-compose.yml` and `config.yaml`. You don't need to modify anything to get started.

Run the full stack with:

```bash
docker-compose up --build
```

Accessible services:

- **Swagger API Docs**:  
  [http://localhost:8081/swagger/index.html#/](http://localhost:8081/swagger/index.html#/)

- **Kafka UI**:  
  [http://localhost:8082/](http://localhost:8082/)