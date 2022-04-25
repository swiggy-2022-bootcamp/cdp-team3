package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	//"github.com/cdp-team3/categories-service/app/grpcs"
	//"github.com/cdp-team3/categories-service/app/grpcs/products"
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
// func toPersistedDynamodbEntitySA(o models.Category) *models.Category {
// 	return &models.Category{

// 		CategoryId :        uuid.New().String(),
// 		CategoryDescription: o.CategoryDescription,
		
// 	}
// }
func (th CategoryHandler) AddCategory() (gin.HandlerFunc){
	return func(ctx *gin.Context) {
	//userId := c.Param("userId")
	var category models.Category

	if err := ctx.BindJSON(&category); err != nil {
		ctx.Error(err)
		er := apperros.NewBadRequestError(err.Error())
		ctx.JSON(er.Code, gin.H{"message": er.Message})
		return
	}

	
	
	// categoryRecord := toPersistedDynamodbEntitySA(category)
	// fmt.Println(categoryRecord)
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
	category_id := ctx.Param("category_id")
	category, err := th.categoryService.GetCategory(category_id)
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
			fmt.Println("Unable to bind json",err)
		//	logger.WithFields(logrus.Fields{"message": err.Error(), "status": http.StatusBadRequest}).
			//	Error("unable to bind json")
			return
		}
		fmt.Println("clist",categoryList)
		// client, _ := grpcs.GetProductsGrpcClient()
		// val,er:=client.DeleteCategories(ctx.Request.Context(),&products.CategoriesDeleteRequest{CategoriesId:categoryList})
		// fmt.Println(val)
		// fmt.Println(er)
		// if err != nil {
		// 	logger.Error("Error while deleting categories: ", err)
		// 	ctx.JSON(http.StatusInternalServerError, app_errors.NewHTTPErrorDTO(http.StatusInternalServerError, err, "Error Getting Shipping Address"))
		// 	return
		// } else if validateShippingAddress(shippingOut) == false {
		// 	//ctx.JSON(http.StatusBadRequest, app_errors.NewHTTPErrorDTO(http.StatusBadRequest, nil, "Shipping Address is not valid."))
		// 	return
		// }
	_,err := th.categoryService.DeleteCategories(categoryList)
	if err != nil {
		ctx.Error(err.Error())
		ctx.JSON(err.Code, gin.H{"message": err.Message})
		return
	}
	ctx.JSON(http.StatusOK,gin.H{"message": "Categories deleted successfully"})}
}
func (th CategoryHandler) DeleteCategory() (gin.HandlerFunc) {
	return func(ctx *gin.Context){
		category_id := ctx.Param("category_id")
	fmt.Println("Inside category id",category_id)
		// client, _ := grpcs.GetProductsGrpcClient()
		// val,err:=client.DeleteCategory(ctx.Request.Context(),&products.CategoryDeleteRequest{CategoryId:category_id})
		// fmt.Println(val)
		// fmt.Println(err)
		// if err != nil {
		// 	logger.Error("Error while deleting categories: ", err)
		// 	ctx.JSON(http.StatusInternalServerError, app_errors.NewHTTPErrorDTO(http.StatusInternalServerError, err, "Error Getting Shipping Address"))
		// 	return
		// } else if validateShippingAddress(shippingOut) == false {
		// 	//ctx.JSON(http.StatusBadRequest, app_errors.NewHTTPErrorDTO(http.StatusBadRequest, nil, "Shipping Address is not valid."))
		// 	return
		// }
	 _,err_ := th.categoryService.DeleteCategoryByID(category_id)
	if err_ != nil {
		ctx.Error(err_.Error())
		ctx.JSON(err_.Code, gin.H{"message": err_.Message})
		return
	}
	ctx.JSON(http.StatusOK,gin.H{"message": "Category deleted successfully"})
	}
}
func (th CategoryHandler) UpdateCategory() (gin.HandlerFunc) {
	return func(ctx *gin.Context){
	category_id := ctx.Param("category_id")
	var category *models.Category
      //  defer cancel()
	    // validate the request body
        if err := ctx.BindJSON(&category); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
	res, err := th.categoryService.UpdateCategoryByID(category_id,category)
	if err != nil {
		ctx.Error(err.Error())
		ctx.JSON(err.Code, gin.H{"message": err.Message})
		return
	}
	ctx.JSON(http.StatusOK,res)
}
}

