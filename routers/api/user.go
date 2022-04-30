/**
 @author: 李映飞
 @date:   2022/4/28
 @note:
**/
package api

import (
	"log"
	"net/http"
	"webDesign/models"
	"webDesign/pkg/e"
	"webDesign/pkg/util"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

type user struct {
	Username string `valid:"Required; MaxSize(50)" json:"username"`
	Password string `valid:"Required; MaxSize(50)" json:"password"`
}

func GetUser(c *gin.Context) {
	valid := validation.Validation{}
	var a user
	c.BindJSON(&a)
	ok, _ := valid.Valid(&a)

	code := e.INVALID_PARAMS
	if ok {
		isExist := models.CheckUser(a.Username, a.Password)
		if isExist {
			token, err := util.GenerateToken(a.Username, a.Username)
			if err != nil {
				code = e.ERROR_AUTH_TOKEN
			} else {
				c.Writer.Header().Add("token",token)
				code = e.SUCCESS
			}

		} else {
			code = e.ERROR_AUTH
		}
	} else {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
	})
}
