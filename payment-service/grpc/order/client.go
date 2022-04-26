package order

import (
	"context"
	"fmt"
	"log"

	"github.com/swiggy-2022-bootcamp/cdp-team3/payment-service/configs"
	"github.com/swiggy-2022-bootcamp/cdp-team3/payment-service/errors"
	ordergrpc "github.com/swiggy-2022-bootcamp/cdp-team3/payment-service/grpc/order/proto"

	"google.golang.org/grpc"
)

func GetOrderStatus(orderId string) (string, *errors.AppError) {
	var addr = configs.EnvOrderHost() + ":" + configs.EnvOrderServiceGRPCPort()
	cc, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()

	c := ordergrpc.NewOrderServiceClient(cc)

	orderRes, err := c.GetOrder(context.Background(), &ordergrpc.GetOrderRequest{OrderId: orderId})
	if err != nil {
		log.Fatalf("Unexpected error: %v", err)
		return "", errors.NewUnexpectedError(err.Error())
	}
	fmt.Printf("Order has been created: %v", orderRes)
	return orderRes.Order.Status, nil
}
