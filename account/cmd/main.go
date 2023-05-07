package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	accountpb "github.com/fiazbilal/accounts-microservice/accountpb"
)

type server struct{}

func (s *server) CreateAccount(ctx context.Context, req *accountpb.CreateAccountRequest) (*accountpb.CreateAccountResponse, error) {
	return &accountpb.CreateAccountResponse{}, nil
}

func (s *server) GetAccounts(req *accountpb.GetAccountsRequest, stream accountpb.AccountService_GetAccountsServer) error {
	return nil
}

func (s *server) GenerateAddress(ctx context.Context, req *accountpb.GenerateAddressRequest) (*accountpb.GenerateAddressResponse, error) {
	return &accountpb.GenerateAddressResponse{}, nil
}

func (s *server) Deposit(ctx context.Context, req *accountpb.DepositRequest) (*accountpb.DepositResponse, error) {
	return &accountpb.DepositResponse{}, nil
}

func (s *server) Withdraw(ctx context.Context, req *accountpb.WithdrawRequest) (*accountpb.WithdrawResponse, error) {
	return &accountpb.WithdrawResponse{}, nil
}

// func (s *server) mustEmbedUnimplementedAccountServiceServer() {}

func main() {
	fmt.Println("Starting AccountService...")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	accountpb.RegisterAccountServiceServer(s, &server{})

	fmt.Println("Starting ...")

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

	fmt.Println("Starting 76890-...")
}
