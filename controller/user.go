package controller

import (
	"github.com/Penetration-Platform-Go/Mysql-Service/model"
	pb "github.com/Penetration-Platform-Go/gRPC-Files/Mysql-Service"
)

// QueryUser model about mysql
func QueryUser(con []*pb.Value) []User {
	condition := "1 = 1"
	for _, each := range con {
		condition += " and " + each.Key + " = '" + each.Value + "'"
	}
	result := model.Query([]string{"username,nickname,password,email,photo"}, "users", condition)
	if !result.IsValid {
		return []User{}
	}
	var users []User
	for result.Value.Next() {
		var user User
		var username, nickname, password, email, photo string
		result.Value.Scan(&username, &nickname, &password, &email, &photo)
		user.Username = username
		user.Nickname = nickname
		user.Password = password
		user.Email = email
		user.Photo = photo
		users = append(users, user)
	}
	return users
}

// InsertUser model into mysql
func InsertUser(user *User) bool {
	column := []string{"username", "password", "nickname", "email", "photo"}
	value := []string{user.Username, user.Password, user.Nickname, user.Email, user.Photo}
	return model.Insert("users", column, value)
}

// UpdateUser model about mysql
func UpdateUser(user *User) bool {
	column := []string{"password", "nickname", "email", "photo"}
	value := []string{user.Password, user.Nickname, user.Email, user.Photo}
	return model.Update("users", "username='"+user.Username+"'", column, value)
}

// DeleteUser model about mysql
func DeleteUser(con []*pb.Value) bool {
	condition := "1 = 1"
	for _, each := range con {
		condition += " and " + each.Key + " = '" + each.Value + "'"
	}
	return model.Delete("users", condition)
}
