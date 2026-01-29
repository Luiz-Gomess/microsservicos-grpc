FROM golang:1.24

WORKDIR /app

COPY microservices-proto /app/microservices-proto
COPY microservices/shipping /app/microservices/shipping

WORKDIR /app/microservices/shipping

RUN go mod download

RUN go version

EXPOSE 8082

CMD ["go", "run", "cmd/main.go"]