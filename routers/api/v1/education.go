/**
 @author: 李映飞
 @date:   2022/6/6
 @note:
**/
package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"webDesign/models"
	"webDesign/pkg/e"
)

type ViewFaculty struct {
	CODE string `json:"code"`
	Name string `json:"name"`
}

func AddFaculty(c *gin.Context) {
	var faculty ViewFaculty
	code := e.SUCCESS
	if err := c.ShouldBindJSON(&faculty); err != nil {
		code = e.INVALID_PARAMS
		c.JSON(200, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
		})
		return
	}
	fid := models.AddFaculty(models.Faculty{
		FCODE: faculty.CODE,
		Name:  faculty.Name,
	})
	if fid == 0 {
		code = e.ERROR
	}
	c.JSON(200, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"fid":  fid,
	})
}

type ID struct {
	ID uint `json:"id"`
}

func DeleteFaculty(c *gin.Context) {

	var id ID
	code := e.SUCCESS
	if err := c.ShouldBindJSON(&id); err != nil {
		code = e.INVALID_PARAMS
		c.JSON(200, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
		})
		return
	}
	if err := models.DeleteFaculty(id.ID); err != nil {
		code = e.ERROR
	}
	c.JSON(200, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
	})
}

type UpdateDepartment struct {
	ID   uint   `json:"id"`
	CODE string `json:"code"`
	Name string `json:"name"`
}

func AddDepartment(c *gin.Context) {

	var department UpdateDepartment
	code := e.SUCCESS
	err := c.ShouldBindJSON(&department)
	if err != nil {
		code = e.INVALID_PARAMS
		c.JSON(200, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
		})
		return
	}
	did := models.AddDepartment(models.Department{
		FID:   department.ID,
		DCODE: department.CODE,
		Name:  department.Name,
	})
	if did == 0 {
		code = e.ERROR
	}
	c.JSON(200, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"did":  did,
	})
}

func DeleteDepartment(c *gin.Context) {

	var id ID
	code := e.SUCCESS
	if err := c.ShouldBindJSON(&id); err != nil {
		code = e.INVALID_PARAMS
		c.JSON(200, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
		})
		return
	}
	if err := models.DeleteDepartment(id.ID); err != nil {
		code = e.ERROR
	}
	c.JSON(200, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
	})
}

func AddSession(c *gin.Context) {
	var session models.Session
	code := e.SUCCESS
	if err := c.ShouldBindJSON(&session); err != nil || len(session.Session) != 2 || session.Session[0] < '0' || session.Session[0] > '9' || session.Session[1] < '0' || session.Session[1] > '9' {
		code = e.INVALID_PARAMS
		c.JSON(200, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
		})
		return
	}
	if err := models.AddSession(session); err != nil {
		code = e.ERROR
	}
	c.JSON(200, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
	})
}

func DeleteSession(c *gin.Context) {
	var session models.Session
	code := e.SUCCESS
	if err := c.ShouldBindJSON(&session); err != nil {
		code = e.INVALID_PARAMS
		c.JSON(200, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
		})
		return
	}
	if err := models.DeleteSession(session); err != nil {
		code = e.ERROR
	}
	c.JSON(200, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
	})
}

func AddClass(c *gin.Context) {
	var class models.Class
	code := e.SUCCESS
	if err := c.ShouldBindJSON(&class); err != nil {
		code = e.INVALID_PARAMS
		c.JSON(200, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
		})
		return
	}
	if err := models.AddClass(&class); err != nil {
		code = e.ERROR
	}
	c.JSON(200, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"sid":  class.SID,
	})
}

func DeleteClass(c *gin.Context) {
	var id ID
	code := e.SUCCESS
	if err := c.ShouldBindJSON(&id); err != nil {
		code = e.INVALID_PARAMS
		c.JSON(200, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
		})
		return
	}
	if err := models.DeleteClass(id.ID); err != nil {
		code = e.ERROR
	}
	c.JSON(200, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
	})
}

func AddCourse(c *gin.Context) {
	var course models.Course
	code := e.SUCCESS
	if err := c.ShouldBindJSON(&course); err != nil {
		code = e.INVALID_PARAMS
		c.JSON(200, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
		})
		return
	}

	if err := models.AddCourse(&course); err != nil {
		code = e.ERROR
	}

	c.JSON(200, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"cid":  course.CID,
	})
}

func DeleteCourse(c *gin.Context) {
	var id ID
	code := e.SUCCESS
	if err := c.ShouldBindJSON(&id); err != nil {
		code = e.INVALID_PARAMS
		c.JSON(200, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
		})
		return
	}
	if err := models.DeleteCourse(id.ID); err != nil {
		code = e.ERROR
	}
	c.JSON(200, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
	})
}

func AllCourses(c *gin.Context) {
	courses := models.GetAllCourses()
	c.JSON(200, gin.H{
		"code": e.SUCCESS,
		"msg":  e.GetMsg(e.SUCCESS),
		"data": courses,
	})
}

func AddDC(c *gin.Context) {
	var dc models.DC
	code := e.SUCCESS
	if err := c.ShouldBindJSON(&dc); err != nil {
		code = e.INVALID_PARAMS
		c.JSON(200, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
		})
		return
	}
	if err := models.AddDC(&dc); err != nil {
		code = e.ERROR
	}
	c.JSON(200, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
	})
}

func DeleteDC(c *gin.Context) {
	var dc models.DC
	code := e.SUCCESS
	if err := c.ShouldBindJSON(&dc); err != nil {
		code = e.INVALID_PARAMS
		c.JSON(200, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
		})
		return
	}

	fmt.Println(dc)
	if err := models.DeleteDC(&dc); err != nil {
		code = e.ERROR
	}

	c.JSON(200, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
	})
}
