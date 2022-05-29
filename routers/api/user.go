/**
 @author: 李映飞
 @date:   2022/4/28
 @note:
**/
package api

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"webDesign/middleware/crypto"
	"webDesign/middleware/detection"
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
	// md5加密
	a.Password = crypto.Encrypt(a.Password)
	ok, _ := valid.Valid(&a)

	code := e.INVALID_PARAMS
	var userId uint
	if ok {
		userId = models.CheckUser(a.Username, a.Password)
		if userId > 0 {
			token, err := util.GenerateToken(a.Username, a.Username)
			if err != nil {
				code = e.ERROR_AUTH_TOKEN
			} else {
				c.Writer.Header().Add("token", token)
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

	data := make(map[string]interface{})
	if userId > 0 {
		user := models.GetUserInfo(userId)
		data["head"] = user.Head
		data["username"] = user.Username
		data["level"] = models.GetLevel(user.Level)
		err := models.UpdateVisitsNumber()
		if err != nil {
			log.Println(err)
		}
		system, _ := models.GetSystemMessage()
		fmt.Println(system)
		data["createTime"] = system.CreateTime.Format("2006-01-02 15:04:05")
		data["visitsNumber"] = system.VisitsNumber
		data["version"] = system.Version
		data["population"] = models.GetPopulationMessage()
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

// 限定字符串为a-z A-Z 0-9 _ %
func CheckString(str string) bool {
	re := regexp.MustCompile("^[a-zA-Z0-9_]+$")
	return re.MatchString(str)
}

func RegisterUser(c *gin.Context) {
	var user models.Message
	c.BindJSON(&user)
	// 暂时没上传头像这个功能
	fmt.Println(user)
	code := e.SUCCESS
	// 检查输入信息是否合法
	if CheckString(user.Password) && detection.CheckSensitiveWord(user.Username) && detection.CheckSensitiveWord(user.Name) {
		user.Password = crypto.Encrypt(user.Password)
		user.Level = models.User
		user.Head = "https://github.com/Lyfive/MyPictures/blob/master/head/go1.png?raw=true"
		err := models.AddUser(&user)
		if err != nil {
			code = e.ERROR_EXIST_USER
		}
	} else {
		code = e.CONTAINING_SENSITIVE_WORD
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
	})
}
