package handler

import (
	"api/domain"
	"api/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CocktailHandler struct {
	CocktailRepository   *repository.CocktailRepository
	IngredientRepository *repository.IngredientRepository
}

func NewCocktailHandler(cocktailRepository *repository.CocktailRepository, ingredientRepository *repository.IngredientRepository) *CocktailHandler {
	return &CocktailHandler{CocktailRepository: cocktailRepository, IngredientRepository: ingredientRepository}
}

func (h *CocktailHandler) GetCocktails(c *gin.Context) {
	cocktails, err := h.CocktailRepository.GetCocktails()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "申し訳ございません。そのレシピは存在しません。",
		})
		return
	}
	c.JSON(http.StatusOK, cocktails)
}

func (h *CocktailHandler) GetCocktailByName(c *gin.Context) {
	cocktailName := c.Param("cocktail_name")

	// cocktailsテーブルのレコードを取得
	cocktail, err := h.CocktailRepository.GetCocktailByName(cocktailName)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "申し訳ございません。そのレシピは存在しません。",
		})
		return
	}
	//Ingredientの要素を取得
	ingredients, err := h.IngredientRepository.GetIngredientsByCocktailId(int(cocktail.ID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "申し訳ございません。そのレシピは存在しません。",
		})
		return
	}

	c.JSON(http.StatusOK, domain.Cocktail{
		Name:        cocktail.Name,
		Recipe:      cocktail.Recipe,
		Ingredients: *ingredients,
	})
}

func (h *CocktailHandler) GetCocktailNamesByIngredient(c *gin.Context) {
	ingredientName := c.Param("ingredient_name")

	//ingredientsテーブルのレコードを取得
	ingredient, err := h.IngredientRepository.GetIngredientByName(ingredientName)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "申し訳ございません。その材料は存在しません。",
		})
		return
	}
	//Cocktailsの要素を取得
	cocktails, _ := h.CocktailRepository.GetCocktailByIngredient(int(ingredient.ID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "申し訳ございません。"+ ingredientName + "を材料に持つカクテルは存在しません。",
		})
		return
	}

	c.JSON(http.StatusOK, domain.CocktailByIngredient{
		IngredientName: ingredient.Name,
		CocktailName:   cocktails,
	})
}
