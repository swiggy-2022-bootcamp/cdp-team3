package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/swiggy-2022-bootcamp/cdp-team3/payment-method-service/grpc"
	"github.com/swiggy-2022-bootcamp/cdp-team3/payment-method-service/utils"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

const (
	port = ":3080"
)

type Server struct {
	pb.UnimplementedAddPaymentMethodServiceServer
}

func (s *Server) AddPaymentMethod(ctx context.Context, in *pb.PaymentMethod) (*pb.PaymentMethod, error) {
	log := utils.InitializeLogger()

	zap.ReplaceGlobals(log)
	defer log.Sync()
	log.Info("Received add payment method request client")
	fmt.Println("Received add payment method request client")
	if in.Body != "COD" &&
		in.Body != "Net" &&
		in.Body != "CC" &&
		in.Body != "DC" &&
		in.Body != "UPI" {
		return &pb.PaymentMethod{Body: "False"}, nil
	} else {
		return &pb.PaymentMethod{Body: "True"}, nil
	}
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterAddPaymentMethodServiceServer(s, &Server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
