package pgquery

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"

	databaseQuery "github.com/wecercle/test-devjr-fullstack-cercle/core/database/postgres/query/sqlc"
	"github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/domain/exception"
	queryrepo "github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/domain/repository/query"
)

type ResaleQueryRepository struct {
	querier *databaseQuery.Queries
}

func NewResaleQueryRepository(querier *databaseQuery.Queries) *ResaleQueryRepository {
	return &ResaleQueryRepository{
		querier: querier,
	}
}

func (r *ResaleQueryRepository) SelectOrderItemsByCPFAndOrderID(
	ctx context.Context,
	cpf string,
	orderID string,
) ([]queryrepo.OrderItem, error) {

	parsedOrderID, err := uuid.Parse(orderID)
	if err != nil {
		return nil, exception.ErrInvalidOrderID
	}

	rows, err := r.querier.SelectOrderItemsByCPFAndOrderID(
		ctx,
		databaseQuery.SelectOrderItemsByCPFAndOrderIDParams{
			OrderID: parsedOrderID,
			DocumentNumber: sql.NullString{
				String: cpf,
				Valid:  true,
			},
		},
	)
	if err != nil {
		return nil, err
	}

	items := make([]queryrepo.OrderItem, 0, len(rows))

	for _, row := range rows {
		item := queryrepo.OrderItem{
			ID:              row.ID.String(),
			FkResaleOrderID: row.FkResaleOrderID.String(),
			Sku:             row.Sku,
			Name:            row.Name,
			Quantity:        row.Quantity,
			AmountValue:     row.AmountValue,
			ShippingStatus:  row.ShippingStatus.String,
			ShippingCode:    row.ShippingCode.String,
		}

		items = append(items, item)
	}

	return items, nil
}

func (r *ResaleQueryRepository) SelectOrderItemForCancel(
	ctx context.Context,
	cpf string,
	orderID string,
	itemID string,
) (*queryrepo.OrderItemForCancel, error) {

	parsedOrderID, err := uuid.Parse(orderID)
	if err != nil {
		return nil, exception.ErrInvalidOrderID
	}

	parsedItemID, err := uuid.Parse(itemID)
	if err != nil {
		return nil, exception.ErrInvalidOrderItemID
	}

	row, err := r.querier.SelectOrderItemForCancel(
		ctx,
		databaseQuery.SelectOrderItemForCancelParams{
			ItemID:  parsedItemID,
			OrderID: parsedOrderID,
			DocumentNumber: sql.NullString{
				String: cpf,
				Valid:  true,
			},
		},
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, exception.ErrOrderItemNotFound
		}
		return nil, err
	}

	var deliveredAt *string
	if row.DeliveredAt.Valid {
		strTime := row.DeliveredAt.Time.Format(time.RFC3339)
		deliveredAt = &strTime
	}

	item := &queryrepo.OrderItemForCancel{
		ID:              row.ID.String(),
		FkResaleOrderID: row.FkResaleOrderID.String(),
		ShippingStatus:  row.ShippingStatus.String,
		DeliveredAt:     deliveredAt,
	}

	return item, nil
}
