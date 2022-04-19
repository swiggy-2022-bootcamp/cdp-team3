package services

import (
	"context"
	"fmt"
	//app_errors "github.com/cdp-team3/shipping-address-service/app-errors"
	"github.com/cdp-team3/shipping-address-service/domain/repository"
	"github.com/cdp-team3/shipping-address-service/app/protobuf"
)

var shippingRepository repository.ShippingRepository

type ShippingProtoServer struct {
	protobuf.UnimplementedShippingServer
}

func NewShippingProtoService(sr repository.ShippingRepository) ShippingProtoServer {
	shippingRepository = sr
	return ShippingProtoServer{}
}

func (s ShippingProtoServer) GetShippingAddress(ctx context.Context, shippingRequest *protobuf.ShippingAddressRequest) (*protobuf.ShippingAddressResponse, error) {
	id := shippingRequest.ShippingAddressID
	fmt.Println("Id in grpc",id)
	res, err := shippingRepository.FindShippingAddressByIdFromDB(id)
	if err != nil {
		return &protobuf.ShippingAddressResponse{},  err.Error()
	}
	return &protobuf.ShippingAddressResponse{
		Firstname: res.FirstName,
		Lastname:  res.LastName,
		City:      res.City,
		Address1:  res.Address1,
		Address2:  res.Address2,
		Countryid: uint32(res.CountryID),
		Postcode:  uint32(res.PostCode),
	}, nil



}