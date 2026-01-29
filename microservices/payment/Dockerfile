FROM golang:1.24-alpine as builder

WORKDIR /app

COPY microservices-proto /app/microservices-proto

COPY microservices/payment /app/microservices/payment

WORKDIR /app/microservices/payment

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o payment-service cmd/main.go

FROM alpine:3.19
WORKDIR /app
COPY --from=builder /app/microservices/payment/payment-service .

EXPOSE 8081
CMD ["./payment-service"]