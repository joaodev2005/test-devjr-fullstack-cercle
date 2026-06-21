package exception

var (
	ErrOrderNotFound       = &DomainError{Code: "order_not_found", Message: "order not found"}
	ErrOrderItemNotFound   = &DomainError{Code: "order_item_not_found", Message: "order item not found"}
	ErrInvalidOrderID      = &DomainError{Code: "invalid_order_id", Message: "invalid order id"}
	ErrInvalidOrderItemID  = &DomainError{Code: "invalid_order_item_id", Message: "invalid order item id"}
	ErrReturnPeriodExpired = &DomainError{Code: "return_period_expired", Message: "return period expired"}
)
