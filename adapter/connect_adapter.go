package adapter

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"project_T4/proto/connect/pb_connect"
)

type ConnectBankAdapter interface {
	GetAccount(ctx context.Context, in *pb_connect.GetAccountRequest) (*pb_connect.GetUserResponse, error)
}

type connectBankAdapter struct {
	connectBankClient pb_connect.ConnectBankClient
}

func NewConnectBankAdapter(addr string) *connectBankAdapter {
	connect, err := grpc.Dial(
		addr,
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}

	clientConnect := pb_connect.NewConnectBankClient(connect)
	return &connectBankAdapter{
		connectBankClient: clientConnect,
	}
}

func (a *connectBankAdapter) GetUserFromID(ctx context.Context, in *pb_connect.GetAccountRequest) (*pb_connect.GetUserResponse, error) {
	return a.connectBankClient.GetUserFromID(ctx, in)
}
