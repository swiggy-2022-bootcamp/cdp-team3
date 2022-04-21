package repository

import (
	"github.com/cdp-team3/categories-service/domain/models"
	apperrors "github.com/cdp-team3/categories-service/app-errors"
	
)
type CategoryRepository interface {
	AddCategoryToDB(category *models.Category) *apperrors.AppError
	FindAllCategoryFromDB() ([]models.Category, *apperrors.AppError)
	// DeleteCategoryByIDFromDB(categoryId string) (bool, *apperros.AppError)
	// GetCategoryByCategoryIdFromDB(categoryId string) (*models.Category, *apperros.AppError)
	 DBHealthCheck() bool
	 GetCategoryFromDB(category_id string)(*models.Category, *apperrors.AppError)
	 DeleteCategoryByIDFromDB(category_id string) (bool,*apperrors.AppError)
	 DeleteCategoriesFromDB(categoryIds []string) (bool,*apperrors.AppError)
	 UpdateCategoryByIdFromDB(category_id string,category *models.Category) (bool, *apperrors.AppError)
	// UpdateCategoryByIdFromDB(categoryId string,category *models.Category) *apperros.AppError
}