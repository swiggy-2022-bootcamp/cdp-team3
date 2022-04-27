package services

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/cdp-team3/shipping-address-service/domain/models"
	"github.com/cdp-team3/shipping-address-service/domain/repository"
	"github.com/cdp-team3/shipping-address-service/app/grpcs/shipping"
)

var shippingRepository repository.ShippingRepository

type ShippingProtoServer struct {
	shipping.UnimplementedShippingServer
}

func NewShippingProtoService(sr repository.ShippingRepository) ShippingProtoServer {
	shippingRepository = sr
	return ShippingProtoServer{}
}

// Returns shipping address to checkout service when grpc call is made
func (s ShippingProtoServer) GetShippingAddressForCheckout(ctx context.Context, ShippingRequest *shipping.ShippingAddressRequestFromCheckout) (*shipping.ShippingAddressResponseForCheckout, error) {
	id :=ShippingRequest.UserID;

	res, err := shippingRepository.GetDefaultShippingAddressOfUserByIdFromDB(id)
	if err != nil {
		return &shipping.ShippingAddressResponseForCheckout{},  err.Error()
	}
	return &shipping.ShippingAddressResponseForCheckout{
		ShippingAddressID : res.Id,
		Firstname: res.FirstName,
		Lastname:  res.LastName,
		City:      res.City,
		Address1:  res.Address1,
		Address2:  res.Address2,
		Countryid: uint32(res.CountryID),
		Postcode:  uint32(res.PostCode),
	
	}, nil

}

// returns shipping address to user and admin service when grpc call is made with shipping address ID
func (s ShippingProtoServer) GetShippingAddress(ctx context.Context, shippingRequest *shipping.ShippingAddressGetRequest) (*shipping.ShippingAddressGetResponse, error) {
	id := shippingRequest.ShippingAddressID
	
	res, err := shippingRepository.FindShippingAddressByIdFromDB(id)
	if err != nil {
		return &shipping.ShippingAddressGetResponse{},  err.Error()
	}
	return &shipping.ShippingAddressGetResponse{
		ShippingAddressID : res.Id,
		Firstname: res.FirstName,
		Lastname:  res.LastName,
		City:      res.City,
		Address1:  res.Address1,
		Address2:  res.Address2,
		Countryid: uint32(res.CountryID),
		Postcode:  uint32(res.PostCode),
		Userid   : res.UserID,
		Default : res.DefaultAddress,
	
	}, nil

}

// stores new shipping address in shipping address DB when a grpc call is made frm admin or customer service while creating a new user and the returns shipping address Id after storing address in Shipping DB
func (s ShippingProtoServer) AddShippingAddress(ctx context.Context,  shippingRequest *shipping.ShippingAddressAddRequest) (*shipping.ShippingAddressAddResponse,  error) {

	
	shippingAddress := &models.ShippingAddress{
		Id:        uuid.New().String(),
		FirstName: shippingRequest.Firstname,
		LastName:  shippingRequest.Lastname,
		City:      shippingRequest.City,
		Address1:  shippingRequest.Address1,
		Address2:  shippingRequest.Address2,
		CountryID: uint32(shippingRequest.Countryid),
		PostCode:  uint32(shippingRequest.Postcode),
		UserID:    shippingRequest.Userid,
		DefaultAddress: shippingRequest.Default,

	}
	res, err := shippingRepository.InsertShippingAddressToDB(shippingAddress)
	if err != nil {
		return &shipping.ShippingAddressAddResponse{}, err.Error()
	}
	return &shipping.ShippingAddressAddResponse{
		ShippingAddressID: res,
	}, nil
}

// deletes shipping address when admin or customer deletes his profile
func (s ShippingProtoServer) DeleteShippingAddress(ctx context.Context, shippingRequest *shipping.ShippingAddressDeleteRequest) (*shipping.ShippingAddressDeleteResponse,  error) {

	res, err := shippingRepository.DeleteShippingAddressByIdFromDB(shippingRequest.ShippingAddressID)
	if err != nil {
		return &shipping.ShippingAddressDeleteResponse{Confirm: false},err.Error()
	}
	return &shipping.ShippingAddressDeleteResponse{Confirm: res}, nil
}

// updates shipping address when admin or customer updates his profile
func (s ShippingProtoServer) UpdateShippingAddress(ctx context.Context, shippingRequest *shipping.ShippingAddressUpdateRequest) (*shipping.ShippingAddressUpdateResponse,  error) {


	newDaModel := &models.ShippingAddress{
		FirstName: shippingRequest.Firstname,
		LastName:  shippingRequest.Lastname,
		City:      shippingRequest.City,
		Address1:  shippingRequest.Address1,
		Address2:  shippingRequest.Address2,
		PostCode:  uint32(shippingRequest.Postcode),
		CountryID: uint32(shippingRequest.Countryid),
		UserID:    shippingRequest.Userid,
		DefaultAddress: shippingRequest.Default,
	}
	res, err := shippingRepository.UpdateShippingAddressByIdFromDB(shippingRequest.ShippingAddressID, newDaModel)
	if err != nil {
		return &shipping.ShippingAddressUpdateResponse{Confirm: false}, err.Error()
	}
	return &shipping.ShippingAddressUpdateResponse{Confirm: res}, nil
}