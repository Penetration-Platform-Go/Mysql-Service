package grpc

import (
	"context"

	"github.com/Penetration-Platform-Go/Mysql-Service/controller"
	mysql "github.com/Penetration-Platform-Go/gRPC-Files/Mysql-Service"
)

// MysqlService define
type MysqlService struct {
}

// QueryUsers gRPC
func (u *MysqlService) QueryUsers(condition *mysql.Condition, stream mysql.Mysql_QueryUsersServer) error {
	result := controller.QueryUser(condition.Value)
	for _, user := range result {
		if err := stream.Send(&mysql.UserInformation{
			Username: user.Username,
			Nickname: user.Nickname,
			Password: user.Password,
			Email:    user.Email,
			Photo:    user.Photo,
		}); err != nil {
			return err
		}
	}
	return nil
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

// DeleteUser gRPC
func (u *MysqlService) DeleteUser(ctx context.Context, condition *mysql.Condition) (*mysql.Result, error) {
	return &mysql.Result{
		IsVaild: controller.DeleteUser(condition.Value),
	}, nil

}
