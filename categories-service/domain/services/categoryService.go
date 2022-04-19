package services

import (
	"github.com/cdp-team3/categories-service/domain/models"
	apperros "github.com/cdp-team3/categories-service/app-errors"
)

type CategoryService interface {
	AddCategory(category *models.Category) *apperros.AppError
	 GetAllCategory() ([]models.Category, *apperros.AppError)
	 GetCategory(category_id string) (*models.Category, *apperros.AppError)
	 DeleteCategoryByID(category_id string) *apperros.AppError
	 DeleteCategories([]string) (bool,*apperros.AppError)
	 UpdateCategoryByID(category_id string,category *models.Category) (bool,*apperros.AppError)
	
	
}