#build
FROM golang:1.21.1 AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN GOOS=linux go build -o main ./cmd/auth

CMD ["./main"]