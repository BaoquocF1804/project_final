package api_user

import (
	"fmt"
	db "project_T4/db/sqlc"
	"project_T4/pb_user"
	"project_T4/token"
	"project_T4/util"
)

type Server struct {
	pb_user.UnimplementedUserBankServer
	config     util.Config
	store      *db.Queries
	tokenMaker token.Maker
}

// NewServer creates a new gRPC
func NewSever(config util.Config, store *db.Queries) (*Server, error) {
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
