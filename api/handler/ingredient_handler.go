package handler

import (
	"api/repository"
	"github.com/gin-gonic/gin"
	"net/http"
)

type IngredientHandler struct {
	CocktailRepository   *repository.CocktailRepository
	IngredientRepository *repository.IngredientRepository
}

func NewIngredientHandler(cocktailRepository *repository.CocktailRepository, ingredientRepository *repository.IngredientRepository) *IngredientHandler {
	return &IngredientHandler{CocktailRepository: cocktailRepository, IngredientRepository: ingredientRepository}
}

func (h *IngredientHandler) GetIngredients(c *gin.Context) {
	ingredients, err := h.IngredientRepository.GetIngredients()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "申し訳ございません。そのレシピは存在しません。",
		})
		return
	}
	c.JSON(http.StatusOK, ingredients)
}
