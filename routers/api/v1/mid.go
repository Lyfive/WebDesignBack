/**
 @author: 李映飞
 @date:   2022/5/22
 @note:
**/
package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"webDesign/models"
	"webDesign/pkg/e"
)

func FindFaculties(c *gin.Context) {
	faculties, err := models.GetAllFaculties()
	code := e.SUCCESS
	if err != nil {
		code = e.ERROR
	}
	c.JSON(code, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": faculties,
	})
}

func FindDepartments(c *gin.Context) {
	facultyId := c.Query("fid")

	id, _ := strconv.Atoi(facultyId)

	departments, err := models.GetAllDepartments(uint(id))
	code := e.SUCCESS
	if err != nil {
		code = e.ERROR
	}
	c.JSON(code, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": departments,
	})
}

func FindSessions(c *gin.Context) {
	departmentId := c.Query("did")

	id, _ := strconv.Atoi(departmentId)

	sessions, err := models.GetAllSessions(uint(id))
	code := e.SUCCESS
	if err != nil {
		code = e.ERROR
	}
	c.JSON(code, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": sessions,
	})
}

func FindClasses(c *gin.Context) {
	departmentId := c.Query("did")
	session := c.Query("session")

	did, _ := strconv.Atoi(departmentId)

	classes, err := models.GetAllClasses(uint(did), session)
	code := e.SUCCESS
	if err != nil {
		code = e.ERROR
	}
	c.JSON(code, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": classes,
	})
}

func FindStudents(c *gin.Context) {
	classId := c.Query("sid")

	id, _ := strconv.Atoi(classId)

	students, err := models.GetStudentsByClass(uint(id))
	code := e.SUCCESS
	if err != nil {
		code = e.ERROR
	}
	c.JSON(code, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": students,
	})
}

func FindCourses(c *gin.Context) {
	departmentId := c.Query("did")

	id, _ := strconv.Atoi(departmentId)
	fmt.Println(id)
	courses, err := models.GetCoursesByDepartment(uint(id))
	fmt.Println(courses)
	code := e.SUCCESS
	if err != nil {
		code = e.ERROR
	}
	c.JSON(code, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": courses,
	})
}
