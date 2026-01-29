package db

import (
	"fmt"

	"github.com/Luiz-Gomess/microservices/shipping/internal/application/core/domain"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ShippingEntity struct {
	gorm.Model
	OrderID       int64
	Status        string
	DeliveryDays  int64
	ShippingItems []ShippingItemEntity `gorm:"foreignKey:ShippingID"`
}

type ShippingItemEntity struct {
	gorm.Model
	ProductCode string
	Quantity    int32
	ShippingID  uint
}

type Adapter struct {
	db *gorm.DB
}

func NewAdapter(dataSourceUrl string) (*Adapter, error) {
	db, openErr := gorm.Open(mysql.Open(dataSourceUrl), &gorm.Config{})
	if openErr != nil {
		return nil, fmt.Errorf("db connection error: %v", openErr)
	}

	err := db.AutoMigrate(&ShippingEntity{}, &ShippingItemEntity{})
	if err != nil {
		return nil, fmt.Errorf("db migration error: %v", err)
	}

	return &Adapter{db: db}, nil
}

func (a Adapter) Save(shipping *domain.Shipping) error {
	var shippingItems []ShippingItemEntity
	
	for _, item := range shipping.ShippingItems {
		shippingItems = append(shippingItems, ShippingItemEntity{
			ProductCode: item.ProductCode,
			Quantity:    item.Quantity,
		})
	}

	shippingModel := ShippingEntity{
		OrderID:       shipping.OrderID,
		Status:        shipping.Status,
		DeliveryDays:  shipping.DeliveryDays,
		ShippingItems: shippingItems,
	}

	res := a.db.Create(&shippingModel)
	
	if res.Error == nil {
		shipping.ID = int64(shippingModel.ID)
	}

	return res.Error
}