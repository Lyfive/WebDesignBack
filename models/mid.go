/*
*

	@author: 李映飞
	@date:   2022/5/22
	@note:

*
*/
package models

type ViewFaculty struct {
	// 学院ID
	FID uint `json:"fid" gorm:"primaryKey;autoIncrement"`
	// 学院名称
	Name string `json:"name" gorm:"not null;varchar(10)"`
}

func GetAllFaculties() ([]*ViewFaculty, error) {
	faculties := make([]*ViewFaculty, 0)
	err := db.Model(&Faculty{}).Find(&faculties).Error
	return faculties, err
}

type ViewDepartment struct {
	// 系ID
	DID uint `json:"did" gorm:"primaryKey;autoIncrement"`

	// 系名称
	Name string `json:"name" gorm:"not null;varchar(10)"`
}

func GetAllDepartments(fid uint) ([]*ViewDepartment, error) {
	departments := make([]*ViewDepartment, 0)
	err := db.Model(&Department{}).Where("f_id = ?", fid).Find(&departments).Error
	return departments, err
}

type ViewSession struct {
	Session string `json:"session" gorm:"char(2)"`
}

func GetAllSessions(did uint) ([]*ViewSession, error) {
	sessions := make([]*ViewSession, 0)
	err := db.Model(&Session{}).Where("d_id = ?", did).Find(&sessions).Error
	return sessions, err
}

type ViewClass struct {
	// 班级ID
	SID uint `json:"sid"`

	// 班级名称
	Name string `json:"name" gorm:"not null;"`
}

func GetAllClasses(did uint, session string) ([]*ViewClass, error) {
	classes := make([]*ViewClass, 0)
	err := db.Model(&Class{}).Where("d_id = ? and session = ?", did, session).Find(&classes).Error
	return classes, err
}

type ViewStudentMessage struct {
	Number string `json:"number"`
	Name   string `json:"name"`
}

func GetStudentsByClass(sid uint) ([]*ViewStudentMessage, error) {
	students := make([]*ViewStudentMessage, 0)
	err := db.Model(&Student{}).Where("s_id = ?", sid).Find(&students).Error
	return students, err
}

func GetCoursesByDepartment(did uint) ([]Course, error) {
	courses := make([]Course, 0)
	err := db.Raw("select courses.c_id c_id ,courses.title title from courses,dcs where courses.c_id = dcs.c_id and dcs.d_id = ?", did).Scan(&courses).Error
	return courses, err
}
