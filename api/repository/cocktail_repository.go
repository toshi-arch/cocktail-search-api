package repository

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	model_database "api/model_database"
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
