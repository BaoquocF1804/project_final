package api_user

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"project_T4/pb_user"
)

func (server *Server) GetUser(ctx context.Context, req *pb_user.GetUserRequest) (*pb_user.GetUserResponse, error) {
	arg, err := server.store.GetUser(ctx, req.GetUsername())
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "no account")
	}
	rsp := &pb_user.GetUserResponse{
		User: convertUser(arg),
	}
	fmt.Println(grpc.Version)
	return rsp, nil
}
