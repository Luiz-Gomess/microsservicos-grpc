package ports

import "github.com/Luiz-Gomess/microservices/shipping/internal/application/core/domain"

type APIPort interface {
	Create(shipping domain.Shipping) (domain.Shipping, error)
}