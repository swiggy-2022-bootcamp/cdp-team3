package services

import (
	"context"

	auth "github.com/cdp-team3/categories-service/app/grpcs/auth"
//	"github.com/cdp-team3/categories-service/utils/logger"
	"google.golang.org/grpc"
)

//var logger = utils.NewLoggerService("auth-grpc")

func VerifyToken(token string) (*auth.VerifyTokenResponse, error) {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
	//	logger.Log("Failed to dial: %v", err)
		return nil, err
	}
	defer conn.Close()
	c := auth.NewAuthServiceClient(conn)
	r, err := c.VerifyToken(context.Background(), &auth.VerifyTokenRequest{
		Token: token,
	})
	if err != nil {
	//	logger("Failed to verify token: %v", err)
		return nil, err
	}
	return r, nil
}