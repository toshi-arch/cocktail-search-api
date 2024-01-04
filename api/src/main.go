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

type Ingredients_Cocktails struct {
	gorm.Model

	Ingredient_id int `json:"ingredient_id"`
	Cocktail_id int `json:"cocktail_id"`
	Amount int `json:"ingredient_amount"`
	Unit   int `json:"ingredient_unit"`
}

type CocktailDetails struct {
	CocktailName string
	CocktailRecipe string
	Ingredient []IngredientDetails
}

type IngredientDetails struct {
	IngredientName string
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
	// dbに接続
	db := gormConnect()

	defer db.Close()
	db.LogMode(true)

	// cocktailsテーブルの全レコードを取得
	cocktails := []Cocktails{}
	db.Find(&cocktails) // 全レコード

	// ingredientsテーブルの全レコードを取得
	ingredients := []Ingredients{}
	db.Find(&ingredients) // 全レコード

	// ingredients_cocktailsテーブルの全レコードを取得
	ingredients_cocktails := []Ingredients_Cocktails{}
	db.Find(&ingredients_cocktails) // 全レコード

	//カクテル名の入力
	fmt.Println("カクテル名を入力してください。")
	var str string
    fmt.Scan(&str)

	//レシピの検索
	x := new(CocktailDetails)
	y := []IngredientDetails{}

	for i1,v1 := range cocktails {
		if cocktails[i1].Name == str {
			x.CocktailName = v1.Name
			x.CocktailRecipe = v1.Recipe
			for _,v2 := range ingredients_cocktails {
				if v1.ID == uint(v2.Cocktail_id) {
					for _,v3 := range ingredients {
						if int(v3.ID) == v2.Ingredient_id {
							y1 := IngredientDetails{Amount: v2.Amount, Unit: v2.Unit, IngredientName: v3.Name}
							y = append(y, y1)
							x.Ingredient = y
						}
					}
				}
			}
		}
	}

	// サーバ立ち上げ
	r := gin.Default()

	r.GET("/cocktails", func(c *gin.Context) {
		c.JSON(http.StatusOK, cocktails)
	})

	r.GET("/cocktail_details", func(c *gin.Context) {
		c.JSON(http.StatusOK, x)
	})

	r.Run(":8080")
}
