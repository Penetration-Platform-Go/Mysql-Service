package grpc

import (
	"context"

	"github.com/Penetration-Platform-Go/Mysql-Service/controller"
	mysql "github.com/Penetration-Platform-Go/gRPC-Files/Mysql-Service"
)

// MysqlService define
type MysqlService struct {
}

// QueryUserByUsername gRPC
func (u *MysqlService) QueryUserByUsername(ctx context.Context, username *mysql.Username) (*mysql.UserInformation, error) {
	r := controller.QueryUser(username.Username)
	return &mysql.UserInformation{
		Username: r.Username,
		Nickname: r.Nickname,
		Password: r.Password,
		Email:    r.Email,
		Photo:    r.Photo,
	}, nil
}

// InsertUser gRPC
func (u *MysqlService) InsertUser(ctx context.Context, user *mysql.UserInformation) (*mysql.Result, error) {
	return &mysql.Result{
		IsVaild: controller.InsertUser(&controller.User{
			Username: user.Username,
			Nickname: user.Nickname,
			Password: user.Password,
			Email:    user.Email,
			Photo:    user.Photo,
		}),
	}, nil
}

// UpdateUser gRPC
func (u *MysqlService) UpdateUser(ctx context.Context, user *mysql.UserInformation) (*mysql.Result, error) {
	return &mysql.Result{
		IsVaild: controller.UpdateUser(&controller.User{
			Username: user.Username,
			Nickname: user.Nickname,
			Password: user.Password,
			Email:    user.Email,
			Photo:    user.Photo,
		}),
	}, nil
}
