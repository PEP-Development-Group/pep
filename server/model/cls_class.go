// 自动生成模板Class
package model

import (
	"gin-vue-admin/global"
	"time"
)

// 如果含有time.Time 请自行import time包
type Class struct {
	global.GVA_MODEL
	Ccredit int       `json:"ccredit" form:"ccredit" gorm:"column:ccredit;comment:学分;type:int;size:10;"`
	Cname   string    `json:"cname" form:"cname" gorm:"column:cname;comment:课程名;type:varchar(20);size:20;"`
	Etime   time.Time `json:"etime" form:"etime" gorm:"column:etime;comment:选课结束;type:datetime;"`
	Stime   time.Time `json:"stime" form:"stime" gorm:"column:stime;comment:选课开始;type:datetime;"`
	Time    time.Time `json:"time" form:"time" gorm:"column:time;comment:上课时间;type:datetime;"`
	Tname   string    `json:"tname" form:"tname" gorm:"column:tname;comment:教师名;type:varchar(5);size:5;"`
	Seleted int       `json:"seleted" form:"seleted" gorm:"column:seleted;comment:已选人数;type:int;size:10;"`
	Total   int       `json:"total" form:"total" gorm:"column:total;comment:总人数;type:int;size:10;"`
}

func (Class) TableName() string {
	return "cls_class"
}

// 如果使用工作流功能 需要打开下方注释 并到initialize的workflow中进行注册 且必须指定TableName
// type ClassWorkflow struct {
// 	// 工作流操作结构体
// 	WorkflowBase      `json:"wf"`
// 	Class   `json:"business"`
// }

// func (Class) TableName() string {
// 	return "cls_class"
// }

// 工作流注册代码

// initWorkflowModel内部注册
// model.WorkflowBusinessStruct["class"] = func() model.GVA_Workflow {
//   return new(model.ClassWorkflow)
// }

// initWorkflowTable内部注册
// model.WorkflowBusinessTable["class"] = func() interface{} {
// 	return new(model.Class)
// }
