FROM golang:1.22-alpine as build

WORKDIR /app

COPY . .
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o hexagonal-architecture-go cmd/main.go

EXPOSE 8080

CMD ["./hexagonal-architecture-go"]
