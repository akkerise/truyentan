SWAGGER_DIR=backend
BACKEND_DIR=backend
WEB_DIR=web

.PHONY: dev build swagger lint seed

dev:
	wails dev

build:
	wails build

swagger:
	cd $(SWAGGER_DIR) && swag init -g cmd/api/main.go -o docs --parseDependency --parseInternal

lint:
	cd $(BACKEND_DIR) && golangci-lint run ./...
	cd $(WEB_DIR) && npm run lint

seed:
	go run $(BACKEND_DIR)/cmd/seed/main.go
