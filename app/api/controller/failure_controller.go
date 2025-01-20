package controller

import (
	"app/app/domain"
	"github.com/gin-gonic/gin"
	"net/http"
)

type FailureController struct {
	FailureRepository domain.FailureRepository
}

func NewFailureController(failureRepository domain.FailureRepository) *FailureController {
	return &FailureController{FailureRepository: failureRepository}
}

// RegisterFailure godoc
// @Summary	Register a new failure manually
// @Tags Failures
// @Accepts json
// @Produce json
// @Success 200 {object} domain.Failure
// @Failure 400 {object} domain.ErrorResponse
// @Failure 500 {object} domain.ErrorResponse
// @Router /failure/register [post]
func (fc *FailureController) RegisterFailure(c *gin.Context) {
	var failure domain.Failure
	if err := c.ShouldBind(&failure); err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	if err := fc.FailureRepository.Create(&failure); err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, failure)
}

// ListFailures godoc
// @Summary	List all failures occurred in the system
// @Tags Failures
// @Produce json
// @Success 200 {object} []domain.Failure
// @Success 204 {object} domain.SuccessResponse
// @Failure 400 {object} domain.ErrorResponse
// @Failure 500 {object} domain.ErrorResponse
// @Router /failure/list [get]
func (fc *FailureController) ListFailures(c *gin.Context) {
	if failures, err := fc.FailureRepository.FetchAll(); err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	} else if len(failures) == 0 {
		c.JSON(http.StatusNoContent, domain.SuccessResponse{Message: "No failures collected"})
	} else {
		c.JSON(http.StatusOK, failures)
	}
}
