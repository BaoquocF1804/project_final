package server

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	db "project_T4/db/sqlc"
	"project_T4/service_user/user/pb_user"
	"project_T4/util"
)

func (server *Server) CreateUser(ctx context.Context, req *pb_user.CreateUserRequest) (*pb_user.CreateUserResponse, error) {
	hashedPassword, err := util.HashPassword(req.GetPassword())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to hash password: %s", err)
	}

	arg := db.CreateUserParams{
		Username:       req.GetUsername(),
		HashedPassword: hashedPassword,
		FullName:       req.GetFullName(),
		Email:          req.GetEmail(),
	}

	err = server.store.CreateUser(ctx, arg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create use: %s", err)
	}

	user, err := server.store.GetUser(ctx, arg.Username)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create user: %s", err)
	}

	rsp := &pb_user.CreateUserResponse{
		User: convertUser(user),
	}
	return rsp, nil
}
