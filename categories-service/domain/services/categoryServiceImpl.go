package services

import (
	"fmt"
	//"net/http"
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
	fmt.Println("Inside category service")
	fmt.Println("category ",category)
	// //calculate transaction points earned
	// points := t.CalculateTransactionPoints(transactionAmount)
	// transaction := &models.Transaction{
	// 	TransactionPoints: points,
	// 	UserId:            transactionAmount.UserId,
	// }

	// //Fetch user transaction details from DB
	// userTransactionPoints, err := t.transactionRepository.GetTransactionPointsByUserIdFromDB(transaction.UserId)

	//If there is no record for the given userId, create a new record
	// if err != nil {
	// 	if err.Code == http.StatusNotFound {
	// 		err_ := t.transactionRepository.AddTransactionPointsFromDB(transaction)
	// 		if err_ != nil {
	// 			return err_
	// 		} else {
	// 			return nil
	// 		}
	// 	} else {
	// 		return err
	// 	}
	// }

	
	err := t.categoryRepository.AddCategoryToDB(category)
	if err != nil {
		return err
	}
	return nil
}
func (t CategoryServiceImpl) GetAllCategory() ([]models.Category,*apperros.AppError) {
	fmt.Println("Inside get all category service")
	//fmt.Println("category ",category)
	// //calculate transaction points earned
	// points := t.CalculateTransactionPoints(transactionAmount)
	// transaction := &models.Transaction{
	// 	TransactionPoints: points,
	// 	UserId:            transactionAmount.UserId,
	// }

	// //Fetch user transaction details from DB
	// userTransactionPoints, err := t.transactionRepository.GetTransactionPointsByUserIdFromDB(transaction.UserId)

	//If there is no record for the given userId, create a new record
	// if err != nil {
	// 	if err.Code == http.StatusNotFound {
	// 		err_ := t.transactionRepository.AddTransactionPointsFromDB(transaction)
	// 		if err_ != nil {
	// 			return err_
	// 		} else {
	// 			return nil
	// 		}
	// 	} else {
	// 		return err
	// 	}
	// }

//	var categories []models.Category
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

	// Fetch payment modes for the given user
	category, err := p.categoryRepository.GetCategoryFromDB(category_id)
	if err != nil {
		fmt.Println(err)
		logger.Error(err)
		return nil, err
	}
	return category, nil
}
func (p CategoryServiceImpl) DeleteCategories(categories []string) (bool,*apperros.AppError) {

	
	 _,err := p.categoryRepository.DeleteCategoriesFromDB(categories)
	if err != nil {
		fmt.Println(err)
		logger.Error(err)
		return  false,err
	}
	return true,nil
}
func (p CategoryServiceImpl) DeleteCategoryByID(category_id string) (*apperros.AppError) {

	// Fetch payment modes for the given user
	_, err := p.categoryRepository.DeleteCategoryByIDFromDB(category_id)
	if err != nil {
		fmt.Println(err)
		logger.Error(err)
		return  err
	}
	return nil
}
func (p CategoryServiceImpl) UpdateCategoryByID(category_id string,category *models.Category) (bool,*apperros.AppError) {

	// Fetch payment modes for the given user
	 _,err := p.categoryRepository.UpdateCategoryByIdFromDB(category_id,category)
	if err != nil {
		fmt.Println(err)
		logger.Error(err)
		return false, err
	}
	return true,nil
}


// func (t TransactionServiceImpl) GetTransactionPointsByUserId(userId string) (int, *apperros.AppError) {
// 	userTransactionPoints, err := t.transactionRepository.GetTransactionPointsByUserIdFromDB(userId)
// 	if err != nil {
// 		return -1, err
// 	}
// 	return userTransactionPoints, nil
// }

// func (t TransactionServiceImpl) CalculateTransactionPoints(transactionAmount *models.TransactionAmount) int {
// 	//For every 100 Rupees user will get 1 transaction point
// 	var points int
// 	points = transactionAmount.Amount / 100
// 	return points
// }

// func (t TransactionServiceImpl) UseTransactionPoints(transactionAmount *models.TransactionAmount) (bool, *models.TransactionAmount, *apperros.AppError) {
// 	//Fetch user transaction details from DB
// 	userTransactionPoints, err := t.transactionRepository.GetTransactionPointsByUserIdFromDB(transactionAmount.UserId)

// 	if err != nil {
// 		return false, transactionAmount, err
// 	} else {
// 		// For every 1 transaction point, user will get a discount of 1 Rupee
// 		if userTransactionPoints == 0 {
// 			return false, transactionAmount, apperros.NewExpectationFailed("You have 0 transaction points")
// 		} else if userTransactionPoints < transactionAmount.Amount {
// 			//user using all his transaction points (Will only apply if order amount is greater than points)
// 			transactionAmount.Amount -= userTransactionPoints
// 			userTransactionPoints = 0

// 			err_ := t.UpdateTransactionPoints(userTransactionPoints, transactionAmount.UserId)
// 			if err_ != nil {
// 				log.Error("Failed to update transaction points")
// 				return false, transactionAmount, err_
// 			}

// 			return true, transactionAmount, nil
// 		} else {
// 			log.Info("Cannot use transaction points as order amount is lesser than available points")
// 			return false, transactionAmount, apperros.NewExpectationFailed("Cannot use transaction points as order amount is lesser than available points")
// 		}
// 	}
// }

// func (t TransactionServiceImpl) UpdateTransactionPoints(transactionPoint int, userId string) *apperros.AppError {
// 	transaction := &models.Transaction{
// 		UserId:            userId,
// 		TransactionPoints: transactionPoint,
// 	}
// 	err := t.transactionRepository.UpdateTransactionPointsToDB(transaction)
// 	return err
// }