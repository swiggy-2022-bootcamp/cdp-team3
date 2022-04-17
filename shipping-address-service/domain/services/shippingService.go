package services

import (
	"github.com/cdp-team3/shipping-address-service/domain/models"
	apperrors "github.com/cdp-team3/shipping-address-service/app-errors"
)

type ShippingService interface {
	
	InsertShippingAddress(shippingAddress *models.ShippingAddress) ( *apperrors.AppError)
	FindShippingAddressById(ShippingAddressID string) (*models.ShippingAddress,*apperrors.AppError)
	UpdateShippingAddressById(id string,shippingAddress *models.ShippingAddress) (bool, *apperrors.AppError)
	DeleteShippingAddressById(shippingAddressId string) (bool, *apperrors.AppError) 
	
}