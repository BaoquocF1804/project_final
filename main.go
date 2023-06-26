package main

import (
	"context"
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"log"
	"net"
	"net/http"
	"project_T4/api_account"
	"project_T4/api_connect"
	"project_T4/api_user"
	db "project_T4/db/sqlc"
	"project_T4/pb_account"
	"project_T4/pb_connect"
	"project_T4/pb_user"
	"project_T4/util"
)

var conn *sql.DB

const serverAddressUser = "0.0.0.0:8080"
const serverAddressAccount = "0.0.0.0:8082"
const serverAddressConnect = "0.0.0.0:8084"

func main() {
	config_1, err := util.LoadConfig(".")
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
	//server := api.NewSever(store)
	//err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
	go runGatewayServerAccount(config_1, store)
	go runGatewayServerUser(config_1, store)

	runGatewayServerConnect(config_1, store)
	//go runGatewayServer(config, store)

}

func runGatewayServerUser(config util.Config, store *db.Queries) {
	server, err := api_user.NewSever(config, store)
	if err != nil {
		log.Fatal("cannot create server: ", err)
	}
	//fmt.Println(grpc.Version)
	grpcMux := runtime.NewServeMux()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err = pb_user.RegisterUserBankHandlerServer(ctx, grpcMux, server)
	if err != nil {
		log.Fatal("cannot register handler sever", err)
	}

	mux := http.NewServeMux()
	mux.Handle("/", grpcMux)

	listener, err := net.Listen("tcp", serverAddressUser)
	if err != nil {
		log.Fatal("cannot create listener: ", err)
	}
	log.Printf("start API_user gateway server at %s", listener.Addr().String())
	err = http.Serve(listener, mux)
	if err != nil {
		log.Fatal("cannot start API_user gateway server: ", err)
	}
}

func runGatewayServerAccount(config util.Config, store *db.Queries) {
	server, err := api_account.NewSever(config, store)
	if err != nil {
		log.Fatal("cannot create server: ", err)
	}
	//fmt.Println(grpc.Version)
	grpcMux := runtime.NewServeMux()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err = pb_account.RegisterAccountBankHandlerServer(ctx, grpcMux, server)
	if err != nil {
		log.Fatal("cannot register handler sever", err)
	}

	mux := http.NewServeMux()
	mux.Handle("/", grpcMux)

	listener, err := net.Listen("tcp", serverAddressAccount)
	if err != nil {
		log.Fatal("cannot create listener: ", err)
	}
	log.Printf("start API_Account gateway server at %s", listener.Addr().String())
	err = http.Serve(listener, mux)
	if err != nil {
		log.Fatal("cannot start API_Account gateway server: ", err)
	}
}
func runGatewayServerConnect(config util.Config, store *db.Queries) {
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
