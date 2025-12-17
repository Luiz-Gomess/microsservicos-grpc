package ports

import "github.com/Luiz-Gomess/microservices/order/internal/application/core/domain"

type PaymentPort interface {
	Charge(*domain.Order) error
}
