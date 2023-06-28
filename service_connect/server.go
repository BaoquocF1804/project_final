package main

import (
	"context"
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"log"
	"net"
	"net/http"
	"project_T4/config"
	db "project_T4/db/sqlc"
	"project_T4/service_connect/connect/pb_connect"
	"project_T4/service_connect/server"
)

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

	runGatewayServerConnect(config_1, store)
}

func runGatewayServerConnect(config config.Config, store *db.Queries) {
	server2, err := server.NewSever(config, store)
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
