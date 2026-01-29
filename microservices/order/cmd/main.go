package main

import (
	"log"

	"github.com/Luiz-Gomess/microservices/order/config"
	"github.com/Luiz-Gomess/microservices/order/internal/adapters/db"
	"github.com/Luiz-Gomess/microservices/order/internal/adapters/grpc"
	payment_adapter "github.com/Luiz-Gomess/microservices/order/internal/adapters/payment"
	"github.com/Luiz-Gomess/microservices/order/internal/adapters/shipping" // Import novo
	"github.com/Luiz-Gomess/microservices/order/internal/application/core/api"
)

func main() {
	dbAdapter, err := db.NewAdapter(config.GetDataSourceURL())
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	paymentAdapter, err := payment_adapter.NewAdapter(config.GetPaymentServiceURL())
	if err != nil {
		log.Fatalf("Failed to connect to payment service: %v", err)
	}

	shippingAdapter, err := shipping.NewAdapter(config.GetShippingServiceUrl())
	if err != nil {
		log.Fatalf("Failed to connect to shipping service: %v", err)
	}

	application := api.NewApplication(dbAdapter, paymentAdapter, shippingAdapter)

	grpcAdapter := grpc.NewAdapter(application, config.GetApplicationPort())
	grpcAdapter.Run()
}
