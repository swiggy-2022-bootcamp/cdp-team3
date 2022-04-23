package grpcs

import (
	"github.com/swiggy-ipp/checkout-service/grpcs/cart_checkout"
	"github.com/swiggy-ipp/checkout-service/grpcs/shipping_checkout"
	order_checkout "github.com/swiggy-ipp/checkout-service/grpcs/order/proto"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	// Error Channels
	ErrChanGRPC chan error = make(chan error)

	// GRPC Client Channels
	CartCheckoutGRPCChannel     chan cart_checkout.CartCheckoutServiceClient = make(chan cart_checkout.CartCheckoutServiceClient)
	ShippingCheckoutGRPCChannel chan shipping_checkout.ShippingClient        = make(chan shipping_checkout.ShippingClient)
	OrderCheckoutGRPCChannel    chan order_checkout.OrderServiceClient = make(chan order_checkout.OrderServiceClient)
)

/// Function with logic for becoming GRPC Client
func BecomeGRPCClient(cartAddress string, shippingAddress string, orderAddress string) {
	// For Cart Checkout
	conn, err := grpc.Dial(cartAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
		ErrChanGRPC <- err
	} else {
		CartCheckoutGRPCChannel <- cart_checkout.NewCartCheckoutServiceClient(conn)
	}
	// For Shipping Checkout
	conn, err = grpc.Dial(shippingAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
		ErrChanGRPC <- err
	} else {
		ShippingCheckoutGRPCChannel <- shipping_checkout.NewShippingClient(conn)
	}
	// For Order Checkout
	conn, err = grpc.Dial(orderAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
		ErrChanGRPC <- err
	} else {
		OrderCheckoutGRPCChannel <- order_checkout.NewOrderServiceClient(conn)
	}
}
