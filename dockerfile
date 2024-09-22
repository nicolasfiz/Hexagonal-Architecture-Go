FROM golang:1.22-alpine as build

WORKDIR /app

COPY . .
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o hexagonal-architecture-go cmd/main.go

FROM gcr.io/distroless/base-debian11

WORKDIR /

COPY --from=build /app/hexagonal-architecture-go ./hexagonal-architecture-go

EXPOSE 8080

CMD ["./hexagonal-architecture-go"]
