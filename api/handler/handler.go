package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	//"github.com/jinzhu/gorm"

	"api/model"
	"api/repository"
)

type TaskHandler struct {
	Repo *repository.TaskRepository
}

func NewTaskhandler(repo *repository.TaskRepository) *TaskHandler {
	return &TaskHandler{Repo: repo}
}

func (h *TaskHandler) GetAllCocktails(c *gin.Context) {
	cocktails, err := h.Repo.GetAllCocktails()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "申し訳ございません。そのレシピは存在しません。",
		})
		return
	}
	c.JSON(http.StatusOK, cocktails)
}

func (h *TaskHandler) GetOneCocktail(c *gin.Context) {
	cocktail_name := c.Param("cocktail_name")

	cocktail_details := new(model.CocktailDetails)

	// cocktailsテーブルのレコードを取得
	target_cocktail, err := h.Repo.GetTargetCocktail(cocktail_name)
	if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "申し訳ございません。そのレシピは存在しません。",
			})
	} else {
		cocktail_details.Name = target_cocktail.Name
		cocktail_details.Recipe = target_cocktail.Recipe

		//IngredientDetailsの要素を取得
		ingredient_details, err := h.Repo.GetIngredientDetails(int(target_cocktail.ID))
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "申し訳ございません。そのレシピは存在しません。",
			
			})
		}
		cocktail_details.Ingredients = ingredient_details

		c.JSON(http.StatusOK, cocktail_details)
	}
}