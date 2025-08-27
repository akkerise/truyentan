SWAGGER_DIR=backend

.PHONY: swagger
swagger:
	cd $(SWAGGER_DIR) && swag init -g cmd/api/main.go -o docs --parseDependency --parseInternal
