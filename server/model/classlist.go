// 自动生成模板ClassList
package model

import (
	"gin-vue-admin/global"
)

// 如果含有time.Time 请自行import time包
type ClassList struct {
	global.GVA_MODEL
	Cname     string `json:"cname" form:"cname" gorm:"column:cname;comment:'课程名';type:varchar(20)"`
	Ccredit   int    `json:"ccredit" form:"ccredit" gorm:"column:ccredit;comment:'课程学时';type:tinyint"`
	Classroom string `json:"classroom" form:"classroom" gorm:"column:classroom;comment:'上课地点';type:varchar(10)"`
}

func (ClassList) TableName() string {
	return "class_list"
}
