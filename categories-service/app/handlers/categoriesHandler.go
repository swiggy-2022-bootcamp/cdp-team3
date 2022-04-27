package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"github.com/cdp-team3/categories-service/app/grpcs"
	"github.com/cdp-team3/categories-service/app/grpcs/products"
	//"github.com/google/uuid"
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

func (th CategoryHandler) AddCategory() (gin.HandlerFunc){
	return func(ctx *gin.Context) {
	var category models.Category

	if err := ctx.BindJSON(&category); err != nil {
		ctx.Error(err)
		er := apperros.NewBadRequestError(err.Error())
		ctx.JSON(er.Code, gin.H{"message": er.Message})
		return
	}

	// validate request body
	if validationErr := validate.Struct(&category); validationErr != nil {
		ctx.Error(validationErr)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": validationErr.Error()})
		return
	}
	err := th.categoryService.AddCategory(&category)
		if err != nil {
		ctx.Error(err.Error())
		ctx.JSON(err.Code, gin.H{"message": err.Message})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Category added successfully"})
}
}
func (th CategoryHandler) GetAllCategory() (gin.HandlerFunc){
	return func(ctx *gin.Context) {

	
	result,err := th.categoryService.GetAllCategory()
		if err != nil {
		ctx.Error(err.Error())
		ctx.JSON(err.Code, gin.H{"message": err.Message})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"categories": result})
}
}


func (th CategoryHandler) GetCategory()  (gin.HandlerFunc){
	return func(ctx *gin.Context) {
	id := ctx.Param("category_id")
	category, err := th.categoryService.GetCategory(id)
	if err != nil {
		ctx.Error(err.Error())
		ctx.JSON(err.Code, gin.H{"message": err.Message})
		return
	}
	ctx.JSON(http.StatusOK, category)
}
}

func (th CategoryHandler) DeleteCategories() (gin.HandlerFunc){
	return func(ctx *gin.Context) {
	var categoryList []string
		if err := ctx.BindJSON(&categoryList); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		
		// Makes a grpc call to product service and check if any products is associated with the categories we want to delete

		client, _ := grpcs.GetProductsGrpcClient()
		val,_:=client.DeleteCategories(ctx.Request.Context(),&products.CategoriesDeleteRequest{CategoriesId:categoryList})


		if(!val){
			_,err := th.categoryService.DeleteCategories(categoryList)
			if err != nil {
				ctx.Error(err.Error())
				ctx.JSON(err.Code, gin.H{"message": err.Message})
				return
			}
			ctx.JSON(http.StatusOK,gin.H{"message": "Categories deleted successfully"})
		}
		
		ctx.JSON(http.StatusConflict,gin.H{"message": "Categories can't be deleted."})}
}
func (th CategoryHandler) DeleteCategory() (gin.HandlerFunc) {
	return func(ctx *gin.Context){
		id := ctx.Param("category_id")
	
	 // Makes a grpc call to product service and check if any products is associated with the categories we want to delete

		client, _ := grpcs.GetProductsGrpcClient()
		val,_:=client.DeleteCategory(ctx.Request.Context(),&products.CategoryDeleteRequest{CategoryId:id})
		if(!val){
			_,err_ := th.categoryService.DeleteCategoryByID(id)
			if err_ != nil {
				ctx.Error(err_.Error())
				ctx.JSON(err_.Code, gin.H{"message": err_.Message})
				return
			}
			ctx.JSON(http.StatusOK,gin.H{"message": "Category deleted successfully"})
		}
		ctx.JSON(http.StatusConflict,gin.H{"message": "Category can't be deleted."})
	}
}

// Handler function to update the category
func (th CategoryHandler) UpdateCategory() (gin.HandlerFunc) {
	return func(ctx *gin.Context){
	id := ctx.Param("category_id")
	var category *models.Category
      
	    // validate the request body
        if err := ctx.BindJSON(&category); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
	res, err := th.categoryService.UpdateCategoryByID(id,category)
	if err != nil {
		ctx.Error(err.Error())
		ctx.JSON(err.Code, gin.H{"message": err.Message})
		return
	}
	ctx.JSON(http.StatusOK,res)
}
}

