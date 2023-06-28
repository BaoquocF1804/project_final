package api_user

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	db "project_T4/db/sqlc"
	"project_T4/proto/user/pb_user"
)

func convertUser(user db.User) *pb_user.User {
	return &pb_user.User{
		Username:          user.Username,
		FullName:          user.FullName,
		Email:             user.Email,
		PasswordChangedAt: timestamppb.New(user.PasswordChangedAt),
		CreatedAt:         timestamppb.New(user.CreatedAt),
	}
}
