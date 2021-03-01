// 自动生成模板Class
package model

import (
	"gin-vue-admin/global"
	"time"
)

// 如果含有time.Time 请自行import time包
type Class struct {
	global.GVA_MODEL
	Ccredit int       `json:"ccredit" form:"ccredit" gorm:"column:ccredit;comment:学分;type:int"`
	Cname   string    `json:"cname" form:"cname" gorm:"column:cname;comment:课程名;type:varchar(20);"`
	Desc    string    `json:"desc" form:"desc" gorm:"column:desc;comment:上课时间中文描述;type:varchar(15);"`
	Etime   time.Time `json:"etime" form:"etime" gorm:"column:etime;comment:选课结束;type:datetime;"`
	Stime   time.Time `json:"stime" form:"stime" gorm:"column:stime;comment:选课开始;type:datetime;"`
	// Time      time.Time `json:"time" form:"time" gorm:"column:time;comment:上课时间;type:datetime;"`
	Tname     string `json:"tname" form:"tname" gorm:"column:tname;comment:教师名;type:varchar(5);"`
	Selected  int    `json:"selected" form:"selected" gorm:"column:selected;comment:已选人数;type:int;default:0"`
	Total     int    `json:"total" form:"total" gorm:"column:total;comment:总人数;type:int;default:0"`
	Classroom string `json:"classroom" form:"classroom" gorm:"type:varchar(10);comment:教室"`
}

type SelectClass struct {
	global.GVA_MODEL
	Username string `json:"username" gorm:"column:username;comment:学号;type:varchar(20);"`
	Cid      uint   `json:"cid" gorm:"column:class_id;comment:课程id;type:tinyint;"`
	Grade    uint   `json:"grade" gorm:"column:grade;default:102;comment:成绩/101旷课/102未上成绩;type:tinyint;"`
}

func (Class) TableName() string {
	return "cls_class"
}

func (SelectClass) TableName() string {
	return "user_classes"
}
