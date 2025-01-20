package controller

import (
	"app/app/domain"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ToolingController struct {
	ToolingUsecase domain.ToolingUsecase
}

// Create godoc
// @Summary	Create a new tooling
// @Tags Toolings
// @Accept json
// @Produce json
// @Param        data    body   domain.ToolingRequest true  "scheme of login"
// @Success 200 {object} domain.Tooling
// @Failure 400 {object} domain.ErrorResponse
// @Failure 500 {object} domain.ErrorResponse
// @Router /tooling/create [post]
func (tc *ToolingController) Create(c *gin.Context) {
	var request domain.ToolingRequest

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	if request.Name == "" {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Article field can't be empty"})
		return
	}

	if request.Type == nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Type field can't be empty"})
		return
	}

	tooling, err := tc.ToolingUsecase.MapProperties(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	err = tc.ToolingUsecase.Create(tooling)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, tooling)
}

// Get godoc
// @Summary	Get a list of all toolings
// @Tags Toolings
// @Accept json
// @Produce json
// @Success 200 {object} []domain.ToolingResponse
// @Failure 400 {object} domain.ErrorResponse
// @Failure 404 {object} domain.ErrorResponse
// @Failure 500 {object} domain.ErrorResponse
// @Router /tooling/get [get]
func (tc *ToolingController) Get(c *gin.Context) {
	//var request domain.ToolingRequest
	//
	//err := c.ShouldBind(&request)
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
	//}

	//value := reflect.ValueOf(request)
	//queryAll := true
	//for i := 0; i < value.NumField(); i++ {
	//	if !(value.Field(i).IsZero() || value.Field(i).IsNil()) {
	//		queryAll = false
	//	}
	//}

	var toolings []domain.ToolingResponse
	//if queryAll {
	toolings, err := tc.ToolingUsecase.GetAll()
	//} else {
	//	toolings, err = tc.ToolingUsecase.GetByConditions(request)
	//}

	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	if len(toolings) == 0 {
		c.JSON(http.StatusNotFound, domain.ErrorResponse{Message: "No tooling found"})
		return
	}

	c.JSON(http.StatusOK, toolings)
}

// Edit godoc
// @Summary	Update a tooling
// @Tags Toolings
// @Accept json
// @Produce json
// @Param        data    body   domain.ToolingRequest true  "scheme of login"
// @Success 200 {object} domain.Tooling
// @Failure 400 {object} domain.ErrorResponse
// @Failure 500 {object} domain.ErrorResponse
// @Router /tooling/edit [post]
func (tc *ToolingController) Edit(c *gin.Context) {
	var request domain.ToolingRequest

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	tooling, err := tc.ToolingUsecase.MapProperties(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	err = tc.ToolingUsecase.Update(tooling)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, tooling)
}

// Delete godoc
// @Summary	Delete a tooling
// @Tags Toolings
// @Accept application/text
// @Produce json
// @Param Tooling_ID body string false "ID of a tooling to be deleted"
// @Success 200 {object} string
// @Failure 400 {object} domain.ErrorResponse
// @Failure 500 {object} domain.ErrorResponse
// @Router /tooling/delete [delete]
func (tc *ToolingController) Delete(c *gin.Context) {
	var marking string

	err := c.ShouldBind(&marking)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	err = tc.ToolingUsecase.Delete(marking)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, "The tool was deleted successfully")
}
