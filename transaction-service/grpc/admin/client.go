package admin

import (
	"context"

	"github.com/swiggy-2022-bootcamp/cdp-team3/transaction-service/configs"
	admin "github.com/swiggy-2022-bootcamp/cdp-team3/transaction-service/grpc/admin/proto"
	"go.uber.org/zap"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func SendTransactionAmount(transaction *admin.TransactionDetails) (*admin.SuccessMessage, error) {
	port := configs.EnvGrpcAdminClientPORT()
	serverHost := configs.EnvGrpcAdminClientHost()
	conn, err := grpc.Dial(serverHost+":"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		zap.L().Error("Failed to dial:" + port + " " + err.Error())
		return nil, err
	}
	defer conn.Close()
	c := admin.NewTransactionAmountClient(conn)
	r, err := c.SendTransactionAmount(context.Background(), &admin.TransactionDetails{
		TransactionAmount: transaction.TransactionAmount,
		UserId:            transaction.UserId,
	})
	if err != nil {
		zap.L().Error("Failed to perform grpc call to update transaction points" + err.Error())
		return nil, err
	}
	return r, nil
}
