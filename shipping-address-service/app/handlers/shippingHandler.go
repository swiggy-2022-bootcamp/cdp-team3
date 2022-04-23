package handlers

import (
	"fmt"
//	"time"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"github.com/google/uuid"
	apperros "github.com/cdp-team3/shipping-address-service/app-errors"
	"github.com/cdp-team3/shipping-address-service/domain/models"
	"github.com/cdp-team3/shipping-address-service/domain/services"
)

var validate = validator.New()

type ShippingHandler struct {
	shippingService services.ShippingService
}

func NewShippingHandler(shippingService services.ShippingService) ShippingHandler {
	return ShippingHandler{shippingService: shippingService}
}
func toPersistedDynamodbEntitySA(o *models.ShippingAddress) *models.ShippingAddress {
	return &models.ShippingAddress{
		Id:        uuid.New().String(),
		FirstName: o.FirstName,
		LastName:  o.LastName,
		City:      o.City,
		Address1:  o.Address1,
		Address2:  o.Address2,
		CountryID: o.CountryID,
		PostCode:  o.PostCode,
	}
}

// Create Shipping Address
// @Summary      Create Shipping Address
// @Description  This Handler allow user to create new Shipping Address
// @Tags         Shipping Address
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {number} 	http.StatusBadRequest
// @Router       /shippingaddress    [post]
func (th ShippingHandler) AddNewShippingAddress() gin.HandlerFunc {
	return func(ctx *gin.Context) {
	//userId := c.Param("userId")
	var shippingAddress *models.ShippingAddress

	if err := ctx.BindJSON(&shippingAddress); err != nil {
		ctx.Error(err)
		err := apperros.NewBadRequestError(err.Error())
		ctx.JSON(err.Code, gin.H{"message": err.Message})
		return
	}
    fmt.Println(shippingAddress)
//use the validator library to validate required fields
if validationErr := validate.Struct(&shippingAddress); validationErr != nil {
	ctx.JSON(http.StatusBadRequest, gin.H{"message": validationErr.Error()})
	return
}
	shippingAddressRecord := toPersistedDynamodbEntitySA(shippingAddress)
	
	err := th.shippingService.InsertShippingAddress(shippingAddressRecord)
		if err != nil {
		ctx.Error(err.Error())
		ctx.JSON(err.Code, gin.H{"message": err.Message})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Shipping Address added successfully"})
}
}

// Get Shipping Address by Id
// @Summary      Get Shipping Address by id
// @Description  This Handle returns shippingAddress given id
// @Tags         Shipping Address
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {number} 	http.StatusBadRequest
// @Router       /shippingaddress/:id    [get]
func (th ShippingHandler) GetShippingAddress() gin.HandlerFunc  {
	return func(ctx *gin.Context) {
	Id := ctx.Param("id")
	fmt.Println("Inside handler",Id)
	res,err := th.shippingService.FindShippingAddressById(Id)
		if err != nil {
		ctx.Error(err.Error())
		ctx.JSON(err.Code, gin.H{"message": err.Message})
		return
	}
	ctx.JSON(http.StatusOK,res)
}
}


// Update Shipping Address
// @Summary      Update Shipping Address
// @Description  This Handle Update shippingAddress given id
// @Tags         Shipping Address
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {number} 	http.StatusBadRequest
// @Router       /shippingaddress/:id     [put]
func (sh ShippingHandler) HandleUpdateShippingAddressByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var shippingAddress ShippingAddressRecordDTO
		id := ctx.Param("id")
		if err := ctx.BindJSON(&shippingAddress); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		newshipAddr := convertShippingAddressDTOtoShippingAddressModel(shippingAddress)
		fmt.Println("new",newshipAddr)
		fmt.Println("old",shippingAddress)

		ok, err := sh.shippingService.UpdateShippingAddressById(id,newshipAddr)
		if !ok {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		ctx.JSON(http.StatusAccepted, gin.H{"message": "Shipping Address record updated"})
	}
}

// Delete Shipping Address
// @Summary      Delete Shipping Address
// @Description  This Handle deletes Delete Shipping Address given sid
// @Tags         Shipping Address
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {number} 	http.StatusBadRequest
// @Router       /shippingaddress/:id   [delete]
func (sh ShippingHandler) HandleDeleteShippingAddressById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		fmt.Println(id)
		_, err := sh.shippingService.DeleteShippingAddressById(id)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		ctx.JSON(http.StatusAccepted, gin.H{"message": "Deleted Succesfully"})
	}
}


type ShippingAddressRecordDTO struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	City      string `json:"city"`
	Address1  string `json:"address_1"`
	Address2  string `json:"address_2"`
	CountryID int    `json:"country_id"`
	PostCode  int    `json:"postcode"`
}
func convertShippingAddressDTOtoShippingAddressModel(saDto ShippingAddressRecordDTO) *models.ShippingAddress {

	return &models.ShippingAddress{
		FirstName: saDto.FirstName,
		LastName:  saDto.LastName,
		City:      saDto.City,
		Address1:  saDto.Address1,
		Address2:  saDto.Address2,
		PostCode:  saDto.PostCode,
		CountryID: saDto.CountryID,
	}
}