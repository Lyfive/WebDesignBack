/*
*

	@author: 李映飞
	@date:   2022/4/28
	@note:

*
*/
package models

import (
	"fmt"
	"log"
	setting "webDesign/pkg"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	var (
		err                          error
		dbName, user, password, host string
	)
	sec, err := setting.Cfg.GetSection("database")
	if err != nil {
		log.Fatal(2, "Fail to get section 'database': %v", err)
	}

	dbName = sec.Key("NAME").String()
	user = sec.Key("USER").String()
	password = sec.Key("PASSWORD").String()
	host = sec.Key("HOST").String()
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName)

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Printf("Fail to open 'database' :%v in 'models/models.go' ", err)
	}
	InitModels()
}

func InitModels() {
	err := db.AutoMigrate(&System{})
	if err == nil {
		AddSystemMessage()
	}
	db.AutoMigrate(&Message{})
	db.AutoMigrate(&Faculty{})
	db.AutoMigrate(&Department{})
	db.AutoMigrate(&Session{})
	db.AutoMigrate(&Class{})
	db.AutoMigrate(&Student{})
	db.AutoMigrate(&Course{})
	db.AutoMigrate(&Grade{})
	db.AutoMigrate(&DC{})
}
