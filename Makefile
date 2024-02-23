swag:
	swag init -g cmd/auth/main.go

lint:
	golangci-lint run ./... -c .golangci-lint.yml

gci:
	gci write .