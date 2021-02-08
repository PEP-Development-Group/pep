// 自动生成模板Boats
package model

import (
	"gin-vue-admin/global"
)

// 如果含有time.Time 请自行import time包
type Boats struct {
      global.GVA_MODEL
      Bid  int `json:"bid" form:"bid" gorm:"column:bid;comment:;type:int;size:10;"`
      Bname  string `json:"bname" form:"bname" gorm:"column:bname;comment:;type:varchar(255);size:255;"`
      Color  string `json:"color" form:"color" gorm:"column:color;comment:;type:varchar(255);size:255;"`
      Sex  *bool `json:"sex" form:"sex" gorm:"column:sex;comment:;type:tinyint;"`
}


func (Boats) TableName() string {
  return "boats"
}


// 如果使用工作流功能 需要打开下方注释 并到initialize的workflow中进行注册 且必须指定TableName
// type BoatsWorkflow struct {
// 	// 工作流操作结构体
// 	WorkflowBase      `json:"wf"`
// 	Boats   `json:"business"`
// }

// func (Boats) TableName() string {
// 	return "boats"
// }

// 工作流注册代码

// initWorkflowModel内部注册
// model.WorkflowBusinessStruct["boats"] = func() model.GVA_Workflow {
//   return new(model.BoatsWorkflow)
// }

// initWorkflowTable内部注册
// model.WorkflowBusinessTable["boats"] = func() interface{} {
// 	return new(model.Boats)
// }
