package shipping

import (
	"context"
	"fmt"

	"github.com/Luiz-Gomess/microservices-proto/golang/shipping"
	"github.com/Luiz-Gomess/microservices/order/internal/application/core/domain"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Adapter struct {
	client shipping.ShippingClient
}

func NewAdapter(shippingServiceUrl string) (*Adapter, error) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.NewClient(shippingServiceUrl, opts...)
	if err != nil {
		return nil, err
	}

	client := shipping.NewShippingClient(conn)
	return &Adapter{client: client}, nil
}

func (a *Adapter) CreateShipping(orderID int64, items []domain.OrderItem) (int64, error) {
	var grpcItems []*shipping.ShippingItem
	for _, item := range items {
		grpcItems = append(grpcItems, &shipping.ShippingItem{
			ProductCode: item.ProductCode,
			Quantity:    item.Quantity,
		})
	}

	req := &shipping.CreateShippingRequest{
		OrderId: orderID,
		Items:   grpcItems,
	}

	resp, err := a.client.Create(context.Background(), req)
	if err != nil {
		return 0, fmt.Errorf("error calling shipping service: %v", err)
	}

	return resp.DeliveryDays, nil
}
