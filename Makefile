ensure-reflex:
	@if ! command -v reflex >/dev/null 2>&1; then \
		echo "🔍 Reflex not found. Installing..."; \
		go install github.com/cespare/reflex@latest; \
		echo "✅ Reflex installed."; \
	fi

ensure-goose:
	@if ! command -v goose >/dev/null 2>&1; then \
		echo "🔍 Goose not found. Installing..."; \
		go install github.com/pressly/goose/v3/cmd/goose@latest; \
		echo "✅ Goose installed."; \
	fi

ensure-swagger:
	@if ! command -v swag >/dev/null 2>&1; then \
		echo "🔍 Swag not found. Installing..."; \
		go install github.com/swaggo/swag/cmd/swag@latest; \
		echo "✅ Swag installed."; \
	fi

ensure-mockgen:
	@if ! command -v mockgen >/dev/null 2>&1; then \
		echo "🔍 Mockgen not found. Installing..."; \
		go install github.com/golang/mock/mockgen@v1.6.0; \
		echo "✅ Mockgen installed."; \
	fi

# mockgen dependency
ensure-gomock: 
	@go get -d github.com/golang/mock/gomock@v1.6.0 >/dev/null 2>&1 || true

tidy:
	go mod tidy

mock: ensure-mockgen ensure-gomock
	go generate ./...

unit-test:
	go test -cover $$(go list ./... | grep -v -E '/(mocks|vendor|docs)')

test: mock unit-test

api-docs: ensure-swagger
	swag init -g server/http_router.go --parseDependency true --parseInternal true

# go install github.com/pressly/goose/v3/cmd/goose@latest
# create migration file
# goose -dir ./migrations create add_some_column sql

migrate-up: ensure-goose
	goose -dir ./migrations -table migration_version postgres "postgres://mkpuser:mkppassword@mkp-postgresql:5432/eticket_db?sslmode=disable" up

migrate-down: ensure-goose
	goose -dir ./migrations -table migration_version postgres "postgres://mkpuser:mkppassword@mkp-postgresql:5432/eticket_db?sslmode=disable" down

pre-run: tidy mock unit-test

run-http: ensure-reflex
	reflex -r '\.go' -s -- sh -c "go run main.go serve-http"

run-all: pre-run migrate-up run-http
	reflex -r '\.go' -s -- sh -c "go run main.go serve-http"