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
	Password string `json:"password" gorm:"not null;type:varchar(128)"`
	Level    int    `json:"level" gorm:"not null;type:tinyint"`
	Name     string `json:"name" gorm:"not null;type:varchar(20)"`
	Head     string `json:"head" gorm:"not null"`
}

const (
	SuperAdmin = 1
	Admin      = 2
	User       = 3
)

var LEVEL = map[int]string{
	SuperAdmin: "超级管理员",
	Admin:      "管理员",
	User:       "普通用户",
}

func AddUser(user *Message) (err error) {
	if err = db.Omit("id").Create(user).Error; err != nil {
		return err
	}
	return nil
}

func GetLevel(level int) string {
	return LEVEL[level]
}

func CheckUser(username, password string) uint {
	var message Message
	db.Select("id").Where(Message{Username: username, Password: password}).First(&message)
	return message.ID
}

func GetUser(username string) (msg Message) {
	db.Debug().Select("level").Where("username = ? ", username).First(&msg)
	return
}

func GetUserInfo(id uint) (msg Message) {
	db.Select("id,username,level,name,head").Where("id = ? ", id).First(&msg)
	return
}
func AddMessage(message Message) int {
	if err := db.Omit("ID").Create(&message).Error; err != nil {
		return e.ERROR_EXIST_USER
	}
	return e.SUCCESS
}
