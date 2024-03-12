package repository

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	model_detail "api/model_detail"
)

type IngredientRepository struct {
	DB *gorm.DB
}

func NewIngredientRepository(db *gorm.DB) *IngredientRepository {
	return &IngredientRepository{DB: db}
}

func (r *IngredientRepository) GetIngredientDetail(cocktail_id int) ([]model_detail.Ingredient, error) {
	ingredient_detail := []model_detail.Ingredient{}
	err := r.DB.Table("ingredients").
		Select("ingredients.name, ingredients_cocktails.amount, ingredients_cocktails.unit").
		Where("Cocktail_id = ?", cocktail_id).
		Joins("left join ingredients_cocktails on ingredients.id = ingredients_cocktails.ingredient_id").
		Find(&ingredient_detail).Error
	return ingredient_detail, err
}
