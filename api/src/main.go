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
		r.GET("/cocktails", func(c *gin.Context) {
			service.GetAll(c, db)
		})
		r.GET("/cocktail/:cocktail_name", func(c *gin.Context) {
			service.GetOneCocktail(c, db)
		})
	}

	g.Run(":8080")
}
