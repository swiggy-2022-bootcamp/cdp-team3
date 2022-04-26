package auth

import (
	"context"

	log "github.com/sirupsen/logrus"
	auth "github.com/swiggy-ipp/cart-service/grpcs/auth/proto"
	"google.golang.org/grpc"
)

func VerifyToken(token string) (*auth.VerifyTokenResponse, error) {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Errorf("Failed to dial: %v", err)
		return nil, err
	}
	defer conn.Close()
	c := auth.NewAuthServiceClient(conn)
	r, err := c.VerifyToken(context.Background(), &auth.VerifyTokenRequest{
		Token: token,
	})
	if err != nil {
		log.Errorf("Failed to verify token: %v", err)
		return nil, err
	}
	return r, nil
}
