package grpcs

import (
	"github.com/swiggy-ipp/checkout-service/configs"
	"github.com/swiggy-ipp/checkout-service/grpcs/cart_checkout"
	order_checkout "github.com/swiggy-ipp/checkout-service/grpcs/order/proto"
	"github.com/swiggy-ipp/checkout-service/grpcs/shipping_checkout"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	// Error GRPC Channel
	ErrChanGRPC chan error = make(chan error)

	// GRPC Client Channels
	// Cart <-> Checkout GRPC Client Channel
	CartCheckoutGRPCChannel chan cart_checkout.CartCheckoutServiceClient = make(chan cart_checkout.CartCheckoutServiceClient)
	// Shipping <-> Checkout GRPC Client Channel
	ShippingCheckoutGRPCChannel chan shipping_checkout.ShippingClient = make(chan shipping_checkout.ShippingClient)
	// Order <-> Checkout GRPC Client Channel
	OrderCheckoutGRPCChannel chan order_checkout.OrderServiceClient = make(chan order_checkout.OrderServiceClient)
)

const didNotConnectErrorMessage string = "Did not connect: %v"

/// Function with logic for becoming GRPC Client
func BecomeGRPCClient() {
	// For Cart Checkout
	conn, err := grpc.Dial(
		configs.EnvCartHost() + ":" + configs.EnvCartServiceGRPCPort(), 
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf(didNotConnectErrorMessage, err)
		ErrChanGRPC <- err
	} else {
		CartCheckoutGRPCChannel <- cart_checkout.NewCartCheckoutServiceClient(conn)
	}
	// For Shipping Checkout
	conn, err = grpc.Dial(
		configs.EnvShippingHost() + ":" + configs.EnvShippingServiceGRPCPort(), 
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf(didNotConnectErrorMessage, err)
		ErrChanGRPC <- err
	} else {
		ShippingCheckoutGRPCChannel <- shipping_checkout.NewShippingClient(conn)
	}
	// For Order Checkout
	conn, err = grpc.Dial(
		configs.EnvOrderHost() + ":" + configs.EnvOrderServiceGRPCPort(), 
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf(didNotConnectErrorMessage, err)
		ErrChanGRPC <- err
	} else {
		OrderCheckoutGRPCChannel <- order_checkout.NewOrderServiceClient(conn)
	}
}
