package services

import (
	"context"
	"fmt"
	//app_errors "github.com/cdp-team3/shipping-address-service/app-errors"
	"github.com/cdp-team3/shipping-address-service/domain/repository"
	"github.com/cdp-team3/shipping-address-service/app/grpcs/shipping_checkout"
)

var shippingRepository repository.ShippingRepository

type ShippingProtoServer struct {
	shipping_checkout.UnimplementedShippingServer
}

func NewShippingProtoService(sr repository.ShippingRepository) ShippingProtoServer {
	shippingRepository = sr
	return ShippingProtoServer{}
}

func (s ShippingProtoServer) GetShippingAddress(ctx context.Context, shippingRequest *shipping_checkout.ShippingAddressRequest) (*shipping_checkout.ShippingAddressResponse, error) {
	id := shippingRequest.ShippingAddressID
	fmt.Println("Id in grpc",id)
	res, err := shippingRepository.FindShippingAddressByIdFromDB(id)
	if err != nil {
		return &shipping_checkout.ShippingAddressResponse{},  err.Error()
	}
	return &shipping_checkout.ShippingAddressResponse{
		Firstname: res.FirstName,
		Lastname:  res.LastName,
		City:      res.City,
		Address1:  res.Address1,
		Address2:  res.Address2,
		Countryid: uint32(res.CountryID),
		Postcode:  uint32(res.PostCode),
	}, nil



}