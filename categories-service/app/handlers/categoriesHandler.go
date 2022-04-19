package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"github.com/google/uuid"
	apperros "github.com/cdp-team3/categories-service/app-errors"
	"github.com/cdp-team3/categories-service/domain/models"
	"github.com/cdp-team3/categories-service/domain/services"
)

var validate = validator.New()

type CategoryHandler struct {
	categoryService services.CategoryService
}

func NewCategoryHandler(categoryService services.CategoryService) CategoryHandler {
	return CategoryHandler{categoryService: categoryService}
}
func toPersistedDynamodbEntitySA(o models.Category) *models.Category {
	return &models.Category{

		CategoryId :        uuid.New().String(),
		CategoryDescription: o.CategoryDescription,
		
	}
}
func (th CategoryHandler) AddCategory(c *gin.Context) {
	//userId := c.Param("userId")
	var category models.Category

	if err := c.BindJSON(&category); err != nil {
		c.Error(err)
		err_ := apperros.NewBadRequestError(err.Error())
		c.JSON(err_.Code, gin.H{"message": err_.Message})
		return
	}

	
	//transactionAmount.UserId = userId
	// err := th.categoryService.AddTransactionPoints(&category)
	// if err != nil {
	// 	c.Error(err.Error())
	// 	c.JSON(err.Code, gin.H{"message": err.Message})
	// 	return
	// }
	
	categoryRecord := toPersistedDynamodbEntitySA(category)
	fmt.Println(categoryRecord)
	//validate request body
	// if validationErr := validate.Struct(&categoryRecord); validationErr != nil {
	// 	c.Error(validationErr)
	// 	c.JSON(http.StatusBadRequest, gin.H{"message": validationErr.Error()})
	// 	return
	// }
	err := th.categoryService.AddCategory(categoryRecord)
		if err != nil {
		c.Error(err.Error())
		c.JSON(err.Code, gin.H{"message": err.Message})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Category added successfully"})
}
func (th CategoryHandler) GetAllCategory(c *gin.Context) {
	//userId := c.Param("userId")
	//var category models.Category

	// if err := c.BindJSON(&category); err != nil {
	// 	c.Error(err)
	// 	err_ := apperros.NewBadRequestError(err.Error())
	// 	c.JSON(err_.Code, gin.H{"message": err_.Message})
	// 	return
	// }

	// //validate request body
	// if validationErr := validate.Struct(&category); validationErr != nil {
	// 	c.Error(validationErr)
	// 	c.JSON(http.StatusBadRequest, gin.H{"message": validationErr.Error()})
	// 	return
	// }
	//transactionAmount.UserId = userId
	// err := th.categoryService.AddTransactionPoints(&category)
	// if err != nil {
	// 	c.Error(err.Error())
	// 	c.JSON(err.Code, gin.H{"message": err.Message})
	// 	return
	// }
	result,err := th.categoryService.GetAllCategory()
		if err != nil {
		c.Error(err.Error())
		c.JSON(err.Code, gin.H{"message": err.Message})
		return
	}
	c.JSON(http.StatusOK, gin.H{"categories": result})
}


func (th CategoryHandler) GetCategory(c *gin.Context) {
	category_id := c.Param("category_id")
	category, err := th.categoryService.GetCategory(category_id)
	if err != nil {
		c.Error(err.Error())
		c.JSON(err.Code, gin.H{"message": err.Message})
		return
	}
	c.JSON(http.StatusOK, category)
}

func (th CategoryHandler) DeleteCategories(c *gin.Context) {
	
	 err := th.categoryService.DeleteCategories()
	if err != nil {
		c.Error(err.Error())
		c.JSON(err.Code, gin.H{"message": err.Message})
		return
	}
	c.JSON(http.StatusOK,gin.H{"message": "Categories deleted successfully"})
}
func (th CategoryHandler) DeleteCategory(c *gin.Context) {
	category_id := c.Param("category_id")
	 err := th.categoryService.DeleteCategoryByID(category_id)
	if err != nil {
		c.Error(err.Error())
		c.JSON(err.Code, gin.H{"message": err.Message})
		return
	}
	c.JSON(http.StatusOK,gin.H{"message": "Category deleted successfully"})
}
func (th CategoryHandler) UpdateCategory(c *gin.Context) {
	category_id := c.Param("category_id")
	var category *models.Category
      //  defer cancel()
	    // validate the request body
        if err := c.BindJSON(&category); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
	res, err := th.categoryService.UpdateCategoryByID(category_id,category)
	if err != nil {
		c.Error(err.Error())
		c.JSON(err.Code, gin.H{"message": err.Message})
		return
	}
	c.JSON(http.StatusOK,res)
}

