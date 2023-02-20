/*
*

	@author: 李映飞
	@date:   2022/4/29
	@note:

*
*/
package models

import (
	"fmt"
	"time"
	"webDesign/pkg/e"
)

type Student struct {
	// 学号
	Number   string    `json:"number" gorm:"primaryKey;type:char(10)"`
	Name     string    `json:"name" gorm:"not null;type:varchar(6)"`
	Sex      string    `json:"sex" gorm:"not null;char(2)"`
	Birthday time.Time `json:"birthday" gorm:"not null;type:date"`
	// 归属班级ID
	SID uint `json:"sid" gorm:"not null"`
}

type StudentInfo struct {
	// 学号
	Number   string `json:"number" gorm:"primaryKey;type:char(10)"`
	Name     string `json:"name" gorm:"not null;type:varchar(6)"`
	Sex      string `json:"sex" gorm:"not null;char(2)"`
	Birthday string `json:"birthday" gorm:"not null;type:date"`
	// 归属班级ID
	SID uint `json:"sid" gorm:"not null"`
}

type ViewStudent struct {
	Number   string `json:"number" gorm:"column:number"`
	Name     string `json:"name" gorm:"column:name"`
	Sex      string `json:"sex" gorm:"column:sex"`
	Birthday string `json:"birthday" gorm:"column:birthday"`
	Class    string `json:"class" gorm:"column:class"`
	Session  string `json:"session" gorm:"column:session"`
}

// CheckStudent 检查是否存在学生
func CheckStudent(number string) bool {
	if err := db.First(&Student{}, number).Error; err != nil {
		return false
	}
	return true
}

type StudentModel struct {
	Name     string `json:"name"`
	Sex      string `json:"sex"`
	Birthday string `json:"birthday"`
	SID      uint   `json:"sid"`
	Session  string `json:"session"`
}

// JoinStudent 添加学生
func JoinStudent(studentModel StudentModel) int {
	// 分别获得学生的班级、系、学院和班级人数生成学号
	birthday, _ := time.Parse("2006-01-02", studentModel.Birthday)
	c := GetClassByID(studentModel.SID)
	d := GetDepartmentByID(c.DID)
	f := GetFacultyByID(d.FID)

	var cnt uint
	number := studentModel.Session // 学号年级
	number += f.FCODE              // 学号学院
	number += d.DCODE              // 学号系
	number += c.SCODE              // 学号班级
	db.Raw("select count(*) from students where s_id = ?", c.SID).Scan(&cnt)
	number = fmt.Sprintf("%s%02d", number, cnt+1)
	stu := Student{
		Number:   number,
		Name:     studentModel.Name,
		Sex:      studentModel.Sex,
		Birthday: birthday,
		SID:      c.SID,
	}

	if err := db.Debug().Create(&stu).Error; err != nil {

		// 获取班级最后一个学生
		var lastStu Student
		db.Last(&lastStu).Where("SID = ?", c.SID)

		str := lastStu.Number[8:]
		fmt.Println(str)
		fmt.Sscanf(str, "%02d", &cnt)
		cnt = cnt + 1
		stu.Number = fmt.Sprintf("%s%02d", lastStu.Number[0:8], cnt)

		fmt.Println(stu.Number)
		if err := db.Debug().Create(&stu).Error; err != nil {
			return e.ERROR
		}
	}
	return e.SUCCESS
}

// FindStudent 查询学生信息
func FindStudent(number string) []ViewStudent {
	stu := make([]ViewStudent, 0)
	if err := db.Raw("select students.number number,students.name name,students.sex sex,students.birthday birthday,classes.name class,classes.session session from students,classes where number like ? and students.s_id = classes.s_id", number).Scan(&stu).Error; err != nil {
		return []ViewStudent{}
	} else {
		return stu
	}
}

func FindStudentByMessage(fid, did, sid, session string) []ViewStudent {
	stu := make([]ViewStudent, 0)
	db.Raw("select students.number number,students.name name,students.sex sex,students.birthday birthday,classes.name class,classes.session session from students,classes,departments,faculties where classes.session like ? and classes.s_id like ? and departments.d_id like ? and faculties.f_id like ? and students.s_id = classes.s_id and classes.d_id = departments.d_id and departments.f_id = faculties.f_id", session, sid, did, fid).Scan(&stu)
	return stu
}

// UpdateStudent 更新学生信息
func UpdateStudent(studentInfo StudentInfo) int {
	birthday, _ := time.Parse("2006-01-02", studentInfo.Birthday)
	stu := Student{
		Number:   studentInfo.Number,
		Name:     studentInfo.Name,
		Sex:      studentInfo.Sex,
		SID:      studentInfo.SID,
		Birthday: birthday,
	}
	if err := db.Updates(stu).Error; err != nil {
		return e.ERROR
	}
	return e.SUCCESS
}

// DeleteStudents 删除学生信息
func DeleteStudents(numbers *[]string) int {
	if err := db.Delete(&Student{}, *numbers).Error; err != nil {
		return e.ERROR_NOT_EXIST_NUMBER
	} else {
		return e.SUCCESS
	}
}

// TransferStudents 批量转入班级
func TransferStudents(numbers *[]string, sid uint) int {
	if err := db.Model(&Student{}).Where("number in (?)", *numbers).Update("s_id", sid).Error; err != nil {
		return e.ERROR
	}
	return e.SUCCESS
}
