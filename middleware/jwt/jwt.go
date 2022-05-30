/**
 @author: 李映飞
 @date:   2022/4/28
 @note:
**/
package jwt

import (
	"fmt"
	"net/http"
	"time"
	"webDesign/models"
	"webDesign/pkg/e"
	"webDesign/pkg/util"

	"github.com/gin-gonic/gin"
)

func GetUserLevel(token string) int {
	claims, err := util.ParseToken(token)
	if err != nil {
		return 4
	}
	return models.GetUser(claims.Username).Level
}

func JWT(level int) gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		code = e.SUCCESS
		token := c.GetHeader("token")

		//token := c.Query("token")
		if token == "" {
			code = e.INVALID_PARAMS
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			} else {
				msg := models.GetUser(claims.Username)
				if msg.Level > level {
					code = e.ERROR_INSUFFICIENT_ACCESS_RIGHTS
				}
			}
		}
		fmt.Println(code)
		if code != e.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
			})
			c.Abort()
			return
		}

		fmt.Println(code)
		c.Next()
	}
}
