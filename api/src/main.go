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

	Name string `json:"name"`
	Alcohol int `json:"alcohol"`
	Recipe string `json:"recipe"`
}

type Ingredients struct {
	gorm.Model

	Name string //`json:"name"`
	Type int `json:"type"`
	Alcohol int //`json:"alcohol"`
}

type IngredientsCocktails struct {
	gorm.Model

	Ingredient_id int `json:"ingredients_id"`
	Cocktail_id int `json:"cocktails_id"`
	Amount int `json:"amount"`
	Unit int `json:"unit"`
}

type Gintonic struct {
	IngredientsCocktails
	Cocktails
	Ingredients
}

func gormConnect() *gorm.DB {
	DBMS := "mysql"
	USER := "admin"
	PASS := "password"
	PROTOCOL := "tcp(liquor-mysql-dev)"
	DBNAME := "liquor"
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME

	db, err := gorm.Open(DBMS, CONNECT + "?parseTime=true")

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

	//ジントニックのレコードを取得
	gintonic := []Gintonic{}
	db.Table("ingredients_cocktails").Select([]string{"cocktails.name","ingredients.name", "ingredients_cocktails.amount", "ingredients_cocktails.unit"}).Joins("left join Cocktails on ingredients_cocktails.cocktail_id = cocktails.id").Joins("left join Ingredients on ingredients_cocktails.ingredient_id = ingredients.id").Where(&Cocktails{Name: "ジントニック"}).Scan(&gintonic)
	
	// サーバ立ち上げ
	r := gin.Default()

	r.GET("/cocktails", func(c *gin.Context) {
		c.JSON(http.StatusOK, cocktails)
	})

	//ジントニックのhttp接続
	r.GET("/cocktails/gintonic", func(c *gin.Context) {
		c.JSON(http.StatusOK, gintonic)
	})

	r.Run(":8080")
}
