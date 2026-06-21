package command

import "context"

type ResaleCommandRepository interface {
	UpdateOrderItemShippingStatus(ctx context.Context, orderID string, itemID string, shippingStatus string) error
}
