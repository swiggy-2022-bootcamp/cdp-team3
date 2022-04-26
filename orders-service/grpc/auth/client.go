package auth

import (
	"context"

	"github.com/swiggy-2022-bootcamp/cdp-team3/orders-service/configs"
	auth "github.com/swiggy-2022-bootcamp/cdp-team3/orders-service/grpc/auth/proto"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func VerifyToken(token string) (*auth.VerifyTokenResponse, error) {
	port := configs.EnvGrpcAuthClientPORT()
	serverHost := configs.EnvGrpcAuthClientHost()
	conn, err := grpc.Dial(serverHost+":"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		zap.L().Error("Failed to dial:" + port + " " + err.Error())
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
