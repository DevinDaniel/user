package handler

import (
	"context"
	"user/domain/model"
	"user/domain/service"
	user "user/proto/user"
)

type User struct{
	UserDateService service.IUserDataService
}

//注册
func(u *User)Register(ctx context.Context,userRegisterRequest *user.UserRegisterRequest,userRegisterResponse *user.UserRegisterResponse) error{
	userRegister := &model.User{
		UserName:    userRegisterRequest.UserName ,
		//FirstName:   userRegisterRequest.FirstName,
		HashPassword:userRegisterRequest.Pwd,
	}
	_,err := u.UserDateService.AddUser(userRegister)
	if err!=nil{
		return err
	}
	userRegisterResponse.Message="添加成功"
	return nil
}

//登录
func(u *User)Login(ctx context.Context,userLoginRequest *user.UserLoginRequest,userLoginResponse *user.UserLoginResponse) error{
	isOK,err := u.UserDateService.CheckPwd(userLoginRequest.UserName,userLoginRequest.Pwd)
	if err!=nil{
		return err
	}
	userLoginResponse.IsSuccess=isOK
	return nil
}

//查询用户信息
func(u *User)GetUserInfo(ctx context.Context, userInfoRequest *user.UserInfoRequest,userInfoResponse *user.UserInfoResponse) error{
	userInfo,err := u.UserDateService.FindUserByName(userInfoRequest.UserName)
	if err!=nil{
		return err
	}
	userInfoResponse=UserForResponse(userInfo)
	return nil
}

//类型转化
func UserForResponse(userModel *model.User) *user.UserInfoResponse{
	response := &user.UserInfoResponse{}
	response.UserName=userModel.UserName
	response.UserId=userModel.ID
	return response
}