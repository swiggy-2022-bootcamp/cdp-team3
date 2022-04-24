package auth

import (
	"context"
	"fmt"
	"log"
	"net"
	"strings"

	auth "github.com/swiggy-2022-bootcamp/cdp-team3/auth-service/grpc/auth/proto"
	"github.com/swiggy-2022-bootcamp/cdp-team3/auth-service/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthServer struct {
	auth.AuthServiceServer
}

func (s *AuthServer) VerifyToken(ctx context.Context, req *auth.VerifyTokenRequest) (*auth.VerifyTokenResponse, error) {
	if req.Token == "" {
		return nil, status.Error(codes.InvalidArgument, "Token is empty")
	}
	token := strings.TrimPrefix(req.Token, "Bearer ")
	isValid, _ := utils.ValidateToken(token)
	if isValid != true {
		return nil, status.Error(codes.Unauthenticated, "Invalid token")
	}
	claims, err := utils.GetClaimsFromToken(token)
	if err != nil {
		fmt.Println("error", err)
		return nil, status.Error(codes.Unauthenticated, "Invalid token")
	}

	return &auth.VerifyTokenResponse{
		UserId:  claims.UserId,
		Name:    claims.Name,
		Email:   claims.EmailId,
		IsAdmin: claims.IsAdmin,
	}, nil
}

func InitialiseAuthServer() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	auth.RegisterAuthServiceServer(s, &AuthServer{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
