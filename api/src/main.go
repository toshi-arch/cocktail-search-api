package main

import (
	"github.com/gin-gonic/gin"
	"api/infra"
	"api/repository"
	"api/handler"
)

func main() {
	// dbに接続
	db := infra.GormConnect()
	defer db.Close()
	db.LogMode(true)

	taskRepoCocktail := repository.NewCocktailRepository(db)
	taskRepoIngredient := repository.NewIngredientRepository(db)
	taskCocktailHandler := handler.NewCocktailHandler(taskRepoCocktail, taskRepoIngredient)
	taskIngredientHandler := handler.NewIngredientHandler(taskRepoCocktail, taskRepoIngredient)
	// サーバ立ち上げ
	g := gin.Default()
	r := g.Group("")

	{
		r.GET("/cocktails", taskCocktailHandler.GetAllCocktails)
		r.GET("/cocktail/:cocktail_name", taskCocktailHandler.GetCocktailByName)
		r.GET("/ingredient/:ingredient_name", taskCocktailHandler.GetCocktailByIngredient)
		r.GET("ingredients", taskIngredientHandler.GetAllIngredients)
	}

	g.Run(":8080")
}
