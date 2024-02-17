#build
FROM golang:1.21.1 AS build

WORKDIR /app

COPY . .

EXPOSE 8080:8080

RUN GOOS=linux go build -o main ./cmd/


#run
FROM gcr.io/distroless/base-debian10

WORKDIR /app

COPY --from=build /app/main .

USER root:root

ENTRYPOINT ["./main"]