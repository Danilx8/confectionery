package controller

import (
	"app/app/domain"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ItemController struct {
	domain.ItemRepository
}

// Get godoc
// @Summary	Get all items
// @Tags Orders
// @Produce json
// @Success 200 {object} []domain.Item
// @Failure 500 {object} domain.ErrorResponse
// @Router /item/all [get]
func (ic ItemController) Get(c *gin.Context) {
	if items, err := ic.Fetch(); err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	} else {
		c.JSON(http.StatusOK, items)
	}
}
