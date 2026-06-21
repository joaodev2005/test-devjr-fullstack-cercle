package usecase

import (
	"context"
	"time"

	output "github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/application/dto/output"
	"github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/application/mapper"
	"github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/domain/exception"
	commandrepo "github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/domain/repository/command"
	queryrepo "github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/domain/repository/query"
)

type ResaleUseCase struct {
	queryRepo   queryrepo.ResaleQueryRepository
	commandRepo commandrepo.ResaleCommandRepository
}

func NewResaleUseCase(q queryrepo.ResaleQueryRepository, c commandrepo.ResaleCommandRepository) *ResaleUseCase {
	return &ResaleUseCase{queryRepo: q, commandRepo: c}
}

func (u *ResaleUseCase) ListItems(ctx context.Context, cpf string, orderID string) ([]output.GetOrderItemsResponse, error) {
	items, err := u.queryRepo.SelectOrderItemsByCPFAndOrderID(ctx, cpf, orderID)
	if err != nil {
		return nil, err
	}

	if len(items) == 0 {
		return nil, exception.ErrOrderNotFound
	}

	return mapper.ToItemResponseList(items), nil
}

func (u *ResaleUseCase) CancelItem(ctx context.Context, cpf string, orderID string, itemID string) (bool, error) {
	item, err := u.queryRepo.SelectOrderItemForCancel(ctx, cpf, orderID, itemID)
	if err != nil {
		return false, err
	}

	if item.ShippingStatus == "RETURNED" {
		return true, nil
	}

	if item.ShippingStatus == "DELIVERED" && item.DeliveredAt != nil {
		deliveryTime, err := time.Parse(time.RFC3339, *item.DeliveredAt)
		if err == nil {
			sevenDaysAgo := time.Now().AddDate(0, 0, -7)
			if deliveryTime.Before(sevenDaysAgo) {
				return false, exception.ErrReturnPeriodExpired
			}
		}
	}

	err = u.commandRepo.UpdateOrderItemShippingStatus(ctx, orderID, itemID, "RETURNED")
	if err != nil {
		return false, err
	}

	return false, nil
}
