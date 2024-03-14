package handler

import (
	model_detail "api/model_detail"
	//model_database "api/model_database"
	"api/repository"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CocktailHandler struct {
	Repo_cocktail   *repository.CocktailRepository
	Repo_ingredient *repository.IngredientRepository
}

func NewCocktailHandler(repo_cocktail *repository.CocktailRepository, repo_ingredient *repository.IngredientRepository) *CocktailHandler {
	return &CocktailHandler{Repo_cocktail: repo_cocktail, Repo_ingredient: repo_ingredient}
}

func (h *CocktailHandler) GetAllCocktails(c *gin.Context) {
	cocktails, err := h.Repo_cocktail.GetAllCocktails()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "申し訳ございません。そのレシピは存在しません。",
		})
		return
	}
	c.JSON(http.StatusOK, cocktails)
}

func (h *CocktailHandler) GetCocktailByName(c *gin.Context) {
	cocktail_name := c.Param("cocktail_name")

	// cocktailsテーブルのレコードを取得
	cocktail, err := h.Repo_cocktail.GetCocktailByName(cocktail_name)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "申し訳ございません。そのレシピは存在しません。",
		})
		return
	} else {
		//IngredientDetailsの要素を取得
		ingredient_detail, err := h.Repo_ingredient.GetIngredientDetail(int(cocktail.ID))
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "申し訳ございません。そのレシピは存在しません。",
			})
		}

		c.JSON(http.StatusOK, model_detail.Cocktail{
			Name:        cocktail.Name,
			Recipe:      cocktail.Recipe,
			Ingredients: ingredient_detail,
		})
	}
}

func (h *CocktailHandler) GetCocktailByIngredient(c *gin.Context) {
	ingredient_name := c.Param("ingredient_name")

	// ingredientsテーブルのレコードを取得
	ingredient, err := h.Repo_ingredient.GetIngredientByName(ingredient_name)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "申し訳ございません。その材料は存在しません。",
		})
		return
	} else {
		//Cocktailsの要素を取得
		cocktails, err := h.Repo_cocktail.GetCocktailByIngredient(int(ingredient.ID))
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "申し訳ございません。その材料は存在しません。",
			})
		}

		c.JSON(http.StatusOK, model_detail.CocktailByIngredient{
			IngredientName: ingredient.Name,
			CocktailName: cocktails,
		})
		//c.JSON(http.StatusOK, cocktails)
	}
}
