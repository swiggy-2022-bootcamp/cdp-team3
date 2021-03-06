package grpcs

import (
	context "context"
	"net"

	"github.com/swiggy-ipp/cart-service/dto/requests"
	"github.com/swiggy-ipp/cart-service/grpcs/cart_checkout"
	"github.com/swiggy-ipp/cart-service/services"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

// server is used to implement cart_checkout.CartCheckoutServiceServer
type server struct {
	cart_checkout.UnimplementedCartCheckoutServiceServer
	cartService services.CartService
}

// Procedure Implementation to Empty the Cart after Checkout
func (s *server) EmptyCart(
	ctx context.Context,
	in *cart_checkout.CartIDSignal,
) (*cart_checkout.CartEmptyOutput, error) {
	// Empty Cart in DB for the given user ID or cart ID
	err := s.cartService.EmptyCart(
		ctx,
		requests.CartIDRequest{UserID: in.UserID, CartID: in.CartID},
	)
	if err != nil {
		return &cart_checkout.CartEmptyOutput{Result: false}, err
	}
	log.Info("EmptyCart done for cart ID: ", in.CartID)
	return &cart_checkout.CartEmptyOutput{Result: true}, nil
}

// Procedure Implementation to Get the Cart Items using GRPC.
func (s *server) GetCartItems(
	ctx context.Context,
	in *cart_checkout.CartIDSignal,
) (*cart_checkout.CartItemsResponse, error) {
	// Empty Cart in DB for the given user ID or cart ID
	res, err := s.cartService.GetCartItems(
		ctx,
		in.CartID,
		in.UserID,
	)
	if err != nil {
		return nil, err
	}
	log.Info("GetCartItems done for cart ID: ", in.CartID)
	// Convert to GRPC Response
	cartItems := []*cart_checkout.CartItem{}
	for _, item := range res.CartItems {
		cartItems = append(cartItems, &cart_checkout.CartItem{
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Price:     float32(item.Price),
		})
	}
	return &cart_checkout.CartItemsResponse{
		Result:    true,
		CartItems: cartItems,
	}, nil
}

// Start GRPC Server using a net.Listener TCP Listener
func NewCartCheckoutGRPCServer(lis net.Listener, cartService services.CartService) error {
	// Create a gRPC server object
	s := grpc.NewServer()
	cart_checkout.RegisterCartCheckoutServiceServer(s, &server{cartService: cartService})
	log.Printf("Server listening at %v", lis.Addr())
	// Start serving requests
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
		return err
	}
	return nil
}
