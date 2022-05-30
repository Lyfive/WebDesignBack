/**
 @author: 李映飞
 @date:   2022/4/28
 @note:
**/
package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"webDesign/middleware/jwt"
	"webDesign/models"
	"webDesign/pkg/e"
)

func Register(c *gin.Context) {
	var message models.Message

	// 拉取用户信息，通过token判断用户权限
	c.BindJSON(&message)
	token := c.GetHeader("token")
	level := jwt.GetUserLevel(token)
	code := e.SUCCESS

	// 头像上传比较麻烦 若头像为空则使用默认头像
	if message.Head == "" {
		message.Head = "https://github.com/Lyfive/MyPictures/blob/master/head/LyFiveHead.jpg?raw=true"
	}
	fmt.Println(message)
	if message.Level < level {
		code = e.ERROR_INSUFFICIENT_ACCESS_RIGHTS
		c.JSON(http.StatusUnauthorized, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
		})
		return
	}
	err := models.AddMessage(message)
	if err != nil {
		code = e.ERROR
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
	})
}

type UserID struct {
	ID uint `json:"id"`
}

func DeleteUser(c *gin.Context) {

	var id UserID
	// 拉取用户信息，通过token判断用户权限
	token := c.GetHeader("token")
	level := jwt.GetUserLevel(token)
	code := e.SUCCESS

	// 传入id无效，返回参数错误
	err := c.BindJSON(&id)
	if err != nil {
		fmt.Println(err)
		code = e.INVALID_PARAMS
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
		})
		return
	}

	// 删除用户时，用户不存在 或 用户权限大于等于自己时返回错误
	message, err := models.GetUserInfo(id.ID)
	if err != nil || message.Level <= level {
		code = e.ERROR_INSUFFICIENT_ACCESS_RIGHTS
		c.JSON(http.StatusUnauthorized, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
		})
		return
	}

	// 删除用户 保证待输出用户存在
	err = models.DeleteUser(message.ID)
	if err != nil {
		code = e.ERROR_EXIST_USER
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
	})
}

func GetUserList(c *gin.Context) {

	// 拉取用户信息，通过token判断用户权限
	token := c.GetHeader("token")
	level := jwt.GetUserLevel(token)
	code := e.SUCCESS

	// 获取等级小于当前用户的用户列表
	users := models.GetUsers(level)

	c.JSON(http.StatusOK, gin.H{
		"code":  code,
		"msg":   e.GetMsg(code),
		"users": users,
	})
}

func ModifyUser(c *gin.Context) {

	// 拉取用户信息，通过token判断用户权限
	token := c.GetHeader("token")
	level := jwt.GetUserLevel(token)
	code := e.SUCCESS

	fmt.Println(level)
	// 获取修改信息，判断权限是否合理
	var user models.Message
	c.BindJSON(&user)
	fmt.Println(user)
	mUser, err := models.GetUserInfo(user.ID)
	// 传入id无效、被修改用户权限大于等于当前用户、修改后权限大于当前用户时返回参数错误
	if err != nil || level >= mUser.Level || user.Level < level || user.Level < models.SuperAdmin || user.Level > models.User {
		code = e.INVALID_PARAMS
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
		})
		return
	}

	fmt.Println(user)
	// 修改用户信息
	err = models.UpdateUser(user)
	if err != nil {
		code = e.INVALID_PARAMS
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
	})
}
