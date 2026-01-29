package domain

import "time"

type ShippingItem struct {
	ProductCode string
	Quantity    int32
}

type Shipping struct {
	ID            int64
	OrderID       int64
	Status        string
	ShippingItems []ShippingItem
	DeliveryDays  int64
	CreatedAt     int64
}

func NewShipping(orderID int64, shippingItems []ShippingItem) Shipping {
	return Shipping{
		OrderID:       orderID,
		ShippingItems: shippingItems,
		Status:        "Pending",
		CreatedAt:     time.Now().Unix(),
	}
}

func (s *Shipping) CalculateDeliveryDays() {
	var totalQuantity int32
	for _, item := range s.ShippingItems {
		totalQuantity += item.Quantity
	}

	baseDays := int64(1)
	extraDays := int64(totalQuantity / 5)

	s.DeliveryDays = baseDays + extraDays
}