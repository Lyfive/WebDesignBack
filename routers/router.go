/**
 @author: 李映飞
 @date:   2022/4/28
 @note:
**/
package routers

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"

	_ "webDesign/docs"
	"webDesign/middleware/jwt"
	"webDesign/models"
	setting "webDesign/pkg"
	"webDesign/routers/api"
	v1 "webDesign/routers/api/v1"
)

// @title           Student Management System
// @version         1.0
// @description     使用Go+Gin+Vue3+Element-plus的框架开发的学生管理系统管理后台
// @termsOfService  http://swagger.io/terms/

// @contact.name   LyFive
// @contact.url    https://lyfive.github.io/
// @contact.email  1169442146@qq.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /

// @securityDefinitions.basic  BasicAuth
func InitRouter() *gin.Engine {

	gin.SetMode(setting.RunMode)
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(Cors())
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Token")
	})

	// 静态文件处理
	{
		r.StaticFile("/favicon.ico", "./dist/favicon.ico")
		r.StaticFile("/22_open.png", "./dist/22_open.png")
		r.StaticFile("/22_close.png", "./dist/22_close.png")
		r.StaticFile("/33_open.png", "./dist/33_open.png")
		r.StaticFile("/33_close.png", "./dist/33_close.png")
		r.Static("/static", "./dist/static")
		r.LoadHTMLFiles("dist/index.html")
	}

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", "./index.html")
	})

	// swagger
	{
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	// 重定向
	{
		r.GET("/login", func(c *gin.Context) {
			c.Redirect(http.StatusMovedPermanently, "/")
		})

		r.GET("/manage", func(c *gin.Context) {
			c.Redirect(http.StatusMovedPermanently, "/")
		})

		r.GET("/home", func(c *gin.Context) {
			c.Redirect(http.StatusMovedPermanently, "/")
		})

		r.GET("/userList", func(c *gin.Context) {
			c.Redirect(http.StatusMovedPermanently, "/")
		})

		r.GET("/modifyUser", func(c *gin.Context) {
			c.Redirect(http.StatusMovedPermanently, "/")
		})

		r.GET("/registerUser", func(c *gin.Context) {
			c.Redirect(http.StatusMovedPermanently, "/")
		})

		r.GET("/studentList", func(c *gin.Context) {
			c.Redirect(http.StatusMovedPermanently, "/")
		})

		r.GET("/queryGrade", func(c *gin.Context) {
			c.Redirect(http.StatusMovedPermanently, "/")
		})

		r.GET("/updateGrade", func(c *gin.Context) {
			c.Redirect(http.StatusMovedPermanently, "/")
		})

		r.GET("/admin", func(c *gin.Context) {
			c.Redirect(http.StatusMovedPermanently, "/")
		})

		r.GET("/education", func(c *gin.Context) {
			c.Redirect(http.StatusMovedPermanently, "/")
		})

	}

	// 用户鉴定
	r.POST("/check", api.CheckUser)

	r.POST("/login", api.GetUser)
	// 原界面上的注册 默认注册为普通用户
	r.POST("/register", api.RegisterUser)

	// 用户组
	users := r.Group("/user")
	{
		// 高级注册 可以提权
		users.POST("/register", jwt.JWT(models.Admin), v1.Register)

		// 删除用户，只能删除权限比自己低的用户
		users.DELETE("/delete", jwt.JWT(models.Admin), v1.DeleteUser)

		// 获取可以修改的用户列表
		users.GET("/users", jwt.JWT(models.Admin), v1.GetUserList)

		// 修改用户信息 要求被修改者权限比自己低，同时被修改后权限不能比自己高
		users.PUT("/modify", jwt.JWT(models.Admin), v1.ModifyUser)

		// 上传头像
		users.POST("/upload", jwt.JWT(models.User), v1.Upload)

		// 修改密码
		users.PUT("/modifyPassword", jwt.JWT(models.User), v1.ModifyPassword)
	}

	students := r.Group("/student")
	{
		// 加入学生
		students.POST("/join", jwt.JWT(models.Admin), v1.Join)

		// 查看学生信息
		students.GET("/view", jwt.JWT(models.User), v1.View)

		// 修改学生信息
		students.PUT("/modify", jwt.JWT(models.Admin), v1.Modify)

		// 删除学生
		students.DELETE("/delete", jwt.JWT(models.Admin), v1.Delete)

		// 集体转出其他班级
		students.PUT("/transfer", jwt.JWT(models.Admin), v1.Transfer)
	}

	// 成绩组
	grades := r.Group("/grade")
	{
		// 添加成绩
		grades.POST("/add", jwt.JWT(models.Admin), v1.Add)

		// 学号成绩查询
		grades.GET("/find", jwt.JWT(models.User), v1.Find)

		// 姓名模糊成绩查询
		grades.GET("/query", jwt.JWT(models.User), v1.Query)

		// 生成学生成绩报表
		grades.GET("/create", jwt.JWT(models.User), v1.Create)
	}

	// 中间信息
	mid := r.Group("/mid")
	mid.Use(jwt.JWT(models.User))
	{
		// 查询学院
		mid.GET("/faculties", v1.FindFaculties)

		// 查询系
		mid.GET("/departments", v1.FindDepartments)

		// 查询系学年
		mid.GET("/sessions", v1.FindSessions)

		// 查询班级
		mid.GET("/classes", v1.FindClasses)

		// 查询班级学生
		mid.GET("/students", v1.FindStudents)

		// 查询应修课程
		mid.GET("/courses", v1.FindCourses)

	}

	// 教务处理
	education := r.Group("/education")
	education.Use(jwt.JWT(models.Admin))
	{
		education.POST("/faculty", v1.AddFaculty)
		education.DELETE("/faculty", v1.DeleteFaculty)

		education.POST("/department", v1.AddDepartment)
		education.DELETE("/department", v1.DeleteDepartment)

		education.POST("/session", v1.AddSession)
		education.DELETE("/session", v1.DeleteSession)

		education.POST("/class", v1.AddClass)
		education.DELETE("/class", v1.DeleteClass)

		education.POST("/course", v1.AddCourse)
		education.DELETE("/course", v1.DeleteCourse)
		education.GET("/courses", v1.AllCourses)

		education.POST("/dc", v1.AddDC)
		education.DELETE("/dc", v1.DeleteDC)
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
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization, Token")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")
		}
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}
