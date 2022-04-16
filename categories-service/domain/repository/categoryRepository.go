package repository

import (
	"github.com/cdp-team3/categories-service/domain/models"
	apperros "github.com/cdp-team3/categories-service/app-errors"
	
)
type CategoryRepository interface {
	AddCategoryToDB(category *models.Category) *apperros.AppError
	FindAllCategoryFromDB() ([]models.Category, *apperros.AppError)
	// DeleteCategoryByIDFromDB(categoryId string) (bool, *apperros.AppError)
	// GetCategoryByCategoryIdFromDB(categoryId string) (*models.Category, *apperros.AppError)
	 DBHealthCheck() bool
	 GetCategoryFromDB(category_id string)(*models.Category, *apperros.AppError)
	 DeleteCategoryByIDFromDB(category_id string) *apperros.AppError
	 UpdateCategoryByIDFromDB1(category_id string,category *models.Category) (bool, *apperros.AppError)
	// UpdateCategoryByIdFromDB(categoryId string,category *models.Category) *apperros.AppError
}