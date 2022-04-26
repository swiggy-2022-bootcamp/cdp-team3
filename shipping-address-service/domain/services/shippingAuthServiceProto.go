package services

import (
	"context"
	"github.com/cdp-team3/shipping-address-service/config"
	auth "github.com/cdp-team3/shipping-address-service/app/grpcs/auth"
//	"github.com/cdp-team3/categories-service/utils/logger"
	"google.golang.org/grpc"
)

//var logger = utils.NewLoggerService("auth-grpc")

func VerifyToken(token string) (*auth.VerifyTokenResponse, error) {
	conn, err := grpc.Dial(config.EnvAuthHost() + ":" + config.EnvAuthServiceGRPCPort(), grpc.WithInsecure())
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