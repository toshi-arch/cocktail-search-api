package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Cocktails struct {
	gorm.Model

	Name    string `json:"cocktail_name"`
	Alcohol int    `json:"cocktail_alcohol"`
	Recipe  string `json:"recipe"`
}

type Ingredients struct {
	gorm.Model

	Name    string `json:"ingredient_name"`
	Type    int    `json:"type"`
	Alcohol int    `json:"ingredient_alcohol"`
}

type IngredientsCocktails struct {
	gorm.Model

	Ingredient_id int `json:"ingredient_id"`
	Cocktail_id   int `json:"cocktail_id"`
	Amount        int `json:"ingredient_amount"`
	Unit          int `json:"ingredient_unit"`
}

type CocktailDetails struct {
	Name        string
	Recipe      string
	Ingredients []IngredientDetails
}

type IngredientDetails struct {
	Name   string
	Amount int
	Unit   int
}

func gormConnect() *gorm.DB {
	DBMS := "mysql"
	USER := "admin"
	PASS := "password"
	PROTOCOL := "tcp(liquor-mysql-dev)"
	DBNAME := "liquor"
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME

	db, err := gorm.Open(DBMS, CONNECT+"?parseTime=true")

	if err != nil {
		panic(err.Error())
	}
	fmt.Println("db connected: ", &db)
	return db
}

func main() {
	// サーバ立ち上げ
	r := gin.Default()
	r.GET("/cocktail/:cocktail_name", func(c *gin.Context) {
		cocktail_name := c.Param("cocktail_name")

		// dbに接続
		db := gormConnect()

		defer db.Close()
		db.LogMode(true)

		cocktail_details := new(CocktailDetails)
		ingredient_details := []IngredientDetails{}

		// cocktailsテーブルのレコードを取得
		target_cocktail := Cocktails{}

		if err := db.Select([]string{"id", "name", "recipe"}).
			Where("Name = ?", cocktail_name).
			First(&target_cocktail).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "申し訳ございません。そのレシピは存在しません。",
			})
		} else {
			cocktail_details.Name = target_cocktail.Name
			cocktail_details.Recipe = target_cocktail.Recipe

			//IngredientDetailsの要素を取得
			db.Table("ingredients").
			Select("ingredients.name, ingredients_cocktails.amount, ingredients_cocktails.unit").
			Where("Cocktail_id = ?", int(target_cocktail.ID)).
			Joins("left join ingredients_cocktails on ingredients.id = ingredients_cocktails.ingredient_id").
			Find(&ingredient_details)

			cocktail_details.Ingredients = ingredient_details

			c.JSON(http.StatusOK, cocktail_details)
		}

	})

	r.Run(":8080")
}
