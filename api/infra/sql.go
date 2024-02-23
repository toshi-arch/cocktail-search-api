package infra

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	
	"api/model"
)

func GetAll(db *gorm.DB, cocktails *[]model.Cocktails) {
	db.Find(&cocktails)
}

func GetTargetCocktail(db *gorm.DB, cocktail_name string, target_cocktail *model.Cocktails) error{
	err := db.Select([]string{"id", "name", "recipe"}).
		Where("Name = ?", cocktail_name).
		First(&target_cocktail).Error
	return err
}

func GetIngredientDetails(db *gorm.DB, target_cocktail_id int, ingredient_details *[]model.IngredientDetails) {
	db.Table("ingredients").
	Select("ingredients.name, ingredients_cocktails.amount, ingredients_cocktails.unit").
	Where("Cocktail_id = ?", target_cocktail_id).
	Joins("left join ingredients_cocktails on ingredients.id = ingredients_cocktails.ingredient_id").
	Find(&ingredient_details)
}