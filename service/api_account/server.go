package api_account

import (
	"fmt"
	"project_T4/config"
	db "project_T4/db/sqlc"
	"project_T4/proto/account/pb_account"
	"project_T4/token"
)

type Server struct {
	pb_account.UnimplementedAccountBankServer
	config     config.Config
	store      *db.Queries
	tokenMaker token.Maker
}

// NewServer creates a new gRPC
func NewSever(config config.Config, store *db.Queries) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	return server, nil
}
