package auth

import (
	"context"
	"log"

	"github.com/swiggy-2022-bootcamp/cdp-team3/admin-service/configs"
	auth "github.com/swiggy-2022-bootcamp/cdp-team3/admin-service/grpc/auth/proto"
	"google.golang.org/grpc"
)

func VerifyToken(token string) (*auth.VerifyTokenResponse, error) {
	conn, err := grpc.Dial(
		configs.EnvAuthHost()+":"+configs.EnvGrpcAuthClientPORT(),
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalf("Failed to dial: %v", err)
		return nil, err
	}
	defer conn.Close()
	c := auth.NewAuthServiceClient(conn)
	r, err := c.VerifyToken(context.Background(), &auth.VerifyTokenRequest{
		Token: token,
	})
	if err != nil {
		log.Fatalf("Failed to verify token: %v", err)
		return nil, err
	}
	return r, nil
}
