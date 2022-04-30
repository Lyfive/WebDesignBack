/**
 @author: 李映飞
 @date:   2022/4/28
 @note:
**/
package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"webDesign/middleware/jwt"
	"webDesign/models"
	setting "webDesign/pkg"
	"webDesign/routers/api"
	v1 "webDesign/routers/api/v1"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.RunMode)

	r.POST("/login", api.GetUser)

	// 用户组
	users := r.Group("/user")
	users.Use(jwt.JWT(models.SuperAdmin))
	{
		users.POST("/register", v1.Register)
	}

	students := r.Group("/student")
	//students.Use(jwt.JWT(models.Admin))
	{
		// 加入学生
		students.POST("/join", v1.Join)

		// 查看学生信息
		students.GET("/view", v1.View)

		// 修改学生信息
		students.PUT("/modify", v1.Modify)

		// 删除学生
		students.DELETE("/delete", v1.Delete)

		// 生成学生成绩报表
		students.GET("/create", v1.Create)
	}

	// 成绩组
	grades := r.Group("/grade")
	grades.Use(jwt.JWT(models.Admin))
	{
		// 添加成绩
		grades.POST("/add", v1.Add)

		// 学号成绩查询
		grades.GET("/find", v1.Find)

		// 姓名模糊成绩查询
		grades.GET("/query", v1.Query)
	}

	return r

}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", "*") // 可将将 * 替换为指定的域名
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")
		}
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}
