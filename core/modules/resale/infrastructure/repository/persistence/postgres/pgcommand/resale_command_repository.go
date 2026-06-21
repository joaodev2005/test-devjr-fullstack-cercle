package pgcommand

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	databaseQuery "github.com/wecercle/test-devjr-fullstack-cercle/core/database/postgres/query/sqlc"
)

type ResaleCommandRepository struct {
	querier *databaseQuery.Queries
}

func NewResaleCommandRepository(querier *databaseQuery.Queries) *ResaleCommandRepository {
	return &ResaleCommandRepository{
		querier: querier,
	}
}

func (r *ResaleCommandRepository) UpdateOrderItemShippingStatus(
	ctx context.Context,
	orderID string,
	itemID string,
	shippingStatus string,
) error {
	orderUUID, err := uuid.Parse(orderID)
	if err != nil {
		return err
	}

	itemUUID, err := uuid.Parse(itemID)
	if err != nil {
		return err
	}

	_, err = r.querier.UpdateOrderItemShippingStatus(
		ctx,
		databaseQuery.UpdateOrderItemShippingStatusParams{
			ShippingStatus: sql.NullString{
				String: shippingStatus,
				Valid:  true,
			},
			ResaleOrderID: orderUUID,
			ID:            itemUUID,
		},
	)

	return err
}
