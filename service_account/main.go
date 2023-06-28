package main

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"project_T4/config"
	db "project_T4/db/sqlc"
	"project_T4/service_account/account/pb_account"
	"project_T4/service_account/server"
)

var conn *sql.DB

const serverAddressAccount = "0.0.0.0:8082"

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

	runGatewayServerAccount(config_1, store)

}

func runGatewayServerAccount(config config.Config, store *db.Queries) {
	server, err := server.NewSever(config, store)
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
