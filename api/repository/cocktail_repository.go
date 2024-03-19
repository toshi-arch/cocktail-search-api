package repository

import (
	modelDatabase "api/model_database"
	modelDetail "api/model_detail"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type CocktailRepository struct {
	DB *gorm.DB
}

func NewCocktailRepository(db *gorm.DB) *CocktailRepository {
	return &CocktailRepository{DB: db}
}

func (r *CocktailRepository) GetCocktails() (*[]modelDatabase.Cocktails, error) {
	cocktails := []modelDatabase.Cocktails{}
	err := r.DB.Find(&cocktails).Error
	return &cocktails, err
}

func (r *CocktailRepository) GetCocktailByName(cocktailName string) (*modelDatabase.Cocktails, error) {
	cocktail := modelDatabase.Cocktails{}
	err := r.DB.Select([]string{"id", "name", "recipe"}).
		Where("Name = ?", cocktailName).
		First(&cocktail).Error
	return &cocktail, err
}

func (r *CocktailRepository) GetCocktailByIngredient(ingredientId int) ([]modelDetail.CocktailName, error) {
	cocktails := []modelDetail.CocktailName{}
	err := r.DB.Table("cocktails").
		Select("cocktails.name").
		Where("Ingredient_id = ?", ingredientId).
		Joins("left join ingredients_cocktails on cocktails.id = ingredients_cocktails.cocktail_id").
		Find(&cocktails).Error
	return cocktails, err
}
