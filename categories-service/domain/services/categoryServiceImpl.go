package services

import (
	"fmt"
	
    "github.com/cdp-team3/categories-service/domain/models"

	apperros "github.com/cdp-team3/categories-service/app-errors"
	"github.com/cdp-team3/categories-service/domain/repository"
	"github.com/cdp-team3/categories-service/utils/logger"
)

type CategoryServiceImpl struct {
	categoryRepository repository.CategoryRepository
}

func NewCategoryServiceImpl(categoryRepository repository.CategoryRepository) CategoryService {
	return &CategoryServiceImpl{categoryRepository: categoryRepository}
}

func (t CategoryServiceImpl) AddCategory(category *models.Category) *apperros.AppError {
	
	
	err := t.categoryRepository.AddCategoryToDB(category)
	if err != nil {
		return err
	}
	return nil
}
func (t CategoryServiceImpl) GetAllCategory() ([]models.Category,*apperros.AppError) {

	
	result,err := t.categoryRepository.FindAllCategoryFromDB()
	if err != nil {
	
		logger.Error(err)
		return nil,err
	}
	
	return result,nil
}

func (p CategoryServiceImpl) GetCategory(id string) (*models.Category, *apperros.AppError) {

	
	category, err := p.categoryRepository.GetCategoryFromDB(id)
	if err != nil {
	
		logger.Error(err)
		return nil, err
	}
	return category, nil
}
func (p CategoryServiceImpl) DeleteCategories(categories []string) (bool,*apperros.AppError) {

	
	res, err := p.categoryRepository.DeleteCategoriesFromDB(categories)
	if err != nil {

		logger.Error(err)
		return  res,err
	}
	return res,nil
}
func (p CategoryServiceImpl) DeleteCategoryByID(id string) (bool,*apperros.AppError) {


		res,err := p.categoryRepository.DeleteCategoryByIDFromDB(id)
		if err != nil {
			fmt.Println(err)
			logger.Error(err)
			return  res,err
		}
		return res,nil
	
	
}
func (p CategoryServiceImpl) UpdateCategoryByID(id string,category *models.Category) (bool,*apperros.AppError) {


	 _,err := p.categoryRepository.UpdateCategoryByIdFromDB(id,category)
	if err != nil {
		
		logger.Error(err)
		return false, err
	}
	return true,nil
}

