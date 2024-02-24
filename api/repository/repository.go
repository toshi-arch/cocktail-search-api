package repository

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"api/model"
)

type TaskRepository struct {
	DB *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{DB: db}
}

func (r *TaskRepository) GetAllCocktails() (*[]model.Cocktails, error) {
	cocktails := []model.Cocktails{}
	err := r.DB.Find(&cocktails).Error
	return &cocktails, err
}

func (r *TaskRepository) GetTargetCocktail(cocktail_name string) (*model.Cocktails, error) {
	target_cocktail := model.Cocktails{}
	err := r.DB.Select([]string{"id", "name", "recipe"}).
		Where("Name = ?", cocktail_name).
		First(&target_cocktail).Error
	return &target_cocktail, err
}

func (r *TaskRepository) GetIngredientDetails(target_cocktail_id int) ([]model.IngredientDetails, error) {
	ingredient_details := []model.IngredientDetails{}
	err := r.DB.Table("ingredients").
		Select("ingredients.name, ingredients_cocktails.amount, ingredients_cocktails.unit").
		Where("Cocktail_id = ?", target_cocktail_id).
		Joins("left join ingredients_cocktails on ingredients.id = ingredients_cocktails.ingredient_id").
		Find(&ingredient_details).Error
	return ingredient_details, err
}
