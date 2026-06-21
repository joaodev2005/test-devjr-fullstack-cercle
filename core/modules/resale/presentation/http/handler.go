package http

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/application/usecase"
	"github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/domain/exception"
	httpresponse "github.com/wecercle/test-devjr-fullstack-cercle/core/modules/shared/presentation/http/response"
)

type Handler struct {
	useCase *usecase.ResaleUseCase
}

func NewHandler(u *usecase.ResaleUseCase) *Handler {
	return &Handler{
		useCase: u,
	}
}

// GetOrderItemsByCPFAndOrderID godoc
// @Summary List order items by CPF and order ID
// @Description Returns all items of a specific order identified by user CPF and order ID
// @Tags Resale
// @Accept json
// @Produce json
// @Param cpf path string true "User CPF (digits only, 11 characters)"
// @Param order_id path string true "Order ID (UUID)"
// @Success 200 {object} map[string]interface{} "Order items retrieved successfully"
// @Failure 400 {object} map[string]string "Bad request"
// @Failure 404 {object} map[string]string "Order not found"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /v1/app/users/{cpf}/orders/{order_id}/items [get]
func (h *Handler) GetOrderItemsByCPFAndOrderID(c *gin.Context) {
	cpf := c.Param("cpf")
	orderID := c.Param("order_id")

	if _, err := uuid.Parse(orderID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid order_id format"})
		return
	}

	items, err := h.useCase.ListItems(c.Request.Context(), cpf, orderID)
	if err != nil {
		if errors.Is(err, exception.ErrOrderNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "order not found or does not belong to user",
			})
			return
		}

		if errors.Is(err, exception.ErrInvalidOrderID) {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "invalid order_id format",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	httpresponse.Success(c, http.StatusOK, items)
}

// CancelOrderItem godoc
// @Summary Cancel an order item
// @Description Cancels a specific item of an order identified by user CPF, order ID and item ID
// @Tags Resale
// @Accept json
// @Produce json
// @Param cpf path string true "User CPF (digits only, 11 characters)"
// @Param order_id path string true "Order ID (UUID)"
// @Param item_id path string true "Order Item ID (UUID)"
// @Success 204 "Item cancelled successfully or already returned"
// @Failure 400 {object} map[string]string "Bad request"
// @Failure 404 {object} map[string]string "Order or item not found"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /v1/app/users/{cpf}/orders/{order_id}/items/{item_id}/cancel [put]
func (h *Handler) CancelOrderItem(c *gin.Context) {
	cpf := c.Param("cpf")
	orderID := c.Param("order_id")
	itemID := c.Param("item_id")

	if _, err := uuid.Parse(orderID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid order_id format"})
		return
	}
	if _, err := uuid.Parse(itemID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid item_id format"})
		return
	}

	_, err := h.useCase.CancelItem(c.Request.Context(), cpf, orderID, itemID)
	if err != nil {
		if errors.Is(err, exception.ErrOrderItemNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "item or order not found"})
			return
		}
		if errors.Is(err, exception.ErrReturnPeriodExpired) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "return deadline exceeded"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
