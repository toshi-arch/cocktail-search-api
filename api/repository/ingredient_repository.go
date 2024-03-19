package repository

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	modelDatabase "api/model_database"
	modelDetail "api/model_detail"
)

type IngredientRepository struct {
	DB *gorm.DB
}

func NewIngredientRepository(db *gorm.DB) *IngredientRepository {
	return &IngredientRepository{DB: db}
}

func (r *IngredientRepository) GetIngredients() (*[]modelDatabase.Ingredients, error) {
	ingredients := []modelDatabase.Ingredients{}
	err := r.DB.Find(&ingredients).Error
	return &ingredients, err
}

func (r *IngredientRepository) GetIngredientByName(ingredientName string) (*modelDatabase.Ingredients, error) {
	ingredient := modelDatabase.Ingredients{}
	err := r.DB.Select([]string{"id", "name"}).
		Where("Name = ?", ingredientName).
		First(&ingredient).Error
	return &ingredient, err
}

func (r *IngredientRepository) GetIngredientsByCocktailId(cocktailId int) (*[]modelDetail.Ingredient, error) {
	ingredientDetail := []modelDetail.Ingredient{}
	err := r.DB.Table("ingredients").
		Select("ingredients.name, ingredients_cocktails.amount, ingredients_cocktails.unit").
		Where("Cocktail_id = ?", cocktailId).
		Joins("left join ingredients_cocktails on ingredients.id = ingredients_cocktails.ingredient_id").
		Find(&ingredientDetail).Error
	return &ingredientDetail, err
}
