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

	taskRepo := repository.NewTaskRepository(db)
	taskHandler := handler.NewTaskhandler(taskRepo)
	// サーバ立ち上げ
	g := gin.Default()
	r := g.Group("")

	{
		r.GET("/cocktails", taskHandler.GetAllCocktails)
		r.GET("/cocktail/:cocktail_name", taskHandler.GetOneCocktail)
	}

	g.Run(":8080")
}
