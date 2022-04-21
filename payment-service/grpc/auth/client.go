package auth

import (
	"context"

	auth "github.com/swiggy-2022-bootcamp/cdp-team3/payment-service/grpc/auth/proto"
	"github.com/swiggy-2022-bootcamp/cdp-team3/payment-service/utils"
	"google.golang.org/grpc"
)

var logger = utils.NewLoggerService("auth-grpc")

func VerifyToken(token string) (*auth.VerifyTokenResponse, error) {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		logger.Log("Failed to dial: %v", err)
		return nil, err
	}
	defer conn.Close()
	c := auth.NewAuthServiceClient(conn)
	r, err := c.VerifyToken(context.Background(), &auth.VerifyTokenRequest{
		Token: token,
	})
	if err != nil {
		logger.Log("Failed to verify token: %v", err)
		return nil, err
	}
	return r, nil
}
