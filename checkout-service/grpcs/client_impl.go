package grpcs

import (
	"github.com/swiggy-ipp/checkout-service/grpcs/cart_checkout"
	"github.com/swiggy-ipp/checkout-service/grpcs/shipping_checkout"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	// Error Channels
	ErrChanGRPC chan error = make(chan error)

	// GRPC Client Channels
	CartCheckoutGRPCChannel chan cart_checkout.CartCheckoutServiceClient = make(chan cart_checkout.CartCheckoutServiceClient)
	ShippingCheckoutGRPCChannel chan shipping_checkout.ShippingClient = make(chan shipping_checkout.ShippingClient)
)

/// Function with logic for becoming GRPC Client
func BecomeGRPCClient(cartAddress string, shippingAddress string) {
	// Create a listener on TCP port
	conn, err := grpc.Dial(cartAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
		ErrChanGRPC <- err
	} else {
		// Start GRPC
		CartCheckoutGRPCChannel <- cart_checkout.NewCartCheckoutServiceClient(conn)
	}

	conn, err = grpc.Dial(shippingAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
		ErrChanGRPC <- err
	} else {
		// Start GRPC
		ShippingCheckoutGRPCChannel <- shipping_checkout.NewShippingClient(conn)
	}
}

