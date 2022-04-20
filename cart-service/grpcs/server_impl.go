package grpcs

import (
	context "context"
	"net"

	"github.com/swiggy-ipp/cart-service/grpcs/cart_checkout"
	"github.com/swiggy-ipp/cart-service/repositories"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

// server is used to implement cart_checkout.CartCheckoutServiceServer
type server struct {
	cart_checkout.UnimplementedCartCheckoutServiceServer
	cartRepository repositories.CartRepository
}

// Procedure Implementation to Empty the Cart after Checkout
func (s *server) EmptyCart(
	ctx context.Context, 
	in *cart_checkout.CartEmptySignal,
) (*cart_checkout.CartEmptyOutput, error) {
	// Empty Cart in DB for the given cart ID *in.CartID
	err := s.cartRepository.EmptyCart(ctx, in.CartID)
	if err != nil {
		return &cart_checkout.CartEmptyOutput{Result: false}, err
	}
	log.Info("EmptyCart done for cart ID: ", in.CartID)
	return &cart_checkout.CartEmptyOutput{Result: true}, nil
}

// Start GRPC Server using a net.Listener TCP Listener
func NewCartCheckoutGRPCServer(lis net.Listener, cartRepository repositories.CartRepository) error {
	// Create a gRPC server object
	s := grpc.NewServer()
	cart_checkout.RegisterCartCheckoutServiceServer(s, &server{cartRepository: cartRepository})
	log.Printf("Server listening at %v", lis.Addr())
	// Start serving requests
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
		return err
	}
	return nil
}
