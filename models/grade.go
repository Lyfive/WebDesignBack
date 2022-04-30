/**
 @author: 李映飞
 @date:   2022/4/29
 @note:
**/
package models

type Course struct {
	// 课程ID
	CID uint `json:"cid" gorm:"primaryKey;autoIncrement"`
	// 课程名
	Name string `json:"name" gorm:"not null;varchar(10)"`
}

type Grade struct {
	// 成绩ID
	GID uint `json:"gid" gorm:"primaryKey;autoIncrement"`
	// 绑定学生学号
	Number uint `json:"number" gorm:"not null"`
	// 绑定课程ID
	CID uint `json:"cid" gorm:"not null"`
	// 分数
	Mark uint `json:"mark" gorm:"not null;type:tinyint"`
}
