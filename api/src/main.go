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

	repoCocktail := repository.NewCocktailRepository(db)
	repoIngredient := repository.NewIngredientRepository(db)
	cocktailHandler := handler.NewCocktailHandler(repoCocktail, repoIngredient)
	ingredientHandler := handler.NewIngredientHandler(repoCocktail, repoIngredient)
	// サーバ立ち上げ
	g := gin.Default()
	r := g.Group("")

	{
		r.GET("/cocktails", cocktailHandler.GetCocktails)
		r.GET("/cocktail/:cocktail_name", cocktailHandler.GetCocktailByName)
		r.GET("/cocktails/:ingredient_name", cocktailHandler.GetCocktailByIngredient)
		r.GET("/ingredients", ingredientHandler.GetIngredients)
	}

	g.Run(":8080")
}
