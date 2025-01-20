package controller

import (
	"app/app/domain"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ItemController struct {
	ItemUseCase domain.ItemUseCase
}

// Get godoc
// @Summary	Get all items
// @Tags Items
// @Produce json
// @Success 200 {object} []domain.Item
// @Failure 500 {object} domain.ErrorResponse
// @Router /item/all [get]
func (ic ItemController) Get(c *gin.Context) {
	if items, err := ic.ItemUseCase.FetchAll(); err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	} else {
		c.JSON(http.StatusOK, items)
	}
}

// Specifications godoc
// @Summary	Get specifications of a posted item
// @Tags Items
// @Accept html/text
// @Produce json
// @Param        data    body   string true  "scheme of order request"
// @Success 200 {object} domain.ItemSpecificationsResponse
// @Failure 400 {object} domain.ErrorResponse
// @Failure 500 {object} domain.ErrorResponse
// @Router /item/specifications [post]
func (ic ItemController) Specifications(c *gin.Context) {
	tmp, err := c.GetRawData()
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}
	name := string(tmp)

	if response, err := ic.ItemUseCase.FetchRequired(name); err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, response)
		return
	}
}
