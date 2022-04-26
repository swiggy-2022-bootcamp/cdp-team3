package mode_of_payment_service

import (
	"fmt"
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func Main() {
	fmt.Println("CAME IN")
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := NewAddPaymentMethodServiceClient(conn)

	response, err := c.AddPaymentMethod(context.Background(), &PaymentMethod{Body: "Hello From Client!"})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from server: %s", response.Body)

}

/*

	lis, err := net.Listen("tcp", ":3007")
	if err != nil {
		log.Info("failed to listen to port 3007")
	}

	grpcServer := grpc.NewServer()

	mode_of_payment_service.RegisterAddPaymentMethodServiceServer(grpcServer, mode_of_payment_service.UnimplementedAddPaymentMethodServiceServer{})

	if err := grpcServer.Serve(lis); err != nil {
		log.Info("failed to serve on port 3007")
	}
*/
