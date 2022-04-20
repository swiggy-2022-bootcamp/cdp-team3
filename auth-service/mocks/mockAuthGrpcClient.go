package main

import (
	"context"
	"fmt"
	"log"

	auth "github.com/swiggy-2022-bootcamp/cdp-team3/auth-service/grpc/auth/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial: %v", err)
	}
	defer conn.Close()
	c := auth.NewAuthServiceClient(conn)
	r, err := c.VerifyToken(context.Background(), &auth.VerifyTokenRequest{
		Token: "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOiIxIiwibmFtZSI6IkpvaG4gRG9lIiwiZW1haWxJZCI6ImpvaG5AZ21haWwuY29tIiwiaXNBZG1pbiI6ZmFsc2UsImV4cCI6MTY1MDU0MzQwMX0.4U9JwOiLX1n1dthCSlfdObv9Y0La9ZlO0iW12n3H3qM",
	})
	if err != nil {
		log.Fatalf("Failed to verify token: %v", err)
	}
	fmt.Println(r)
}
