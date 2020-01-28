package controller

import (
	"github.com/Penetration-Platform-Go/Mysql-Service/model"
)

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

// QueryUser model about mysql
func QueryUser(username string) *User {
	result := model.Query([]string{"username,nickname,password,email,photo"}, "users", "username='"+username+"'")
	if !result.IsValid {
		return &User{}
	}
	var user User
	for result.Value.Next() {
		var username, nickname, password, email, photo string
		result.Value.Scan(&username, &nickname, &password, &email, &photo)
		user.Username = username
		user.Nickname = nickname
		user.Password = password
		user.Email = email
		user.Photo = photo
	}
	return &user
}
