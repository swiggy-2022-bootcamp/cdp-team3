package grpcs

import (
	context "context"
	"net"

	"github.com/swiggy-ipp/cart-service/grpcs/cart_checkout"

	log "github.com/sirupsen/logrus"
	grpc "google.golang.org/grpc"
)

type server struct {
	cart_checkout.UnimplementedCheckoutServiceServer
}

func (s *server) EmptyCart(
	ctx context.Context, 
	in *cart_checkout.CartEmptySignal,
) (*cart_checkout.CartEmptyOutput, error) {
	// TODO: Empty Cart in DB for the given cart ID *in.CartID
	log.Info("EmptyCart called with cart ID: ", in.CartID)
	return &cart_checkout.CartEmptyOutput{Result: true}, nil
}

func StartGRPCServer(lis net.Listener) error {
	// Create a gRPC server object
	s := grpc.NewServer()
	cart_checkout.RegisterCheckoutServiceServer(s, &server{})
	log.Printf("Server listening at %v", lis.Addr())
	// Start serving requests
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
		return err
	}
	return nil
}
