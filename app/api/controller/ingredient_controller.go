package controller

import (
	"app/app/domain"
	"github.com/gin-gonic/gin"
	"net/http"
)

type IngredientController struct {
	IngredientUsecase domain.IngredientUsecase
}

// FindAll godoc
// @Summary	Get a list of all ingredients
// @Tags Ingredients
// @Accept json
// @Produce json
// @Success 200 {object} []domain.Ingredient
// @Failure 204 {object} domain.ErrorResponse
// @Failure 500 {object} domain.ErrorResponse
// @Router /ingredient/get [get]
func (ic *IngredientController) FindAll(c *gin.Context) {
	ingredients, err := ic.IngredientUsecase.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	if len(ingredients) == 0 {
		c.JSON(http.StatusNoContent, domain.ErrorResponse{Message: "No ingredients found"})
		return
	}

	c.JSON(http.StatusOK, ingredients)
}

// Edit godoc
// @Summary	Edit a given ingredient
// @Tags Ingredients
// @Accept json
// @Produce json
// @Param        data    body   domain.Ingredient true  "scheme of ingredient"
// @Success 200 {object} domain.Ingredient
// @Failure 400 {object} domain.ErrorResponse
// @Failure 500 {object} domain.ErrorResponse
// @Router /ingredient/edit [post]
func (ic *IngredientController) Edit(c *gin.Context) {
	var ingredient domain.Ingredient
	err := c.ShouldBind(&ingredient)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	err = ic.IngredientUsecase.Edit(&ingredient)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, ingredient)
}

// Delete godoc
// @Summary	Delete a given ingredient by its marking
// @Tags Ingredients
// @Accept application/text
// @Produce json
// @Param Article body string false "Article of an ingredient to be deleted"
// @Success 200 {object} string
// @Failure 400 {object} domain.ErrorResponse
// @Failure 500 {object} domain.ErrorResponse
// @Router /ingredient/delete [delete]
func (ic *IngredientController) Delete(c *gin.Context) {
	var article string
	err := c.ShouldBind(&article)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	err = ic.IngredientUsecase.Delete(article)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, "Ingredient is deleted successfully")
}
