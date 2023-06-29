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
	"project_T4/service_user/server"
	"project_T4/service_user/user/pb_user"
)

const serverAddressUser = "0.0.0.0:8080"

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

	runGatewayServerUser(config_1, store)

}

func runGatewayServerUser(config config.Config, store *db.Queries) {
	server, err := server.NewSever(config, store)
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
