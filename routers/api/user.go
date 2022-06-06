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
	"webDesign/models"
	"webDesign/pkg/crypto"
	"webDesign/pkg/detection"
	"webDesign/pkg/e"
	"webDesign/pkg/util"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

type user struct {
	Username string `valid:"Required; MaxSize(50)" json:"username"`
	Password string `valid:"Required; MaxSize(50)" json:"password"`
}

// CheckUser 检查token是否过期
// @Summary      检查token是否过期
// @Description  根据header的token值检查token是否过期
// @Tags         checkUser
// @Accept       json
// @Produce      json
// @Success      200            {string}  string  "Success"
// @Router       /check [post]
func CheckUser(c *gin.Context) {
	token := c.GetHeader("token")
	code := e.SUCCESS
	if !util.CheckToken(token) {
		code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
	})
}

// GetUser 获取用户信息
// @Summary      登录并获取用户信息
// @Description  发送账密验证登录，返回用户信息和token
// @Tags         login
// @Accept       json
// @Produce      json
// @Param        token header string true "Token"
// @Success      200            {string}  string  "Success"
// @Router       /login [post]
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
		user, err := models.GetUserInfo(userId)
		// 获取ID不存在则返回参数错误
		if err != nil {
			code = e.ERROR_EXIST_USER
			c.JSON(http.StatusOK, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
			})
			return
		}
		data["id"] = user.ID
		data["head"] = "/static/img/" + user.Head + ".jpg"
		data["username"] = user.Username
		data["level"] = models.GetLevel(user.Level)
		err = models.UpdateVisitsNumber()
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

// RegisterUser 注册用户
// @Summary      根据用户信息注册普通用户
// @Description  发送姓名账密注册普通用户
// @Tags         register
// @Accept       json
// @Produce      json
// @Param        user body models.Message true "用户信息"
// @Success      200            {string}  string  "Success"
// @Router       /register [post]
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
