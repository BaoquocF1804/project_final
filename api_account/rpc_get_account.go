package api_account

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"project_T4/pb_account"
)

func (server *Server) GetAccount(ctx context.Context, req *pb_account.GetAccountRequest) (*pb_account.GetAccountResponse, error) {
	//fmt.Println(grpc.Version)

	arg, err := server.store.GetAccount(ctx, req.GetID())
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "no account")
	}

	rsp := &pb_account.GetAccountResponse{
		Account: convertUser(arg),
	}
	return rsp, nil
}
