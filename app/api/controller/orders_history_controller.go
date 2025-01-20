package controller

import (
	"app/app/domain"
	"github.com/gin-gonic/gin"
	"net/http"
)

type OrdersHistoryController struct {
	ordersHistoryRepository domain.OrdersHistoryRepository
}

func NewOrdersHistoryController(repository domain.OrdersHistoryRepository) *OrdersHistoryController {
	return &OrdersHistoryController{
		ordersHistoryRepository: repository,
	}
}

// Get godoc
// @Summary	Get all orders history
// @Tags Orders
// @Produce json
// @Success 200 {object} []domain.OrdersHistory
// @Failure 500 {object} domain.ErrorResponse
// @Router /order/history [get]
func (oh *OrdersHistoryController) Get(c *gin.Context) {
	if history, err := oh.ordersHistoryRepository.FetchAll(); err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, history)
		return
	}
}
