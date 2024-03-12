package model

import (
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

