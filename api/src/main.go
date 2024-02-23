package main

import (
	"github.com/gin-gonic/gin"
	"api/infra"
	"api/service"
)

func main() {
	// dbに接続
	db := infra.GormConnect()
	defer db.Close()
	db.LogMode(true)

	// サーバ立ち上げ
	g := gin.Default()
	r := g.Group("")

	{
		r.GET("/cocktails", service.GetAll)
		r.GET("/cocktail/:cocktail_name", service.GetOneCocktail)
	}

	g.Run(":8080")
}
