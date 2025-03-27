package main

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
	"link_shorten_server/user/dao"
	user "link_shorten_server/user/kitex_gen/user"
	"link_shorten_server/user/response"
	utils2 "link_shorten_server/utils"
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
	//2.1.检查性别是否合法
	if !(req.Gender == "male" || req.Gender == "female") {
		return &user.UserRegisterResponse{Status: &response.WrongGender}, nil
	}
	//2.2.检查邮箱是否唯一
	result, status = dao.IfEmailExists(req.Email)
	if status != emptyStatus {
		return &user.UserRegisterResponse{Status: &status}, nil
	}
	if result == true {
		return &user.UserRegisterResponse{Status: &response.EmailExists}, nil
	}
	//2.3.检查手机号是否唯一
	result, status = dao.IfPhoneNumberExists(req.PhoneNumber)
	if status != emptyStatus {
		return &user.UserRegisterResponse{Status: &status}, nil
	}
	if result == true {
		return &user.UserRegisterResponse{Status: &response.PhoneNumberExists}, nil
	}
	//3.插入新用户信息
	//3.1.加密密码
	hashedPwd, err := utils2.HashPassword(req.Password)
	if err != nil {
		retErr := response.InternalErr(err)
		return &user.UserRegisterResponse{Status: &retErr}, nil
	}
	req.Password = hashedPwd
	//3.2.插入用户信息
	status = dao.InsertUserInfo(*req)
	if status != emptyStatus {
		return &user.UserRegisterResponse{Status: &status}, nil
	}
	return &user.UserRegisterResponse{Status: &response.Ok}, nil
}

// UserLogin implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserLogin(ctx context.Context, req *user.UserLoginRequest) (resp *user.UserLoginResponse, err error) {
	var emptyStatus user.Status
	//1.先寻找用户是否存在
	result, status := dao.IfUsernameExists(req.Username)
	if status != emptyStatus {
		return &user.UserLoginResponse{Status: &status}, nil
	}
	if result == false {
		return &user.UserLoginResponse{Status: &response.WrongUsrName}, nil
	}
	//2.接下来检验密码
	//2.1.获取数据库中存储的密码
	hashedPwd, status := dao.GetUserHashedPassword(req.Username)
	if status != emptyStatus {
		return &user.UserLoginResponse{Status: &status}, nil
	}
	//2.2.检验密码
	result, err = utils2.CompareHashPwdAndPwd(hashedPwd, req.Password)
	if err != nil {
		retErr := response.InternalErr(err)
		return &user.UserLoginResponse{Status: &retErr}, nil
	}
	if result == false {
		return &user.UserLoginResponse{Status: &response.WrongPassword}, nil
	}
	//3.登录成功，返回token
	//3.1.先获取用户id
	userID, status := dao.GetUserIDByName(req.Username)
	if status != emptyStatus {
		return &user.UserLoginResponse{Status: &response.WrongPassword}, nil
	}
	//3.2.生成token
	accessToken, refreshToken, err := utils2.GenerateTokens(int(userID))
	if err != nil {
		retErr := response.InternalErr(err)
		return &user.UserLoginResponse{Status: &retErr}, nil
	}
	var respData user.UserLoginResponseData
	respData.AccessToken = &accessToken
	respData.RefreshToken = &refreshToken
	//3.3.返回结果
	return &user.UserLoginResponse{Status: &response.Ok, Data: &respData}, nil
}

// TokenRefresh implements the UserServiceImpl interface.
func (s *UserServiceImpl) TokenRefresh(ctx context.Context, req *user.TokenRefreshRequest) (resp *user.TokenRefreshResponse, err error) {
	var emptyStatus user.Status
	//1.验证refreshToken
	token, status := utils2.ValidateRefreshToken(req.RefreshToken)
	if token == nil || status != emptyStatus {
		return &user.TokenRefreshResponse{Status: &response.InvalidToken}, nil
	}
	//2.生成新的token
	accessToken, refreshToken, err := utils2.GenerateTokens(int(token.Claims.(jwt.MapClaims)["user_id"].(float64)))
	if err != nil {
		retErr := response.InternalErr(err)
		return &user.TokenRefreshResponse{Status: &retErr}, nil
	}
	//3.返回结果
	var respData user.TokenRefreshResponseData
	respData.AccessToken = accessToken
	respData.RefreshToken = refreshToken
	return &user.TokenRefreshResponse{Status: &response.Ok, Data: &respData}, nil
}
