package grpcs

import (
	"fmt"

	protos "github.com/cdp-team3/categories-service/app/grpcs/products"
	"google.golang.org/grpc"
)

// Function to get the gRPC client object of products service
// Dialing to products service without any security as Dialup option
func GetProductsGrpcClient() (protos.ProductsClient, error) {
	conn, err := grpc.Dial("localhost:2012", grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("unable to connect with grpc server")
	}

	client := protos.NewProductsClient(conn)
	return client, nil
}
