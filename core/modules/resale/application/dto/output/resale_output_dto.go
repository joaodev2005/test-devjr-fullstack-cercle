package dto

type GetOrderItemsResponse struct {
	ID              string `json:"id"`
	FkResaleOrderID string `json:"fk_resale_order_id"`
	Sku             string `json:"sku"`
	Name            string `json:"name"`
	Quantity        int32  `json:"quantity"`
	AmountValue     string `json:"amount_value"`
	ShippingCode    string `json:"shipping_code"`
	ShippingStatus  string `json:"shipping_status"`
}
