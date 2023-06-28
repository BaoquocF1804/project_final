package adapter

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"project_T4/proto/account/pb_account"
)

type AccountBankAdapter interface {
	GetAccount(ctx context.Context, in *pb_account.GetAccountRequest) (*pb_account.GetAccountResponse, error)
}

type accountBankAdapter struct {
	accountBankClient pb_account.AccountBankClient
}

func NewAccountBankAdapter(addr string) *accountBankAdapter {
	connectAccount, err := grpc.Dial(
		addr,
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}

	clientAccount := pb_account.NewAccountBankClient(connectAccount)
	return &accountBankAdapter{
		accountBankClient: clientAccount,
	}
}

func (a *accountBankAdapter) GetAccount(ctx context.Context, in *pb_account.GetAccountRequest) (*pb_account.GetAccountResponse, error) {
	return a.accountBankClient.GetAccount(ctx, in)
}
