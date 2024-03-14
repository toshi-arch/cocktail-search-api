package repository

import (
	model_database "api/model_database"
	model_detail "api/model_detail"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type CocktailRepository struct {
	DB *gorm.DB
}

func NewCocktailRepository(db *gorm.DB) *CocktailRepository {
	return &CocktailRepository{DB: db}
}

func (r *CocktailRepository) GetAllCocktails() (*[]model_database.Cocktails, error) {
	cocktails := []model_database.Cocktails{}
	err := r.DB.Find(&cocktails).Error
	return &cocktails, err
}

func (r *CocktailRepository) GetCocktailByName(cocktail_name string) (*model_database.Cocktails, error) {
	cocktail := model_database.Cocktails{}
	err := r.DB.Select([]string{"id", "name", "recipe"}).
		Where("Name = ?", cocktail_name).
		First(&cocktail).Error
	return &cocktail, err
}

/*func (r *CocktailRepository) GetCocktailByIngredient(ingredient_id int) ([]model_detail.CocktailName, error) {
	cocktails := []model_detail.CocktailName{}
	err := r.DB.Table("cocktails").
		Select("cocktails.name").
		Where("Ingredient_id = ?", ingredient_id).
		Joins("left join ingredients_cocktails on cocktails.id = ingredients_cocktails.cocktail_id").
		Find(&cocktails).Error
	return cocktails, err
}*/

func (r *CocktailRepository) GetCocktailByIngredient(ingredient_id int) ([]model_detail.CocktailName, error) {
	cocktails := []model_detail.CocktailName{}
	err := r.DB.Table("cocktails").
		Select("cocktails.name").
		Where("Ingredient_id = ?", ingredient_id).
		Joins("left join ingredients_cocktails on cocktails.id = ingredients_cocktails.cocktail_id").
		Find(&cocktails).Error
	return cocktails, err
}
