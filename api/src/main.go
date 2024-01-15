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
	Cocktail_id int `json:"cocktail_id"`
	Amount int `json:"ingredient_amount"`
	Unit   int `json:"ingredient_unit"`
}

type CocktailDetails struct {
	Name string
	Recipe string
	Ingredients []IngredientDetails
}

type IngredientDetails struct {
	Name string
	Amount int
	Unit int
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

		x := new(CocktailDetails)
		y := []IngredientDetails{}
	
		// cocktailsテーブルのレコードを取得
		cocktails := []Cocktails{}
		db.Select([]string{"id", "name", "recipe"}).
		Where("Name = ?", cocktail_name).
		Find(&cocktails)  //First(&cocktails)でもいいかも

		x.Name = cocktails[0].Name
		x.Recipe = cocktails[0].Recipe
		
		// ingredients_cocktailsテーブルのレコードを取得
		ingredients_cocktails := []IngredientsCocktails{}
		db.Select([]string{"ingredient_id", "amount", "unit"}).
		Where("Cocktail_id = ?", int(cocktails[0].ID)).
		Find(&ingredients_cocktails) 
		 
		for _,v := range ingredients_cocktails {
			y1 := IngredientDetails{Amount: v.Amount, Unit: v.Unit}
			y = append(y, y1)
		}
		
		// ingredientsテーブルのレコードを取得
		ingredients := []Ingredients{}
		for i,v := range ingredients_cocktails {
			db.Select([]string{"name"}).
			Where("ID = ?", v.Ingredient_id).
			Find(&ingredients)
			y[i].Name = ingredients[0].Name	
		}
		x.Ingredients = y

		c.JSON(http.StatusOK, x)
	})

	r.Run(":8080")
}
