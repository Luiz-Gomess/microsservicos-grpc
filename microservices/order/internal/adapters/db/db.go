package db

import (
	"fmt"
	"time"

	"github.com/Luiz-Gomess/microservices/order/internal/application/core/domain"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	CustomerID   int64
	Status       string
	OrderItems   []OrderItem
	DeliveryDays int64 
}

type Product struct {
	gorm.Model
	ProductCode string `gorm:"type:varchar(100);uniqueIndex"`
	Name        string
	Quantity    int32
}

type OrderItem struct {
	gorm.Model
	ProductCode string
	UnitPrice   float32
	Quantity    int32
	OrderID     uint
}

type Adapter struct {
	db *gorm.DB
}

func NewAdapter(dataSourceUrl string) (*Adapter, error) {
	db, openErr := gorm.Open(mysql.Open(dataSourceUrl), &gorm.Config{})

	if openErr != nil {
		return nil, fmt.Errorf("db connection error : %v", openErr)
	}
	err := db.AutoMigrate(&Order{}, OrderItem{}, &Product{})

	if err != nil {
		return nil, fmt.Errorf("db migration error : %v", err)
	}

	return &Adapter{db: db}, nil

}

func (a Adapter) CheckStock(productCode string) (bool, error) {
	var count int64
	// Verifica se existe algum produto com este código
	err := a.db.Model(&Product{}).Where("product_code = ?", productCode).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (a Adapter) Get(id string) (domain.Order, error) {
	var orderEntity Order
	res := a.db.First(&orderEntity, id)
	var orderItems []domain.OrderItem

	for _, orderItem := range orderEntity.OrderItems {
		orderItems = append(orderItems, domain.OrderItem{
			ProductCode: orderItem.ProductCode,
			UnitPrice:   orderItem.UnitPrice,
			Quantity:    orderItem.Quantity,
		})
	}

	order := domain.Order{
		ID:           int64(orderEntity.ID),
		CustomerID:   orderEntity.CustomerID,
		Status:       orderEntity.Status,
		OrderItems:   orderItems,
		CreatedAt:    orderEntity.CreatedAt.UnixNano(),
		DeliveryDays: orderEntity.DeliveryDays,
	}
	return order, res.Error
}

func (a Adapter) Save(order *domain.Order) error {
    var orderItems []OrderItem
    for _, orderItem := range order.OrderItems {
        orderItems = append(orderItems, OrderItem{
            ProductCode: orderItem.ProductCode,
            UnitPrice:   orderItem.UnitPrice,
            Quantity:    orderItem.Quantity,
        })
    }
    now := time.Now()

    orderModel := Order{
        Model: gorm.Model{
            ID:        uint(order.ID),
            CreatedAt: now, 
            UpdatedAt: now, 
        },
        CustomerID:   order.CustomerID,
        Status:       order.Status,
        OrderItems:   orderItems,
        DeliveryDays: order.DeliveryDays,
    }

    res := a.db.Save(&orderModel)

    if res.Error == nil {
        order.ID = int64(orderModel.ID)
    }
    return res.Error
}
