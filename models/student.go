/**
 @author: 李映飞
 @date:   2022/4/29
 @note:
**/
package models

import (
	"fmt"
	"time"
	"webDesign/pkg/e"
)

type Faculty struct {
	// 学院ID
	FID uint `json:"fid" gorm:"primaryKey;autoIncrement"`
	// 学院代码
	FCODE string `json:"fcode" gorm:"unique;not null;type:char(2)"`
	// 学院名称
	Name string `json:"name" gorm:"not null;varchar(10)"`
}

type Department struct {
	// 系ID
	DID uint `json:"did" gorm:"primaryKey;autoIncrement"`
	// 归属学院ID
	FID uint `json:"fid" gorm:"not null"`
	// 系代码
	DCODE string `json:"dcode" gorm:"unique;not null;type:char(2)"`
	// 系名称
	Name string `json:"name" gorm:"not null;varchar(10)"`
}

type Class struct {
	// 班级ID
	SID uint `json:"sid" gorm:"primaryKey;autoIncrement"`
	// 归属系ID
	DID uint `json:"did" gorm:"not null"`
	// 所属年级
	Session string `json:"session" gorm:"not null;type:char(2)"`
	// 班级代码
	SCODE string `json:"fcode" gorm:"not null;type:char(2)"`
	// 班级名称
	Name string `json:"name" gorm:"not null;varchar(10)"`
}

type Student struct {
	// 学号
	Number   string    `json:"number" gorm:"primaryKey;type:char(10)"`
	Name     string    `json:"name" gorm:"not null;type:varchar(6)"`
	Sex      string    `json:"sex" gorm:"not null;char(2)"`
	Birthday time.Time `json:"birthday" gorm:"not null;type:date"`
	// 归属班级ID
	SID uint `json:"sid" gorm:"not null"`
}

type ViewStudent struct {
	Number   string `json:"number"`
	Name     string `json:"name"`
	Sex      string `json:"sex"`
	Birthday string `json:"birthday"`
	Class    string `json:"class"`
}

func GetFaculty(FID uint) Faculty {
	var faculty Faculty

	if err := db.First(&faculty, FID).Error; err != nil {
		return Faculty{}
	} else {
		return faculty
	}
}

func GetDepartment(DID uint) Department {
	var department Department

	if err := db.First(&department, DID).Error; err != nil {
		return Department{}
	} else {
		return department
	}
}

func GetClass(name string) Class {
	c := Class{}
	if err := db.Debug().Where("name = ?", name).First(&c).Error; err == nil {
		return c
	} else {
		return Class{}
	}
}

func GetClassName(SID uint) string {
	var class Class
	if err := db.Select("name").First(&class, SID).Error; err != nil {
		return ""
	} else {
		return class.Name
	}
}

// GetPrefix 生成学号前缀
func GetPrefix(class Class) string {
	prefix := ""
	department := GetDepartment(class.DID)
	faculty := GetFaculty(department.FID)

	if faculty.FID == 0 {
		return prefix
	} else {
		prefix = class.Session + faculty.FCODE + department.DCODE + class.SCODE
		return prefix
	}
}

// JoinStudent 添加学生
func JoinStudent(viewStudent ViewStudent) int {
	// 找到班级ID
	c := GetClass(viewStudent.Class)
	birthday, _ := time.ParseInLocation("2006-01-02", viewStudent.Birthday, time.Local)
	number := ""
	fmt.Println(c)
	if c.SID > 0 {
		// 生成学号前缀
		number = GetPrefix(c)
	} else {
		return e.ERROR_NOT_EXIST_CLASS
	}
	var cnt uint
	db.Raw("select count(*) from students where s_id = ?", c.SID).Scan(&cnt)
	number = fmt.Sprintf("%s%02d", number, cnt+1)
	stu := Student{
		Number:   number,
		Name:     viewStudent.Name,
		Sex:      viewStudent.Sex,
		Birthday: birthday,
		SID:      c.SID,
	}

	if err := db.Create(&stu).Error; err != nil {

		// 获取班级最后一个学生
		var lastStu Student
		db.Last(&lastStu).Where("SID = ?", c.SID)

		str := lastStu.Number[8:]
		fmt.Println(str)
		fmt.Sscanf(str, "%02d", &cnt)
		cnt = cnt + 1
		stu.Number = fmt.Sprintf("%s%02d", lastStu.Number[0:8], cnt)

		fmt.Println(stu.Number)
		if err := db.Create(&stu).Error; err != nil {
			return e.ERROR
		}
	}
	return e.SUCCESS
}

// FindStudent 查询学生信息
func FindStudent(number string) ViewStudent {
	var stu Student
	if err := db.First(&stu, "number = ?", number).Error; err != nil {
		return ViewStudent{}
	} else {
		return ViewStudent{
			Number:   stu.Number,
			Name:     stu.Name,
			Sex:      stu.Sex,
			Birthday: stu.Birthday.Format("2006-01-02"),
			Class:    GetClassName(stu.SID),
		}
	}
}

// UpdateStudent 更新学生信息
func UpdateStudent(viewStudent ViewStudent) int {
	var stu Student
	if err := db.First(&stu, "number = ?", viewStudent.Number).Error; err != nil {
		return e.ERROR_NOT_EXIST_NUMBER
	}
	birthday, _ := time.ParseInLocation("2006-01-02", viewStudent.Birthday, time.Local)
	SID := GetClass(viewStudent.Class).SID
	if SID > 0 {
		if err := db.Model(&stu).Updates(Student{
			Name:     viewStudent.Name,
			Sex:      viewStudent.Sex,
			Birthday: birthday,
			SID:      SID,
		}).Error; err != nil {
			return e.ERROR
		}
	} else {
		return e.ERROR_NOT_EXIST_CLASS
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
