package repository

import (
	"github.com/cdp-team3/shipping-address-service/domain/models"
	apperrors "github.com/cdp-team3/shipping-address-service/app-errors"
	
)
type ShippingRepository interface {
	
	 DBHealthCheck() bool
	 InsertShippingAddressToDB(shippingAddress *models.ShippingAddress) ( *apperrors.AppError)
	 FindShippingAddressByIdFromDB(ShippingAddressID string) (*models.ShippingAddress,*apperrors.AppError)
	 UpdateShippingAddressByIdFromDB(id string,shippingAddress *models.ShippingAddress) (bool, *apperrors.AppError) 
	 DeleteShippingAddressByIdFromDB(id string) (bool, *apperrors.AppError) 
	
}