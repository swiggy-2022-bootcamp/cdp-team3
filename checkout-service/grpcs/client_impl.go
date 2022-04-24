package grpcs

import (
	"github.com/swiggy-ipp/checkout-service/grpcs/cart_checkout"
	order_checkout "github.com/swiggy-ipp/checkout-service/grpcs/order/proto"
	"github.com/swiggy-ipp/checkout-service/grpcs/shipping_checkout"

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
	OrderCheckoutGRPCChannel    chan order_checkout.OrderServiceClient       = make(chan order_checkout.OrderServiceClient)
)

/// Function with logic for becoming GRPC Client
func BecomeGRPCClient(cartPort string, shippingPort string, orderPort string) {
	// For Cart Checkout
	conn, err := grpc.Dial("0.0.0.0:"+cartPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
		ErrChanGRPC <- err
	} else {
		CartCheckoutGRPCChannel <- cart_checkout.NewCartCheckoutServiceClient(conn)
	}
	// For Shipping Checkout
	conn, err = grpc.Dial("0.0.0.0:"+shippingPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
		ErrChanGRPC <- err
	} else {
		ShippingCheckoutGRPCChannel <- shipping_checkout.NewShippingClient(conn)
	}
	// For Order Checkout
	conn, err = grpc.Dial("0.0.0.0:"+orderPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
		ErrChanGRPC <- err
	} else {
		OrderCheckoutGRPCChannel <- order_checkout.NewOrderServiceClient(conn)
	}
}
