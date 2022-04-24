package services

import (
	"fmt"
	//"net/http"
	
	"github.com/cdp-team3/shipping-address-service/domain/models"
	apperros "github.com/cdp-team3/shipping-address-service/app-errors"
	"github.com/cdp-team3/shipping-address-service/domain/repository"
	"github.com/cdp-team3/shipping-address-service/utils/logger"
)

type ShippingServiceImpl struct {
	shippingRepository repository.ShippingRepository
}

func NewShippingServiceImpl(shippingRepository repository.ShippingRepository) ShippingService {
	return &ShippingServiceImpl{shippingRepository: shippingRepository}
}
func (p ShippingServiceImpl) InsertShippingAddress(shipping *models.ShippingAddress) (string,*apperros.AppError) {
fmt.Println("Inside service",shipping)
	
	 id,err := p.shippingRepository.InsertShippingAddressToDB(shipping)
	if err != nil {
		fmt.Println(err)
		logger.Error(err)
		return  "",err
	}
	return id,nil
}
func (p ShippingServiceImpl) FindShippingAddressById(id string) (*models.ShippingAddress,*apperros.AppError) {
fmt.Println("Inside service",id)
	
	 result,err := p.shippingRepository.FindShippingAddressByIdFromDB(id)
	if err != nil {
		fmt.Println(err)
		logger.Error(err)
		return  nil,err
	}
	return result,nil
}
func (p ShippingServiceImpl) UpdateShippingAddressById(shippingAddressId string, newShippingAddress *models.ShippingAddress) (bool, *apperros.AppError) {
	_, err := p.shippingRepository.UpdateShippingAddressByIdFromDB(shippingAddressId, newShippingAddress)
	if err != nil {
		return false, err
	}
	return true, nil
}
func (p ShippingServiceImpl) DeleteShippingAddressById(shippingAddressId string) (bool, *apperros.AppError) {
	_, err := p.shippingRepository.DeleteShippingAddressByIdFromDB(shippingAddressId)
	if err != nil {
		return false, err
	}
	return true, nil
}
func (p ShippingServiceImpl) HandleSetExistingShippingAddressToDefaultById(shippingAddressId string) (bool, *apperros.AppError) {
	_, err := p.shippingRepository.HandleSetExistingShippingAddressToDefaultByIdToDB(shippingAddressId)
	if err != nil {
		return false, err
	}
	return true, nil
}


func (p ShippingServiceImpl) GetDefaultShippingAddressOfUserById(userId string) (*models.ShippingAddress, *apperros.AppError) {
	res, err := p.shippingRepository.GetDefaultShippingAddressOfUserByIdFromDB(userId)
	if err != nil {
		return nil, err
	}
	return res, nil
}