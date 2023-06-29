package server

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	db "project_T4/db/sqlc"
	"project_T4/service_account/account/pb_account"
)

func (server *Server) CreateAccount(ctx context.Context, req *pb_account.CreateAccountRequest) (*pb_account.CreateAccountResponse, error) {

	arg := db.CreateAccountParams{
		Owner:    req.GetOwner(),
		Balance:  req.GetBalance(),
		Currency: req.GetCurrency(),
	}

	err := server.store.CreateAccount(ctx, arg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create account: %s", err)
	}
	account := pb_account.Account{
		Owner:    arg.Owner,
		Balance:  arg.Balance,
		Currency: arg.Currency,
	}
	rps := &pb_account.CreateAccountResponse{
		Account: &account,
	}
	return rps, nil
}
