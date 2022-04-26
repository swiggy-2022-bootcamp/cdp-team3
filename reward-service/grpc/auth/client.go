package auth

import (
	"context"

	"github.com/swiggy-2022-bootcamp/cdp-team3/rewards-service/configs"
	auth "github.com/swiggy-2022-bootcamp/cdp-team3/rewards-service/grpc/auth/proto"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func VerifyToken(token string) (*auth.VerifyTokenResponse, error) {
	conn, err := grpc.Dial(configs.EnvAuthHost()+":"+configs.EnvGrpcAuthClientPORT(), grpc.WithInsecure())
	if err != nil {
		zap.L().Error("Failed to dial:" + configs.EnvGrpcAuthClientPORT() + " " + err.Error())
		return nil, err
	}
	defer conn.Close()
	c := auth.NewAuthServiceClient(conn)
	r, err := c.VerifyToken(context.Background(), &auth.VerifyTokenRequest{
		Token: token,
	})
	if err != nil {
		zap.L().Error("Failed to verify token:" + err.Error())
		return nil, err
	}
	return r, nil
}
