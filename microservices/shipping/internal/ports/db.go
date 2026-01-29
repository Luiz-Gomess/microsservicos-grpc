package ports

import "github.com/Luiz-Gomess/microservices/shipping/internal/application/core/domain"

type DBPort interface {
	Save(shipping *domain.Shipping) error
}