package api_connect

import (
	"fmt"
	"project_T4/adapter"
	db "project_T4/db/sqlc"
	"project_T4/pb_connect"
	"project_T4/token"
	"project_T4/util"
)

type Server struct {
	pb_connect.UnimplementedConnectBankServer
	config         util.Config
	store          *db.Queries
	tokenMaker     token.Maker
	accountAdapter adapter.AccountBankAdapter
	userAdapter    adapter.UserBankAdapter
}

func NewSever(config util.Config, store *db.Queries) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	accountAdapter := adapter.NewAccountBankAdapter("localhost:8082")
	userAdapter := adapter.NewUserBankAdapter("localhost:8080")
	server := &Server{
		config:         config,
		store:          store,
		tokenMaker:     tokenMaker,
		accountAdapter: accountAdapter,
		userAdapter:    userAdapter,
	}
	return server, nil
}
