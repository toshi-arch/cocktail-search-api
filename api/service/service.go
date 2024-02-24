package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"api/infra"
	"api/model"
)

func GetAll(c *gin.Context, db *gorm.DB) {
	//db := infra.GormConnect()
	cocktails := []model.Cocktails{}
	infra.GetAll(db, &cocktails)

	c.JSON(http.StatusOK, cocktails)
}

func GetOneCocktail(c *gin.Context, db *gorm.DB) {
	//db := infra.GormConnect()
	cocktail_name := c.Param("cocktail_name")

	cocktail_details := new(model.CocktailDetails)
	ingredient_details := []model.IngredientDetails{}

	// cocktailsテーブルのレコードを取得
	target_cocktail := model.Cocktails{}

	if err := infra.GetTargetCocktail(db, cocktail_name, &target_cocktail); err != nil {
		{
			c.JSON(http.StatusNotFound, gin.H{
				"message": "申し訳ございません。そのレシピは存在しません。",
			})
		}
	} else {
		cocktail_details.Name = target_cocktail.Name
		cocktail_details.Recipe = target_cocktail.Recipe

		//IngredientDetailsの要素を取得
		infra.GetIngredientDetails(db, int(target_cocktail.ID), &ingredient_details)

		cocktail_details.Ingredients = ingredient_details

		c.JSON(http.StatusOK, cocktail_details)
	}
}
