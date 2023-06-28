package adapter

import (
	"context"
	"google.golang.org/grpc"
	"log"
	pb_connect2 "project_T4/service_connect/connect/pb_connect"
)

type ConnectBankAdapter interface {
	GetAccount(ctx context.Context, in *pb_connect2.GetAccountRequest) (*pb_connect2.GetUserResponse, error)
}

type connectBankAdapter struct {
	connectBankClient pb_connect2.ConnectBankClient
}

func NewConnectBankAdapter(addr string) *connectBankAdapter {
	connect, err := grpc.Dial(
		addr,
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}

	clientConnect := pb_connect2.NewConnectBankClient(connect)
	return &connectBankAdapter{
		connectBankClient: clientConnect,
	}
}

func (a *connectBankAdapter) GetUserFromID(ctx context.Context, in *pb_connect2.GetAccountRequest) (*pb_connect2.GetUserResponse, error) {
	return a.connectBankClient.GetUserFromID(ctx, in)
}
