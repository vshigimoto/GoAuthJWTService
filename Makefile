swag:
	swag init -g cmd/auth/main.go

lint:
	golangci-lint run ./... -c .golangci-lint.yml

gci:
	gci write .

local:
	docker-compose up -d mongo
	go run cmd/auth/main.go

run:
	docker-compose up -d --build

test:
	go test -v ./...