package server

import (
	"fmt"
	"project_T4/config"
	db "project_T4/db/sqlc"
	"project_T4/service_user/user/pb_user"
	"project_T4/token"
)

type Server struct {
	pb_user.UnimplementedUserBankServer
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
