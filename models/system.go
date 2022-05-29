/**
 @author: 李映飞
 @date:   2022/5/22
 @note:
**/
package models

import (
	"fmt"
	"sync"
	"time"
)

var mutex sync.Mutex

type System struct {
	CreateTime   time.Time `json:"createTime";gorm:"column:create_time;primary_key"`
	Version      string    `json:"version";gorm:"column:version"`
	VisitsNumber uint      `json:"visitsNumber";gorm:"column:visits_number"`
}

type Population struct {
	FacultyNumber    uint `json:"facultyNumber";gorm:"column:faculty_number"`
	DepartmentNumber uint `json:"departmentNumber";gorm:"column:department_number"`
	ClassNumber      uint `json:"classNumber";gorm:"column:class_number"`
	CourseNumber     uint `json:"courseNumber";gorm:"column:course_number"`
	StudentNumber    uint `json:"studentNumber";gorm:"column:student_number"`
	GradeNumber      uint `json:"gradeNumber";gorm:"column:grade_number"`
}

func GetSystemMessage() (System, error) {
	var system System
	err := db.First(&system).Error
	return system, err
}

func UpdateVisitsNumber() error {
	var system System
	mutex.Lock()
	defer mutex.Unlock()
	err := db.Debug().First(&system).Error
	fmt.Println(system)
	if err != nil {
		return err
	}
	system.VisitsNumber++
	fmt.Println(system)
	err = db.Debug().Model(&system).Where("create_time = ?", system.CreateTime).Update("visits_number", system.VisitsNumber).Error
	return err
}

func GetPopulationMessage() Population {
	var population Population
	db.Raw("Select count(*) as faculty_number from faculties").Scan(&population.FacultyNumber)
	db.Raw("Select count(*) as department_number from departments").Scan(&population.DepartmentNumber)
	db.Raw("Select count(*) as class_number from classes").Scan(&population.ClassNumber)
	db.Raw("Select count(*) as course_number from courses").Scan(&population.CourseNumber)
	db.Raw("Select count(*) as student_number from students").Scan(&population.StudentNumber)
	db.Raw("Select count(*) as grade_number from grades").Scan(&population.GradeNumber)
	return population
}
