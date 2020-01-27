package grpc

import (
	"context"

	"github.com/Penetration-Platform-Go/Mysql-Service/controller"
	mysql "github.com/Penetration-Platform-Go/gRPC-Files/Mysql-Service"
)

// MysqlService define
type MysqlService struct {
}

//QueryUserByUsername gRPC
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

func (u *MysqlService) InsertUser(ctx context.Context, username *mysql.UserInformation) (*mysql.Result, error) {
	return &mysql.Result{}, nil
}

func (u *MysqlService) UpdateUser(ctx context.Context, username *mysql.UserInformation) (*mysql.Result, error) {
	return &mysql.Result{}, nil
}
