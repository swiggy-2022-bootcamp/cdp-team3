package repository

import (
	"github.com/cdp-team3/categories-service/domain/models"
	apperrors "github.com/cdp-team3/categories-service/app-errors"
	
)
type CategoryRepository interface {
	
	DBHealthCheck() bool

	 AddCategoryToDB(category *models.Category) *apperrors.AppError
	 FindAllCategoryFromDB() ([]models.Category, *apperrors.AppError)
	 GetCategoryFromDB(category_id string)(*models.Category, *apperrors.AppError)
	 DeleteCategoryByIDFromDB(category_id string) (bool,*apperrors.AppError)
	 DeleteCategoriesFromDB(categoryIds []string) (bool,*apperrors.AppError)
	 UpdateCategoryByIdFromDB(category_id string,category *models.Category) (bool, *apperrors.AppError)
	
}