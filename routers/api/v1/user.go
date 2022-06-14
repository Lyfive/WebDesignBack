/**
 @author: 李映飞
 @date:   2022/4/28
 @note:
**/
package v1

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"webDesign/models"
	setting "webDesign/pkg"
	"webDesign/pkg/crypto"
	"webDesign/pkg/e"
	"webDesign/pkg/util"
)

// Register 用户注册
// @Summary 用户注册
// @Description 在用户界面为用户注册，可以进行提权、添加头像
// @Tags 用户
// @Accept json
// @Produce json
// @Param message body models.Message true "用户注册信息"
// @Success 200 "SUCCESS"
// @Failure 401 "权限不足"
// @Router /user/register [post]
func Register(c *gin.Context) {
	var message models.Message

	// 拉取用户信息，通过token判断用户权限
	c.BindJSON(&message)
	token := c.GetHeader("token")
	level := util.GetUserLevel(token)
	code := e.SUCCESS

	// 头像上传比较麻烦 若头像为空则使用默认头像
	if message.Head == "" {
		message.Head = "30616305ef290e389d9019a0683a8046"
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

// DeleteUser 删除用户
// @Summary 删除用户
// @Description 删除用户，只有管理员可以删除用户
// @Tags 用户
// @Accept json
// @Produce json
// @Param id body UserID true "用户ID"
// @Success 200 "SUCCESS"
// @Failure 401 "权限不足"
// @Router /user/delete [delete]
func DeleteUser(c *gin.Context) {

	var id UserID
	// 拉取用户信息，通过token判断用户权限
	token := c.GetHeader("token")
	level := util.GetUserLevel(token)
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

// GetUserList 获取用户列表
// @Summary 获取用户列表
// @Description 获取用户列表，只有管理员可以获取用户列表
// @Tags 用户
// @Accept json
// @Produce json
// @Success 200 "SUCCESS"
// @Failure 401 "权限不足"
// @Router /user/users [get]
func GetUserList(c *gin.Context) {

	// 拉取用户信息，通过token判断用户权限
	token := c.GetHeader("token")
	level := util.GetUserLevel(token)
	code := e.SUCCESS

	// 获取等级小于当前用户的用户列表
	users := models.GetUsers(level)

	c.JSON(http.StatusOK, gin.H{
		"code":  code,
		"msg":   e.GetMsg(code),
		"users": users,
	})
}

// ModifyUser 修改用户信息
// @Summary 修改用户信息
// @Description 高权限可以修改低权限信息，但是提权和自己同水平就不能修改
// @Tags 用户
// @Accept json
// @Produce json
// @Param user body models.Message true "修改后的用户信息"
// @Success 200 "SUCCESS"
// @Failure 401 "权限不足"
// @Router /user/modify [put]
func ModifyUser(c *gin.Context) {

	// 拉取用户信息，通过token判断用户权限
	token := c.GetHeader("token")
	level := util.GetUserLevel(token)
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

func Upload(c *gin.Context) {
	code := e.ERROR
	url := c.PostForm("url")
	id := c.PostForm("id")
	picture := ""
	var fileName string
	if id != "" {
		if url != "" {
			fileName = beego.Substr(url, 12, len(url)-16)
		} else {
			fileHeader, _ := c.FormFile("img")
			// 如果获取图片成功
			// 读取图片file并生成对应的md5
			file, _ := fileHeader.Open()
			fileName = crypto.EncryptFile(file)

			dst := "./dist/static/img/" + fileName + ".jpg"
			c.SaveUploadedFile(fileHeader, dst)
		}
		picture = setting.HOST + "/static/img/" + fileName + ".jpg"
		uid, err := strconv.Atoi(id)
		if err == nil {
			err := models.UpdateHead(fileName, uint(uid))
			if err == nil {
				code = e.SUCCESS
			}
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"url":  picture,
	})
}

func ModifyPassword(c *gin.Context) {

}
