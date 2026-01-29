package api

import (
	"log"

	"github.com/Luiz-Gomess/microservices/shipping/internal/application/core/domain"
	"github.com/Luiz-Gomess/microservices/shipping/internal/ports"
)

type Application struct {
	db ports.DBPort
}

func NewApplication(db ports.DBPort) *Application {
	return &Application{
		db: db,
	}
}

func (a Application) Create(shipping domain.Shipping) (domain.Shipping, error) {

	log.Printf("[Shipping] Solicitação de entrega para o Pedido ID: %d. Calculando prazo...", shipping.OrderID)

	shipping.CalculateDeliveryDays()

	log.Printf("[Shipping] Prazo calculado: %d dias. Salvando...", shipping.DeliveryDays)

	err := a.db.Save(&shipping)
	if err != nil {
		return domain.Shipping{}, err
	}

	return shipping, nil
}
