package api_connect

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"project_T4/mocks"
	_ "project_T4/mocks"
	"project_T4/pb_account"
	"project_T4/pb_connect"
	"project_T4/pb_user"
	"testing"
)

func TestGetUserFromID(t *testing.T) {
	// Tạo một mocks server instance
	mockServer := &Server{
		accountAdapter: &mocks.AccountBankAdapter{},
		userAdapter:    &mocks.UserBankAdapter{},
	}

	// Tạo một GetAccountRequest mẫu
	req := &pb_connect.GetAccountRequest{
		ID: 2,
	}

	// Thiết lập phản hồi của mocks adapter cho GetAccount
	mockAccountResponse := &pb_account.GetAccountResponse{
		Account: &pb_account.Account{
			Owner: "mocked_username",
		},
	}
	mockServer.accountAdapter.(*mocks.AccountBankAdapter).On("GetAccount", context.Background(), mock.AnythingOfType("*pb_account.GetAccountRequest")).Return(mockAccountResponse, nil)

	// Thiết lập phản hồi của mocks adapter cho GetUser
	mockUserResponse := &pb_user.GetUserResponse{
		User: &pb_user.User{
			Username: "mocked_username",
			FullName: "Mocked User",
			Email:    "mocked_email@example.com",
		},
	}
	mockServer.userAdapter.(*mocks.UserBankAdapter).On("GetUser", context.Background(), mock.AnythingOfType("*pb_user.GetUserRequest")).Return(mockUserResponse, nil)

	// Gọi phương thức GetUserFromID
	res, err := mockServer.GetUserFromID(context.Background(), req)

	// Xác nhận rằng cuộc gọi phương thức thành công
	assert.NoError(t, err)

	// Xác nhận các giá trị mong đợi trong phản hồi
	expectedRes := &pb_connect.GetUserResponse{
		Username: "mocked_username",
		FullName: "Mocked User",
		Email:    "mocked_email@example.com",
	}
	assert.Equal(t, expectedRes, res)

	// Xác minh rằng phương thức GetAccount của adapter đã được gọi với đúng các tham số
	mockServer.accountAdapter.(*mocks.AccountBankAdapter).AssertCalled(t, "GetAccount", context.Background(), mock.AnythingOfType("*pb_account.GetAccountRequest"))

	// Xác minh rằng phương thức GetUser của adapter đã được gọi với đúng các tham số
	mockServer.userAdapter.(*mocks.UserBankAdapter).AssertCalled(t, "GetUser", context.Background(), mock.AnythingOfType("*pb_user.GetUserRequest"))
}
