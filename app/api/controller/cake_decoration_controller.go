package controller

import (
	"app/app/domain"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CakeDecorationController struct {
	CakeDecorationUsecase domain.CakeDecorationUsecase
}

// FindAll godoc
// @Summary	Get a list of all cake decorations
// @Tags Cake Decorations
// @Accept json
// @Produce json
// @Success 200 {object} []domain.CakeDecoration
// @Failure 204 {object} domain.ErrorResponse
// @Failure 500 {object} domain.ErrorResponse
// @Router /decoration/get [get]
func (cd *CakeDecorationController) FindAll(c *gin.Context) {
	ingredients, err := cd.CakeDecorationUsecase.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	if len(ingredients) == 0 {
		c.JSON(http.StatusNoContent, domain.ErrorResponse{Message: "No cake decorations found"})
		return
	}

	c.JSON(http.StatusOK, ingredients)
}

// Edit godoc
// @Summary	Edit a given decoration
// @Tags Cake Decorations
// @Accept json
// @Produce json
// @Param        data    body   domain.CakeDecoration true  "scheme of cake decoration"
// @Success 200 {object} domain.CakeDecoration
// @Failure 400 {object} domain.ErrorResponse
// @Failure 500 {object} domain.ErrorResponse
// @Router /decoration/edit [post]
func (cd *CakeDecorationController) Edit(c *gin.Context) {
	var cakeDecoration domain.CakeDecoration
	err := c.ShouldBind(&cakeDecoration)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	err = cd.CakeDecorationUsecase.Edit(&cakeDecoration)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, cakeDecoration)
}

// Delete godoc
// @Summary	Delete a given cake decoration by its article
// @Tags Cake Decorations
// @Accept application/text
// @Produce json
// @Param Article body string false "Article of a cake decoration to be deleted"
// @Success 200 {object} string
// @Failure 400 {object} domain.ErrorResponse
// @Failure 500 {object} domain.ErrorResponse
// @Router /decoration/delete [delete]
func (cd *CakeDecorationController) Delete(c *gin.Context) {
	var article string
	err := c.ShouldBind(&article)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	err = cd.CakeDecorationUsecase.Delete(article)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, "Cake decoration is deleted successfully")
}
