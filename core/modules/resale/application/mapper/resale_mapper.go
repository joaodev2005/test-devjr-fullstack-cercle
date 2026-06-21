package mapper

import (
	output "github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/application/dto/output"
	queryrepo "github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/domain/repository/query"
)

func ToItemResponseList(domainItems []queryrepo.OrderItem) []output.GetOrderItemsResponse {
	dtos := make([]output.GetOrderItemsResponse, 0, len(domainItems))

	for _, item := range domainItems {
		dtos = append(dtos, output.GetOrderItemsResponse{
			ID:              item.ID,
			FkResaleOrderID: item.FkResaleOrderID,
			Sku:             item.Sku,
			Name:            item.Name,
			Quantity:        item.Quantity,
			AmountValue:     item.AmountValue,
			ShippingCode:    item.ShippingCode,
			ShippingStatus:  item.ShippingStatus,
		})
	}

	return dtos
}
