package services

import (
	"fmt"
	//"net/http"
	"github.com/cdp-team3/categories-service/domain/models"
	//"github.com/cdp-team3/categories-service/app/grpcs"
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
	fmt.Println("Inside category service")
	fmt.Println("category ",category)
	

	
	err := t.categoryRepository.AddCategoryToDB(category)
	if err != nil {
		return err
	}
	return nil
}
func (t CategoryServiceImpl) GetAllCategory() ([]models.Category,*apperros.AppError) {
	fmt.Println("Inside get all category service")
	
	result,err := t.categoryRepository.FindAllCategoryFromDB()
	if err != nil {
		fmt.Println(err)
		logger.Error(err)
		return nil,err
	}
	fmt.Println(result)
	return result,nil
}

func (p CategoryServiceImpl) GetCategory(category_id string) (*models.Category, *apperros.AppError) {

	
	category, err := p.categoryRepository.GetCategoryFromDB(category_id)
	if err != nil {
		fmt.Println(err)
		logger.Error(err)
		return nil, err
	}
	return category, nil
}
func (p CategoryServiceImpl) DeleteCategories(categories []string) (bool,*apperros.AppError) {

	
	res, err := p.categoryRepository.DeleteCategoriesFromDB(categories)
	if err != nil {
		fmt.Println(err)
		logger.Error(err)
		return  res,err
	}
	return res,nil
}
func (p CategoryServiceImpl) DeleteCategoryByID(category_id string) (bool,*apperros.AppError) {


     
	// if res == true{
	//	return false,apperrors.NewConflictRequestError(err.Error())
	// } else{
		res,err := p.categoryRepository.DeleteCategoryByIDFromDB(category_id)
		if err != nil {
			fmt.Println(err)
			logger.Error(err)
			return  res,err
		}
		return res,nil
	// }
	
}
func (p CategoryServiceImpl) UpdateCategoryByID(category_id string,category *models.Category) (bool,*apperros.AppError) {


	 _,err := p.categoryRepository.UpdateCategoryByIdFromDB(category_id,category)
	if err != nil {
		fmt.Println(err)
		logger.Error(err)
		return false, err
	}
	return true,nil
}

