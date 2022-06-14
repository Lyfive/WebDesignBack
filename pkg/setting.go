/**
 @author: 李映飞
 @date:   2022/3/13
 @note:
**/
package setting

import (
	"github.com/go-ini/ini"
	"log"
)

var (
	Cfg *ini.File

	RunMode string

	HOST     string
	HTTPHost string
	HTTPPort int

	PageSize  int
	JwtSecret string
)

func init() {
	var err error
	Cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("Fail to 'conf/app.ini' : %v", err)
	}

	LoadBase()
	LoadServer()
	LoadApp()
}

func LoadBase() {
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}

func LoadServer() {
	sec, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatalf("Fail to get section 'server': %v", err)
	}
	HOST = sec.Key("HOST").MustString("http://127.0.0.1:8080")
	HTTPHost = sec.Key("HTTP_HOST").MustString("127.0.0.1")
	HTTPPort = sec.Key("HTTP_PORT").MustInt(8080)
}

func LoadApp() {
	sec, err := Cfg.GetSection("app")
	if err != nil {
		log.Fatalf("Fail to get section 'app': %v", err)
	}

	JwtSecret = sec.Key("JWT_SECRET").MustString("!@)*#)!@U#@*!@!)")
	PageSize = sec.Key("PAGE_SIZE").MustInt(10)
}
