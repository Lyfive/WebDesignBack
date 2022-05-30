/**
 @author: 李映飞
 @date:   2022/4/29
 @note:
**/
package models

import (
	"fmt"
	"log"
	"webDesign/pkg/e"
)

type Course struct {
	// 课程ID
	CID uint `json:"cid" gorm:"primaryKey;autoIncrement"`
	// 课程名
	Title string `json:"title" gorm:"not null;varchar(10)"`
}

type Grade struct {
	// 成绩ID
	GID uint `json:"gid" gorm:"primaryKey;autoIncrement"`
	// 绑定学生学号
	Number string `json:"number" gorm:"not null;type:char(10)"`
	// 绑定课程ID
	CID uint `json:"cid" gorm:"not null"`
	// 分数
	Mark uint `json:"mark" gorm:"not null;type:tinyint"`
}

// ViewGrade 前端展示成绩
type ViewGrade struct {
	Number string `json:"number"`
	Name   string `json:"name"`
	Title  string `json:"title"`
	Mark   uint   `json:"mark"`
}

// GetCourse  获取课程信息
func GetCourse(title string) Course {
	var course Course
	if err := db.First(&course, "title = ?", title).Error; err != nil {
		return Course{}
	} else {
		return course
	}
}

// AddGrade 添加成绩
func AddGrade(grade Grade) int {
	var gid uint
	db.Model(&Grade{}).Omit("number", "c_id", "mark").Where("number = ? and c_id = ?", grade.Number, grade.CID).First(&gid)
	if gid > 0 {
		grade.GID = gid
		if err := db.Debug().Model(&grade).Update("mark", grade.Mark).Error; err != nil {
			log.Printf("error:%v", err)
			return e.ERROR
		}
		return e.SUCCESS
	}
	if err := db.Omit("g_id").Create(&grade).Error; err != nil {
		return e.INVALID_PARAMS
	}
	return e.SUCCESS
}

// FindGrades 根据学号查询相关成绩
func FindGrades(number string) []ViewGrade {
	var grades []ViewGrade
	db.Debug().Debug().Raw("select students.name, students.number, courses.title,grades.mark "+
		"from grades,students,courses "+
		"where students.number like ? and grades.c_id = courses.c_id and students.number = grades.number", number).Scan(&grades) // 模糊查询
	return grades
}

// QueryGrades 根据姓名查询相关成绩
func QueryGrades(matchStr string, isOpen bool) []ViewGrade {
	var grades []ViewGrade
	mod := "="
	if isOpen {
		mod = "like"
	}
	db.Debug().Raw("select students.name,students.number,courses.title,grades.mark "+
		"from grades,students,courses "+
		"where students.name "+
		mod+
		" ? and grades.number = students.number and grades.c_id = courses.c_id", matchStr).
		Scan(&grades)
	return grades
}

// GetDepartmentCourses  根据class ID获取班级所学课程
func GetDepartmentCourses(DID uint) []Course {
	var courses []Course
	db.Raw("select courses.c_id c_id,courses.title title from dc,courses where dc.d_id = ? and dc.c_id = courses.c_id", DID).
		Scan(&courses)
	//fmt.Println(courses)
	return courses
}

type MiniStudent struct {
	Number string `gorm:"column:number"`
	Name   string `gorm:"column:name"`
}

func GetClassStudents(SID uint) []MiniStudent {
	var students []MiniStudent
	db.Raw("select number,name from students "+
		"where students.s_id = ?", SID).
		Scan(&students)
	fmt.Println(students)
	return students
}

type View struct {
	Number string `gorm:"column:number"`
	Title  string `gorm:"column:title"`
	Mark   uint   `gorm:"column:mark"`
}

// GetStudentsGrades 根据班级ID 查询该班级学生对应课程成绩
func GetStudentsGrades(SID uint) []View {
	var grades []View
	db.Raw("select students.number,courses.title,mark from students,grades,courses "+
		"where students.s_id = ? and students.number = grades.number and grades.c_id = courses.c_id", SID).
		Scan(&grades)
	fmt.Println(grades)
	return grades
}
