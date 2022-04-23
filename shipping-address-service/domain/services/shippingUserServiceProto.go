package services

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	//app_errors "github.com/cdp-team3/shipping-address-service/app-errors"
	"github.com/cdp-team3/shipping-address-service/domain/models"
	"github.com/cdp-team3/shipping-address-service/domain/repository"
	"github.com/cdp-team3/shipping-address-service/app/grpcs/shipping_user"
)

var shippingRepository repository.ShippingRepository

type ShippingProtoServer struct {
	shipping_user.UnimplementedShippingServer
}

func NewShippingProtoService(sr repository.ShippingRepository) ShippingProtoServer {
	shippingRepository = sr
	return ShippingProtoServer{}
}

func (s ShippingProtoServer) GetShippingAddress(ctx context.Context, shippingRequest *shipping_user.ShippingAddressRequest) (*shipping_user.ShippingAddressResponse, error) {
	id := shippingRequest.ShippingAddressID
	fmt.Println("Id in grpc",id)
	res, err := shippingRepository.FindShippingAddressByIdFromDB(id)
	if err != nil {
		return &shipping_user.ShippingAddressResponse{},  err.Error()
	}
	return &shipping_user.ShippingAddressResponse{
		Firstname: res.FirstName,
		Lastname:  res.LastName,
		City:      res.City,
		Address1:  res.Address1,
		Address2:  res.Address2,
		Countryid: int(res.CountryID),
		Postcode:  int(res.PostCode),
	}, nil

}
func (s ShippingProtoServer) AddShippingAddress(ctx context.Context,  shippingRequest *shipping_user.ShippingAddressAddRequest) (*shipping_user.ShippingAddressAddResponse,  error) {
	fmt.Println("Add in grpc",shippingRequest)
	shippingAddress := models.ShippingAddress{
		Id:        uuid.New().String(),
		FirstName: shippingRequest.Firstname,
		LastName:  shippingRequest.Lastname,
		City:      shippingRequest.City,
		Address1:  shippingRequest.Address1,
		Address2:  shippingRequest.Address2,
		CountryID: int(shippingRequest.Countryid),
		PostCode:  int(shippingRequest.Postcode),
	}
	res, err := shippingRepository.InsertShippingAddressToDB(shippingAddress)
	if err != nil {
		return &shipping_user.ShippingAddressAddResponse{}, err.Error()
	}
	return &shipping_user.ShippingAddressAddResponse{
		ShippingAddressID: res,
	}, nil
}

func (s ShippingProtoServer) DeleteShippingAddress(ctx context.Context, shippingRequest *shipping_user.ShippingAddressRequest) (*shipping_user.ShippingAddressDeleteResponse,  error) {
	fmt.Println("Delete in grpc",shippingRequest)
	res, err := shippingRepository.DeleteShippingAddressByIdFromDB(shippingRequest.ShippingAddressID)
	if err != nil {
		return &shipping_user.ShippingAddressDeleteResponse{Confirm: false},err.Error()
	}
	return &shipping_user.ShippingAddressDeleteResponse{Confirm: res}, nil
}
func (s ShippingProtoServer) UpdateShippingAddress(ctx context.Context, shippingRequest *shipping_user.ShippingAddressUpdateRequest) (*shipping_user.ShippingAddressUpdateResponse,  error) {
	fmt.Println("Delete in grpc",shippingRequest)
	newDaModel := models.ShippingAddress{
		FirstName: shippingRequest.Firstname,
		LastName:  shippingRequest.Lastname,
		City:      shippingRequest.City,
		Address1:  shippingRequest.Address1,
		Address2:  shippingRequest.Address2,
		PostCode:  int(shippingRequest.Postcode),
		CountryID: int(shippingRequest.Countryid),
	}
	res, err := shippingRepository.UpdateShippingAddressByIdFromDB(shippingRequest.ShippingAddressID, newDaModel)
	if err != nil {
		return &shipping_user.ShippingAddressUpdateResponse{Confirm: false}, err.Error()
	}
	return &shipping_user.ShippingAddressUpdateResponse{Confirm: res}, nil
}