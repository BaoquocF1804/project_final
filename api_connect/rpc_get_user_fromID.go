package api_connect

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"project_T4/pb_account"
	"project_T4/pb_connect"
	"project_T4/pb_user"
)

func (server *Server) GetUserFromID(ctx context.Context, req *pb_connect.GetAccountRequest) (*pb_connect.GetUserResponse, error) {
	//count := req.GetID()
	accountReq := &pb_account.GetAccountRequest{
		ID: req.GetID(),
	}

	arg, err := server.accountAdapter.GetAccount(ctx, accountReq)
	//fmt.Println(err)
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
