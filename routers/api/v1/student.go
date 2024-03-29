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
	"strconv"
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

func StringToUint(str string) uint {
	atoi, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}
	return uint(atoi)
}
func Search(c *gin.Context) {
	//fmt.Println(c.Query("number"))

	fid := c.Query("fid")
	did := c.Query("did")
	sid := c.Query("sid")
	session := c.Query("session")
	if fid == "" {
		fid = "%"
	}
	if did == "" {
		did = "%"
	}
	if sid == "" {
		sid = "%"
	}
	if session == "" {
		session = "%"
	}
	viewStudent := models.FindStudentByMessage(fid, did, sid, session)

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

type Numbers struct {
	Numbers []string `json:"numbers"`
	ID      uint     `json:"sid"`
}

func Delete(c *gin.Context) {
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

func Transfer(c *gin.Context) {
	numbers := Numbers{
		Numbers: make([]string, 0),
	}
	c.BindJSON(&numbers)
	code := e.ERROR
	if numbers.ID != 0 {
		code = models.TransferStudents(&numbers.Numbers, numbers.ID)
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
	})
}
