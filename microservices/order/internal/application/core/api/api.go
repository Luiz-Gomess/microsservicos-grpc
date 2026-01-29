package api

import (
	"log"

	"github.com/Luiz-Gomess/microservices/order/internal/application/core/domain"
	"github.com/Luiz-Gomess/microservices/order/internal/ports"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Application struct {
	db       ports.DBPort
	payment  ports.PaymentPort
	shipping ports.ShippingPort 
}

func NewApplication(db ports.DBPort, payment ports.PaymentPort, shipping ports.ShippingPort) *Application {
	return &Application{
		db:       db,
		payment:  payment,
		shipping: shipping,
	}
}

func (a Application) PlaceOrder(order domain.Order) (domain.Order, error) {
	log.Printf("[Order] 1. Recebi pedido para o cliente ID: %d", order.CustomerID)

	for _, item := range order.OrderItems {
		log.Printf("[Order] 2. Validando disponibilidade do produto: %s", item.ProductCode)
		exists, err := a.db.CheckStock(item.ProductCode)
		if err != nil {
			log.Printf("[Order] Erro: Produto %s não encontrado!", item.ProductCode)
			return domain.Order{}, status.Errorf(codes.Internal, "error checking stock: %v", err)
		}
		if !exists {
			return domain.Order{}, status.Errorf(codes.NotFound, "item not found in stock: %s", item.ProductCode)
		}
	}

	log.Println("[Order] Todos os produtos são válidos.")

	// Salvar no Banco
	err := a.db.Save(&order)
	if err != nil {
		return domain.Order{}, err
	}
	log.Printf("[Order] 3. Pedido salvo no banco com ID: %d", order.ID)

	// Pagamento
	log.Println("[Order] 4. Chamando serviço de Pagamento...")
	paymentErr := a.payment.Charge(&order)
	if paymentErr != nil {
		log.Printf("[Order] Falha no pagamento: %v", paymentErr)
		return domain.Order{}, paymentErr
	}
	log.Println("[Order] Pagamento aprovado.")

	// Entrega
	log.Println("[Order] 5. Chamando serviço de Entrega (Shipping)...")
	_, err = a.shipping.CreateShipping(order.ID, order.OrderItems)
	if err != nil {
		log.Printf("[Order] Falha ao agendar entrega: %v", err)
		return domain.Order{}, err
	}
	log.Println("[Order] Entrega agendada com sucesso.")

	log.Println("[Order] Processo finalizado.")
	return order, nil
}
