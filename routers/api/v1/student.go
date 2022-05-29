/**
 @author: 李映飞
 @date:   2022/4/29
 @note:
**/
package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"webDesign/models"
	"webDesign/pkg/e"
)

func Join(c *gin.Context) {
	var studentModel models.StudentModel
	c.BindJSON(&studentModel)
	fmt.Println(studentModel)
	fmt.Println(123)
	code := models.JoinStudent(studentModel)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
	})
}

func View(c *gin.Context) {
	//fmt.Println(c.Query("number"))
	number := c.Query("number")
	fmt.Println(number)
	viewStudent := models.FindStudent(number)
	code := e.SUCCESS
	if len(viewStudent) == 0 {
		code = e.ERROR_NOT_EXIST_NUMBER
	}
	for i := 0; i < len(viewStudent); i++ {
		//2002-11-15T00:00:00+08:00
		mid, _ := time.ParseInLocation("2006-01-02T15:04:05+08:00", viewStudent[i].Birthday, time.Local)
		viewStudent[i].Birthday = mid.Format("2006-01-02")
	}
	c.JSON(http.StatusOK, gin.H{
		"code":     code,
		"msg":      e.GetMsg(code),
		"students": viewStudent,
	})
}

func Modify(c *gin.Context) {
	var student models.StudentInfo
	fmt.Println(c.Request)
	c.BindJSON(&student)
	fmt.Println(student)
	code := models.UpdateStudent(student)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
	})
}

func Delete(c *gin.Context) {
	type Numbers struct {
		Numbers []string `json:"numbers"`
	}
	numbers := Numbers{
		Numbers: make([]string, 0),
	}
	c.BindJSON(&numbers)
	fmt.Println(numbers)
	code := models.DeleteStudents(&numbers.Numbers)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
	})
}
