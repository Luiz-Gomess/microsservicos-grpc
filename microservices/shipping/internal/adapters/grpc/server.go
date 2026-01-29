package grpc

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/Luiz-Gomess/microservices-proto/golang/shipping"
	"github.com/Luiz-Gomess/microservices/shipping/internal/application/core/domain"
	"github.com/Luiz-Gomess/microservices/shipping/internal/ports"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Adapter struct {
	api  ports.APIPort
	port int
	shipping.UnimplementedShippingServer
}

func NewAdapter(api ports.APIPort, port int) *Adapter {
	return &Adapter{api: api, port: port}
}

func (a Adapter) Create(ctx context.Context, request *shipping.CreateShippingRequest) (*shipping.CreateShippingResponse, error) {
	var shippingItems []domain.ShippingItem
	for _, item := range request.Items {
		shippingItems = append(shippingItems, domain.ShippingItem{
			ProductCode: item.ProductCode,
			Quantity:    item.Quantity,
		})
	}

	domainShipping := domain.NewShipping(request.OrderId, shippingItems)

	result, err := a.api.Create(domainShipping)
	if err != nil {
		return nil, err
	}

	return &shipping.CreateShippingResponse{
		ShippingId:   result.ID,
		DeliveryDays: result.DeliveryDays,
	}, nil
}

func (a Adapter) Run() {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	shipping.RegisterShippingServer(grpcServer, a)

	reflection.Register(grpcServer)

	log.Printf("starting shipping service on port %d ...", a.port)
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
