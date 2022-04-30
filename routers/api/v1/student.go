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
	"webDesign/models"
	"webDesign/pkg/e"
)

func Join(c *gin.Context) {
	var viewStudent models.ViewStudent
	c.BindJSON(&viewStudent)
	fmt.Println(viewStudent)

	code := models.JoinStudent(viewStudent)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
	})
}

func View(c *gin.Context) {
	number := c.Query("number")
	viewStudent := models.FindStudent(number)
	code := e.SUCCESS
	if viewStudent.Number == "" {
		code = e.ERROR_NOT_EXIST_NUMBER
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"msg":     e.GetMsg(code),
		"student": viewStudent,
	})
}

func Modify(c *gin.Context) {
	var viewStudent models.ViewStudent
	c.BindJSON(&viewStudent)
	code := models.UpdateStudent(viewStudent)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
	})
}

func Delete(c *gin.Context) {
	numbers := c.QueryArray("numbers")

	code := models.DeleteStudents(&numbers)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
	})
}

func Create(c *gin.Context) {

}
