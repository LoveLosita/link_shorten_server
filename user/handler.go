package main

import (
	"context"
	"link_shorten_server/user/dao"
	user "link_shorten_server/user/kitex_gen/user"
	"link_shorten_server/user/response"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// UserRegister implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserRegister(ctx context.Context, req *user.UserRegisterRequest) (resp *user.UserRegisterResponse, err error) {
	var emptyStatus user.Status
	//1.检查用户名是否已经存在
	result, status := dao.IfUsernameExists(req.Username)
	if status != emptyStatus {
		return &user.UserRegisterResponse{Status: &status}, nil
	}
	if result == true { //已经存在
		return &user.UserRegisterResponse{Status: &response.UsernameExists}, nil
	}
	//2.检查用户参数是否合法
	if !(req.Gender == "male" || req.Gender == "female") {
		return &user.UserRegisterResponse{Status: &response.WrongGender}, nil
	}
	//3.插入新用户信息
	status = dao.InsertUserInfo(*req)
	if status != emptyStatus {
		return &user.UserRegisterResponse{Status: &status}, nil
	}
	return &user.UserRegisterResponse{}, nil
}

// UserLogin implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserLogin(ctx context.Context, req *user.UserLoginRequest) (resp *user.UserLoginResponse, err error) {
	// TODO: Your code here...
	return
}

// TokenRefresh implements the UserServiceImpl interface.
func (s *UserServiceImpl) TokenRefresh(ctx context.Context, req *user.TokenRefreshRequest) (resp *user.TokenRefreshResponse, err error) {
	// TODO: Your code here...
	return
}
