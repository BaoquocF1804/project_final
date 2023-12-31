package adapter

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"project_T4/service_user/user/pb_user"
)

type UserBankAdapter interface {
	GetUser(ctx context.Context, in *pb_user.GetUserRequest) (*pb_user.GetUserResponse, error)
}

type userBankAdapter struct {
	userBankClient pb_user.UserBankClient
}

func NewUserBankAdapter(addr string) UserBankAdapter {
	connectUser, err := grpc.Dial(
		addr,
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}

	clientUser := pb_user.NewUserBankClient(connectUser)

	return &userBankAdapter{
		userBankClient: clientUser,
	}
}

func (a *userBankAdapter) GetUser(ctx context.Context, in *pb_user.GetUserRequest) (*pb_user.GetUserResponse, error) {
	return a.userBankClient.GetUser(ctx, in)
}
