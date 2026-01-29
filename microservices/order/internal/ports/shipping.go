package ports

import "github.com/Luiz-Gomess/microservices/order/internal/application/core/domain"

type ShippingPort interface {
	CreateShipping(orderID int64, items []domain.OrderItem) (int64, error)
}