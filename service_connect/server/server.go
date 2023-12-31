package server

import (
	"fmt"
	"project_T4/config"
	db "project_T4/db/sqlc"
	adapter2 "project_T4/service_connect/adapter"
	"project_T4/service_connect/connect/pb_connect"
	"project_T4/token"
)

type Server struct {
	pb_connect.UnimplementedConnectBankServer
	config         config.Config
	store          *db.Queries
	tokenMaker     token.Maker
	accountAdapter adapter2.AccountBankAdapter
	userAdapter    adapter2.UserBankAdapter
}

func NewSever(config config.Config, store *db.Queries) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	accountAdapter := adapter2.NewAccountBankAdapter("localhost:8082")
	userAdapter := adapter2.NewUserBankAdapter("localhost:8080")

	server := &Server{
		config:         config,
		store:          store,
		tokenMaker:     tokenMaker,
		accountAdapter: accountAdapter,
		userAdapter:    userAdapter,
	}

	return server, nil
}
