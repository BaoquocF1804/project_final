package api_connect

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"log"
	db "project_T4/db/sqlc"
	"project_T4/pb_account"
	"project_T4/pb_connect"
	"project_T4/pb_user"
	"project_T4/token"
	"project_T4/util"
	"time"
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
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(1024*1024)), // Kích thước tối đa cho việc nhận dữ liệu
		grpc.WithDefaultCallOptions(grpc.MaxCallSendMsgSize(1024*1024)),
		grpc.WithKeepaliveParams(keepalive.ClientParameters{
			Time:                time.Second * 20, // Khoảng thời gian giữa các tin nhắn keepalive
			Timeout:             time.Second * 10, // Thời gian chờ keepalive trước khi coi kết nối bị mất
			PermitWithoutStream: true,             // Cho phép keepalive khi không có kênh truyền dữ liệu
		}),
	)
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	//defer conn.Close()
	conn1, err := grpc.Dial(
		"localhost:8080",
		grpc.WithInsecure(),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(1024*1024)), // Kích thước tối đa cho việc nhận dữ liệu
		grpc.WithDefaultCallOptions(grpc.MaxCallSendMsgSize(1024*1024)),
		grpc.WithKeepaliveParams(keepalive.ClientParameters{
			Time:                time.Second * 20, // Khoảng thời gian giữa các tin nhắn keepalive
			Timeout:             time.Second * 10, // Thời gian chờ keepalive trước khi coi kết nối bị mất
			PermitWithoutStream: true,             // Cho phép keepalive khi không có kênh truyền dữ liệu
		}),
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
