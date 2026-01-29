FROM golang:1.24-alpine as builder

WORKDIR /app

COPY microservices-proto /app/microservices-proto

COPY microservices/order /app/microservices/order

WORKDIR /app/microservices/order

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o order-service cmd/main.go

FROM alpine:3.19
WORKDIR /app
COPY --from=builder /app/microservices/order/order-service .

EXPOSE 8080
CMD ["./order-service"]