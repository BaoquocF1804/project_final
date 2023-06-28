package adapter

import (
	"context"
	"google.golang.org/grpc"
	"log"
	pb_user2 "project_T4/proto/user/pb_user"
)

type UserBankAdapter interface {
	GetUser(ctx context.Context, in *pb_user2.GetUserRequest) (*pb_user2.GetUserResponse, error)
}

type userBankAdapter struct {
	userBankClient pb_user2.UserBankClient
}

func NewUserBankAdapter(addr string) UserBankAdapter {
	connectUser, err := grpc.Dial(
		addr,
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}

	clientUser := pb_user2.NewUserBankClient(connectUser)

	return &userBankAdapter{
		userBankClient: clientUser,
	}
}

func (a *userBankAdapter) GetUser(ctx context.Context, in *pb_user2.GetUserRequest) (*pb_user2.GetUserResponse, error) {
	return a.userBankClient.GetUser(ctx, in)
}
