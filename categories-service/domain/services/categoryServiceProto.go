package services

import (
	"context"
	"github.com/cdp-team3/categories-service/config"
	auth "github.com/cdp-team3/categories-service/app/grpcs/auth"

	"google.golang.org/grpc"
)



func VerifyToken(token string) (*auth.VerifyTokenResponse, error) {
	conn, err := grpc.Dial(config.EnvAuthHost() + ":" + config.EnvAuthServiceGRPCPort(), grpc.WithInsecure())

	if err != nil {
	
		return nil, err
	}
	defer conn.Close()
	c := auth.NewAuthServiceClient(conn)
	r, err := c.VerifyToken(context.Background(), &auth.VerifyTokenRequest{
		Token: token,
	})
	if err != nil {

		return nil, err
	}
	return r, nil
}