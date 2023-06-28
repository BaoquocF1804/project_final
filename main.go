package main

import (
	"context"
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"net/http"
	"project_T4/config"
	db "project_T4/db/sqlc"
	"project_T4/proto/account/pb_account"
	"project_T4/proto/connect/pb_connect"
	"project_T4/proto/user/pb_user"
	"project_T4/service/api_account"
	"project_T4/service/api_connect"
	"project_T4/service/api_user"
)

var conn *sql.DB

const serverAddressUser = "0.0.0.0:8080"
const serverAddressAccount = "0.0.0.0:8082"
const serverAddressConnect = "0.0.0.0:8084"

func main() {
	config_1, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config")
	}

	cfg := mysql.Config{
		User:                 ("root"),
		Passwd:               ("secret"),
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "simplebank",
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	// Get a database handle.
	conn, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	store := db.New(conn)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}

	go runGatewayServerAccount(config_1, store)
	go runGatewayServerUser(config_1, store)
	runGatewayServerConnect(config_1, store)
}

func runGatewayServerUser(config config.Config, store *db.Queries) {
	server, err := api_user.NewSever(config, store)
	if err != nil {
		log.Fatal("cannot create server: ", err)
	}

	grpcServer := grpc.NewServer()
	pb_user.RegisterUserBankServer(grpcServer, server)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", serverAddressUser)
	if err != nil {
		log.Fatal("cannot create listener: ", err)
	}

	log.Printf("start gRPC %s", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot start gRPC: ", err)
	}
}

func runGatewayServerAccount(config config.Config, store *db.Queries) {
	server, err := api_account.NewSever(config, store)
	if err != nil {
		log.Fatal("cannot create server: ", err)
	}

	grpcServer := grpc.NewServer()
	pb_account.RegisterAccountBankServer(grpcServer, server)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", serverAddressAccount)
	if err != nil {
		log.Fatal("cannot create listener: ", err)
	}

	log.Printf("start gRPC %s", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot start gRPC: ", err)
	}
}
func runGatewayServerConnect(config config.Config, store *db.Queries) {
	server2, err := api_connect.NewSever(config, store)
	if err != nil {
		log.Fatal("cannot create server: ", err)
	}

	grpcMux := runtime.NewServeMux()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err = pb_connect.RegisterConnectBankHandlerServer(ctx, grpcMux, server2)
	if err != nil {
		log.Fatal("cannot register handler sever", err)
	}

	mux := http.NewServeMux()
	mux.Handle("/", grpcMux)

	listener, err := net.Listen("tcp", serverAddressConnect)
	if err != nil {
		log.Fatal("cannot create listener: ", err)
	}

	log.Printf("start Connect 2API gateway server at %s", listener.Addr().String())
	err = http.Serve(listener, mux)
	if err != nil {
		log.Fatal("cannot start Connect 2API gateway server: ", err)
	}
}
