package handler

import (
	"api/repository"
	"github.com/gin-gonic/gin"
	"net/http"
)

type IngredientHandler struct {
	Repo_cocktail   *repository.CocktailRepository
	Repo_ingredient *repository.IngredientRepository
}

func NewIngredientHandler(repo_cocktail *repository.CocktailRepository, repo_ingredient *repository.IngredientRepository) *IngredientHandler {
	return &IngredientHandler{Repo_cocktail: repo_cocktail, Repo_ingredient: repo_ingredient}
}

func (h *IngredientHandler) GetIngredients(c *gin.Context) {
	ingredients, err := h.Repo_ingredient.GetIngredients()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "申し訳ございません。そのレシピは存在しません。",
		})
		return
	}
	c.JSON(http.StatusOK, ingredients)
}
