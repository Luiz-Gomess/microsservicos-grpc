package ports

import "github.com/Luiz-Gomess/microservices/order/internal/application/core/domain"

type DBPort interface {
	Get(id string) (domain.Order, error)
	Save(*domain.Order) error
	CheckStock(productCode string) (bool, error) 
}
