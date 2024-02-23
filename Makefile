swag:
	swag init -g cmd/auth/main.go

lint:
	golangci-lint run ./...