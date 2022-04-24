package auth

import (
	"context"

	"github.com/swiggy-2022-bootcamp/cdp-team3/transaction-service/configs"
	auth "github.com/swiggy-2022-bootcamp/cdp-team3/transaction-service/grpc/auth/proto"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func VerifyToken(token string) (*auth.VerifyTokenResponse, error) {
	port := configs.EnvGrpcAuthClientPORT()
	conn, err := grpc.Dial(":"+port, grpc.WithInsecure())
	if err != nil {
		zap.L().Error("Failed to dial:"+port+" "+err.Error())
		return nil, err
	}
	defer conn.Close()
	c := auth.NewAuthServiceClient(conn)
	r, err := c.VerifyToken(context.Background(), &auth.VerifyTokenRequest{
		Token: token,
	})
	if err != nil {
		zap.L().Error("Failed to verify token:" +err.Error())
		return nil, err
	}
	return r, nil
}