package server

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	db "project_T4/db/sqlc"
	"project_T4/service_account/account/pb_account"
)

func convertUser(account db.Account) *pb_account.Account {
	return &pb_account.Account{
		ID:       account.ID,
		Owner:    account.Owner,
		Balance:  int32(account.Balance),
		Currency: account.Currency,
		CreatAt:  timestamppb.New(account.CreatedAt),
	}
}
