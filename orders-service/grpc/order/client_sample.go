package order

import (
	"context"
	"fmt"
	"log"

	ordergrpc "github.com/swiggy-2022-bootcamp/cdp-team3/orders-service/grpc/order/proto"
	"go.uber.org/zap"

	"google.golang.org/grpc"
)


func Client() {
	cc, err := grpc.Dial(":4004", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()

	c := ordergrpc.NewOrderServiceClient(cc)

	newOrder := &ordergrpc.RequestOrder{
		CustomerId: "1",
		TotalAmount: 200,
		ShippingAddressId: "34543654dfgdfnbdfx-dfgdsfb",
		OrderedProducts: []*ordergrpc.OrderedProduct{
			{
				ProductId: "2",
				Quantity: 5,
			},
			{
				ProductId: "3",
				Quantity: 50,
			},
		},
	}

	createOrderRes, err := c.CreateOrder(context.Background(), &ordergrpc.CreateOrderRequest{Order: newOrder})
	if err != nil {
		log.Fatalf("Unexpected error: %v", err)
	}
	fmt.Printf("Order has been created: %v", createOrderRes)
	zap.L().Info("Order Created"+createOrderRes.String())
}