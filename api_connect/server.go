package api_connect

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	db "project_T4/db/sqlc"
	"project_T4/pb_account"
	"project_T4/pb_connect"
	"project_T4/pb_user"
	"project_T4/token"
	"project_T4/util"
)

type Server struct {
	pb_connect.UnimplementedConnectBankServer
	config        util.Config
	store         *db.Queries
	tokenMaker    token.Maker
	accountClient pb_account.AccountBankClient
	userClient    pb_user.UserBankClient
}

func NewSever(config util.Config, store *db.Queries) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	conn, err := grpc.Dial(
		"localhost:8082",
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	//defer conn.Close()
	conn1, err := grpc.Dial(
		"localhost:8080",
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	//defer conn.Close()
	accountClient := pb_account.NewAccountBankClient(conn)
	//fmt.Print(accountClient)
	userClient := pb_user.NewUserBankClient(conn1)
	server := &Server{
		config:        config,
		store:         store,
		tokenMaker:    tokenMaker,
		accountClient: accountClient,
		userClient:    userClient,
	}
	return server, nil
}
