/**
 @author: 李映飞
 @date:   2022/4/28
 @note:
**/
package models

import (
	"webDesign/pkg/e"
)

type Message struct {
	ID       uint   `gorm:"primary_key"`
	Username string `json:"username" gorm:"unique;not null;type:varchar(12)"`
	Password string `json:"password" gorm:"not null;type:varchar(20)"`
	Level    int   `json:"level" gorm:"not null;type:tinyint"`
	Name     string `json:"name" gorm:"not null;type:varchar(20)"`
	Head     string `json:"head" gorm:"not null"`
}

const (
	SuperAdmin = 1
	Admin      = 2
	User       = 3
)

func CheckUser(username, password string) bool {
	var message Message
	db.Select("id").Where(Message{Username: username, Password: password}).First(&message)
	if message.ID > 0 {
		return true
	}
	return false
}

func GetUser(username string) (msg Message) {
	db.Select("level").Where("username = ? ", username).First(&msg)
	return
}

func AddMessage(message Message) int {
	if err := db.Omit("ID").Create(&message).Error; err != nil {
		return e.ERROR_EXIST_USER
	}
	return e.SUCCESS
}
