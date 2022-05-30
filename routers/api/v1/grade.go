/**
 @author: 李映飞
 @date:   2022/4/29
 @note:
**/
package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
	"net/http"
	"strconv"
	"webDesign/models"
	"webDesign/pkg/e"
)

// Add 添加成绩
func Add(c *gin.Context) {
	var grade models.Grade
	c.BindJSON(&grade)
	code := models.AddGrade(grade)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
	})
}

type GradeMsg struct {
	Title string `json:"title"`
	Mark  uint   `json:"mark"`
}
type Data struct {
	Number string     `json:"number"`
	Name   string     `json:"name"`
	Grades []GradeMsg `json:"grades"`
}

// 将view数据转换成返回的Data数据
func exchange(grades []models.ViewGrade) []Data {
	data := make(map[string]*Data)
	for _, v := range grades {
		if _, ok := data[v.Number]; !ok {
			data[v.Number] = &Data{
				Number: v.Number,
				Name:   v.Name,
				Grades: make([]GradeMsg, 0),
			}
		}
		data[v.Number].Grades = append(data[v.Number].Grades, GradeMsg{Title: v.Title, Mark: v.Mark})
	}

	length := len(data)
	slice := make([]Data, length)
	for _, v := range data {
		length--
		slice[length] = *v
	}
	return slice
}

// Find 学号查询成绩
func Find(c *gin.Context) {
	number := c.Query("number")
	var grades []models.ViewGrade
	code := e.SUCCESS
	grades = models.FindGrades(number)

	if len(grades) == 0 {
		code = e.ERROR_NOT_EXIST_NUMBER
	}
	fmt.Println(grades)

	c.JSON(http.StatusOK, gin.H{
		"code":       code,
		"msg":        e.GetMsg(code),
		"gradesData": exchange(grades),
	})
}

// Query 姓名模糊查询成绩
func Query(c *gin.Context) {
	name := c.Query("name")
	open := c.Query("open")
	isOpen := false
	if open == "Yes" {
		isOpen = true
	}
	grades := models.QueryGrades(name, isOpen)
	fmt.Println(grades)
	c.JSON(http.StatusOK, gin.H{
		"code":       e.SUCCESS,
		"gradesData": exchange(grades),
	})
}

func GetAxis(row int, column int) string {
	return string(row+'A') + strconv.Itoa(column)
}

type GradeList struct {
	Name   string
	Grades map[string]uint
}

func Create(c *gin.Context) {

	number := c.Query("number")
	// 获取班级 ID
	class := models.GetClass(number)

	if class.SID <= 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": e.ERROR_NOT_EXIST_CLASS,
			"msg":  e.GetMsg(e.ERROR_NOT_EXIST_CLASS),
		})
	} else {
		// 根据系ID 查询班级所有学的课程
		courses := models.GetDepartmentCourses(class.DID)
		fmt.Println(courses)

		// 班级学生列表
		students := models.GetClassStudents(class.SID)
		fmt.Println(students)
		scnt := len(students)

		// 成绩表
		grades := models.GetStudentsGrades(class.SID)
		fmt.Println(grades)

		// 成绩信息
		gradeMessage := make(map[string]GradeList)
		for _, student := range students {
			gradeMessage[student.Number] = GradeList{Name: student.Name, Grades: make(map[string]uint)}
		}

		// 总成绩
		avg := make(map[string]uint)
		for _, grade := range grades {
			avg[grade.Title] += grade.Mark
			avg["班级总成绩"] += grade.Mark
			gradeMessage[grade.Number].Grades[grade.Title] = grade.Mark
			gradeMessage[grade.Number].Grades["总成绩"] += grade.Mark
		}

		// 创建file
		f := excelize.NewFile()

		// 创建样式
		s, _ := f.NewStyle(`{"alignment":{
		"horizontal":"center",
		"vertical":"center"
	},"number_format":0
	,"fill":{
		"type":"pattern",
		"pattern":1
	}}`)
		f.SetSheetName("Sheet1", "s1")
		f.Save()
		err := f.SetCellValue("s1", "A1", "学号")
		f.SetColWidth("s1", "A", "A", 20)
		if err != nil {
			fmt.Println(err)
			return
		}
		f.SetCellValue("s1", "B1", "姓名")
		f.SetColWidth("s1", "B", "B", 10)
		index := 2
		for _, course := range courses {
			f.SetCellValue("s1", GetAxis(index, 1), course.Title)
			f.SetColWidth("s1", string('A'+index), string('A'+index), 2*float64(len(course.Title)))
			index++
			f.SetCellValue("s1", GetAxis(index, 1), course.Title+"平均值")
			f.SetColWidth("s1", string('A'+index), string('A'+index), 2*float64(len(course.Title)+3))
			index++
		}
		f.SetCellValue("s1", GetAxis(index, 1), "总成绩")
		f.SetColWidth("s1", string('A'+index), string('A'+index), 15)
		index++
		f.SetCellValue("s1", GetAxis(index, 1), "班级总成绩平均值")
		f.SetColWidth("s1", string('A'+index), string('A'+index), 20)
		f.SetColStyle("s1", "A:"+string('A'+index), s)
		index = 2
		for number, gradeList := range gradeMessage {
			col := 0
			f.SetCellValue("s1", GetAxis(col, index), number)
			col++
			f.SetCellValue("s1", GetAxis(col, index), gradeList.Name)
			col++
			for _, course := range courses {
				f.SetCellValue("s1", GetAxis(col, index), gradeList.Grades[course.Title])
				col++
				f.SetCellValue("s1", GetAxis(col, index), float64(avg[course.Title])/float64(scnt))
				col++
			}
			// 个人总成绩
			f.SetCellValue("s1", GetAxis(col, index), gradeList.Grades["总成绩"])
			col++
			f.SetCellValue("s1", GetAxis(col, index), float64(avg["班级总成绩"])/float64(scnt))
			index++
		}

		f.SetSheetName("s1", class.Name+"班级成绩表")
		filePath := "./runtime/file/Grade.xlsx"
		err = f.SaveAs(filePath)
		if err != nil {
			return
		}
		err = f.Close()
		if err != nil {
			fmt.Println(err)
		}
		c.File(filePath)
	}

}
