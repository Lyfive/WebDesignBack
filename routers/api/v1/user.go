/**
 @author: 李映飞
 @date:   2022/4/28
 @note:
**/
package v1

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"webDesign/models"
	"webDesign/pkg/e"
)

func Register(c *gin.Context) {
	var message models.Message
	err := c.BindJSON(&message)

	if err != nil {
		log.Fatalf("Error from Register v1/user.go :%v", err)
		return
	}

	code := models.AddMessage(message)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
	})
}
