package controller

import (
	"github.com/Penetration-Platform-Go/Mysql-Service/model"
)

// InsertUser model into mysql
func InsertUser(user *User) {

}

// QueryUser model about mysql
func QueryUser(username string) *User {
	result := model.Query("username,nickname,password,email,photo", "users", "username='"+username+"'")
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
