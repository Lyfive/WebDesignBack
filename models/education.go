/**
 @author: 李映飞
 @date:   2022/6/6
 @note:
**/
package models

// 对学院、专业、班级、学生的增删改查
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

type Session struct {
	// 学期ID
	DID uint `json:"did" gorm:"primaryKey;autoIncrement"`
	// 学期名称
	Session string `json:"session" gorm:"primaryKey;char(2)"`
}

type Class struct {
	// 班级ID
	SID uint `json:"sid" gorm:"primaryKey;autoIncrement"`
	// 归属系ID
	DID uint `json:"did" gorm:"not null"`
	// 所属年级
	Session string `json:"session" gorm:"not null;type:char(2)"`
	// 班级代码
	SCODE string `json:"scode" gorm:"not null;type:char(2)"`
	// 班级名称
	Name string `json:"name" gorm:"not null;varchar(10)"`
}

type Course struct {
	// 课程ID
	CID uint `json:"cid" gorm:"primaryKey;autoIncrement"`
	// 课程名
	Title string `json:"title" gorm:"not null;varchar(10)"`
}

type DC struct {
	DID uint `json:"did" gorm:"column:d_id"`
	CID uint `json:"cid" gorm:"column:c_id"`
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

func GetAllCourses() []Course {
	var courses []Course
	db.Find(&courses)
	return courses
}
func GetFacultyByID(FID uint) Faculty {
	var faculty Faculty

	if err := db.First(&faculty, FID).Error; err != nil {
		return Faculty{}
	} else {
		return faculty
	}
}

func GetDepartmentID(number string) uint {
	var student Student
	db.First(&student, number)
	return GetClassByID(student.SID).DID
}

func GetDepartmentByID(DID uint) Department {
	var department Department

	if err := db.First(&department, DID).Error; err != nil {
		return Department{}
	} else {
		return department
	}
}

func GetClass(number string) Class {
	c := Class{}
	var stu Student
	if err := db.Debug().Where("number = ?", number).First(&stu).Error; err == nil {
		c = GetClassByID(stu.SID)
		return c
	}
	return Class{}
}

func GetClassByID(SID uint) Class {
	c := Class{}
	if err := db.Debug().Where("s_id = ?", SID).First(&c).Error; err == nil {
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

func AddFaculty(faculty Faculty) uint {
	if err := db.Omit("f_id").Create(&faculty).Error; err != nil {
		return 0
	} else {
		return faculty.FID
	}
}

func DeleteFaculty(FID uint) error {
	if err := db.Delete(&Faculty{}, FID).Error; err != nil {
		return err
	} else {
		return nil
	}
}

func AddDepartment(department Department) uint {
	if err := db.Omit("d_id").Create(&department).Error; err != nil {
		return 0
	} else {
		return department.DID
	}
}

func DeleteDepartment(DID uint) error {
	if err := db.Delete(&Department{}, DID).Error; err != nil {
		return err
	} else {
		return nil
	}
}

func AddSession(session Session) error {
	if err := db.Create(&session).Error; err != nil {
		return err
	} else {
		return nil
	}
}

func DeleteSession(session Session) error {
	if err := db.Delete(&session).Error; err != nil {
		return err
	} else {
		return nil
	}
}

func AddClass(class *Class) error {
	if err := db.Omit("s_id").Create(&class).Error; err != nil {
		return err
	} else {
		return nil
	}
}

func DeleteClass(SID uint) error {
	if err := db.Delete(&Class{}, SID).Error; err != nil {
		return err
	} else {
		return nil
	}
}

func AddCourse(course *Course) error {
	if err := db.Omit("c_id").Create(&course).Error; err != nil {
		return err
	} else {
		return nil
	}
}

func DeleteCourse(CID uint) error {
	if err := db.Delete(&Course{}, CID).Error; err != nil {
		return err
	} else {
		return nil
	}
}

func AddDC(dc *DC) error {
	if err := db.Table("dc").Create(&dc).Error; err != nil {
		return err
	} else {
		return nil
	}
}

func DeleteDC(dc *DC) error {
	if err := db.Table("dc").Where("d_id = ? and c_id = ?", dc.DID, dc.CID).Delete(&dc).Error; err != nil {
		return err
	} else {
		return nil
	}
}
