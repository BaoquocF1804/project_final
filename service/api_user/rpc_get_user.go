package api_user

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"project_T4/proto/user/pb_user"
)

func (server *Server) GetUser(ctx context.Context, req *pb_user.GetUserRequest) (*pb_user.GetUserResponse, error) {
	arg, err := server.store.GetUser(ctx, req.GetUsername())
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "no account")
	}

	rsp := &pb_user.GetUserResponse{
		User: convertUser(arg),
	}
	return rsp, nil
}
