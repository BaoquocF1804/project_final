package server

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"project_T4/service_account/account/pb_account"
	"project_T4/service_connect/connect/pb_connect"
	"project_T4/service_user/user/pb_user"
)

func (server *Server) GetUserFromID(ctx context.Context, req *pb_connect.GetAccountRequest) (*pb_connect.GetUserResponse, error) {
	accountReq := &pb_account.GetAccountRequest{
		ID: req.GetID(),
	}
	arg, err := server.accountAdapter.GetAccount(ctx, accountReq)
	if err != nil {
		log.Fatalf("GetAccount RPC failed: %v", err)
	}

	accountReq2 := &pb_user.GetUserRequest{
		Username: arg.Account.Owner,
	}
	arg1, err := server.userAdapter.GetUser(ctx, accountReq2)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "no user")
	}

	userRes := &pb_connect.GetUserResponse{
		Username: arg1.User.Username,
		FullName: arg1.User.FullName,
		Email:    arg1.User.Email,
	}
	return userRes, nil
}
