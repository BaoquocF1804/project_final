package adapter

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"project_T4/pb_account"
)

type AccountBankAdapter interface {
	GetAccount(ctx context.Context, in *pb_account.GetAccountRequest) (*pb_account.GetAccountResponse, error)
}

type accountBankAdapter struct {
	accountBankClient pb_account.AccountBankClient
}

func NewAccountBankAdapter(addr string) AccountBankAdapter {
	conn1, err := grpc.Dial(
		addr,
		grpc.WithInsecure(),
	)
	fmt.Println(addr)
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	clientAccount := pb_account.NewAccountBankClient(conn1)
	return &accountBankAdapter{
		accountBankClient: clientAccount,
	}
}

func (a *accountBankAdapter) GetAccount(ctx context.Context, in *pb_account.GetAccountRequest) (*pb_account.GetAccountResponse, error) {
	return a.accountBankClient.GetAccount(ctx, in)

}
