package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/swiggy-2022-bootcamp/cdp-team3/payment-method-service/grpc"
	"google.golang.org/grpc"
)

const (
	address = "localhost:3080"
)

func main() {

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewAddPaymentMethodServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	selectedMode := "COD"

	r, err := c.AddPaymentMethod(ctx, &pb.PaymentMethod{Body: selectedMode})
	if err != nil {
		log.Fatalf("could not create user: %v", err)
	}
	fmt.Println(r.Body)

}
