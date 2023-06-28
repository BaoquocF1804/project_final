package adapter

import (
	"context"
	"google.golang.org/grpc"
	"log"
	pb_account2 "project_T4/service_account/account/pb_account"
)

type AccountBankAdapter interface {
	GetAccount(ctx context.Context, in *pb_account2.GetAccountRequest) (*pb_account2.GetAccountResponse, error)
}

type accountBankAdapter struct {
	accountBankClient pb_account2.AccountBankClient
}

func NewAccountBankAdapter(addr string) *accountBankAdapter {
	connectAccount, err := grpc.Dial(
		addr,
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}

	clientAccount := pb_account2.NewAccountBankClient(connectAccount)
	return &accountBankAdapter{
		accountBankClient: clientAccount,
	}
}

func (a *accountBankAdapter) GetAccount(ctx context.Context, in *pb_account2.GetAccountRequest) (*pb_account2.GetAccountResponse, error) {
	return a.accountBankClient.GetAccount(ctx, in)
}
