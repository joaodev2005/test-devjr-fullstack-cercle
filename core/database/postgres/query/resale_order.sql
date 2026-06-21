-- name: SelectOrderItemsByCPFAndOrderID :many
SELECT
    roi.id,
    roi.fk_resale_order_id,
    roi.sku,
    roi.name,
    roi.quantity,
    roi.amount_value,
    roi.shipping_code,
    roi.shipping_status
FROM cercle_test.resale_order_item roi
JOIN cercle_test.resale_order ro ON ro.id = roi.fk_resale_order_id
JOIN cercle_test.users u ON u.id = ro.fk_users_id
WHERE ro.id = @order_id
  AND u.document_number = @document_number
  AND roi.deleted_at IS NULL
ORDER BY roi.created_at ASC;

-- name: UpdateOrderItemShippingStatus :execrows
UPDATE cercle_test.resale_order_item
SET shipping_status = @shipping_status,
    updated_at = NOW()
WHERE fk_resale_order_id = @resale_order_id
  AND id = @id;

-- name: SelectOrderItemForCancel :one
SELECT
    roi.id,
    roi.fk_resale_order_id,
    roi.shipping_status,
    roi.delivered_at
FROM cercle_test.resale_order_item roi
JOIN cercle_test.resale_order ro
    ON ro.id = roi.fk_resale_order_id
JOIN cercle_test.users u
    ON u.id = ro.fk_users_id
WHERE roi.id = @item_id
  AND ro.id = @order_id
  AND u.document_number = @document_number
  AND roi.deleted_at IS NULL;