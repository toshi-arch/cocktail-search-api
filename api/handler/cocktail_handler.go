package handler

import (
	modelDetail "api/model_detail"
	"api/repository"
	"github.com/gin-gonic/gin"
	"net/http"
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
	//IngredientDetailsの要素を取得
	ingredient_detail, err := h.IngredientRepository.GetIngredientsByCocktailId(int(cocktail.ID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "申し訳ございません。そのレシピは存在しません。",
		})
		return
	}

	c.JSON(http.StatusOK, modelDetail.Cocktail{
		Name:        cocktail.Name,
		Recipe:      cocktail.Recipe,
		Ingredients: *ingredient_detail,
	})
}

func (h *CocktailHandler) GetCocktailNamesByIngredient(c *gin.Context) {
	ingredientName := c.Param("ingredient_name")

	// ingredientsテーブルのレコードを取得
	ingredient, err := h.IngredientRepository.GetIngredientByName(ingredientName)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "申し訳ございません。その材料は存在しません。",
		})
		return
	}
	//Cocktailsの要素を取得
	cocktails, err := h.CocktailRepository.GetCocktailByIngredient(int(ingredient.ID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "申し訳ございません。その材料は存在しません。",
		})
		return
	}

	c.JSON(http.StatusOK, modelDetail.CocktailByIngredient{
		IngredientName: ingredient.Name,
		CocktailName:   cocktails,
	})
}
