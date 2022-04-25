package grpcs

import (
	"fmt"
	"github.com/cdp-team3/categories-service/config"
	protos "github.com/cdp-team3/categories-service/app/grpcs/products"
	"google.golang.org/grpc"
)

// Function to get the gRPC client object of products service
// Dialing to products service without any security as Dialup option
func GetProductsGrpcClient() (protos.ProductsClient, error) {
	conn, err :=  grpc.Dial(config.EnvProductsHost() + ":" + config.EnvProductsServiceGRPCPort(), grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("unable to connect with grpc server")
	}

	client := protos.NewProductsClient(conn)
	return client, nil
}
