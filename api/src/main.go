package main

import (
	"fmt"
	"net/http"
	//"time"


	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Cocktails struct {
	gorm.Model
	/*ID int `gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`*/

	Name string `json:"name"`
	Alcohol int `json:"alcohol"`
	Recipe string `json:"recipe"`

}

func gormConnect() *gorm.DB {
	DBMS := "mysql"
	USER := "admin"
	PASS := "password"
	PROTOCOL := "tcp(liquor-mysql-dev)"
	DBNAME := "liquor"
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME

	db, err := gorm.Open(DBMS, CONNECT)

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

	// サーバ立ち上げ
	r := gin.Default()

	r.GET("/cocktails", func(c *gin.Context) {
		c.JSON(http.StatusOK, cocktails)
	})

	r.Run(":8080")
}
