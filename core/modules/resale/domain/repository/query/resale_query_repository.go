package query

import "context"

type OrderItem struct {
	ID              string
	FkResaleOrderID string
	Sku             string
	Name            string
	Quantity        int32
	AmountValue     string
	ShippingCode    string
	ShippingStatus  string
}

type OrderItemForCancel struct {
	ID              string
	FkResaleOrderID string
	ShippingStatus  string
	DeliveredAt     *string
}

type ResaleQueryRepository interface {
	SelectOrderItemsByCPFAndOrderID(ctx context.Context, cpf string, orderID string) ([]OrderItem, error)
	SelectOrderItemForCancel(ctx context.Context, cpf string, orderID string, itemID string) (*OrderItemForCancel, error)
}
