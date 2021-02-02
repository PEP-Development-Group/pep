// 自动生成模板Stu
package model

import (
	"gin-vue-admin/global"
)

// 如果含有time.Time 请自行import time包
type Stu struct {
      global.GVA_MODEL
      StuId  int `json:"stuId" form:"stuId" gorm:"column:stu_id;comment:学号;type:bigint;size:20;"`
      StuName  string `json:"stuName" form:"stuName" gorm:"column:stu_name;comment:姓名;type:varchar(5);size:5;"`
      College  string `json:"college" form:"college" gorm:"column:college;comment:学院;type:varchar(5);size:5;"`
      Major  string `json:"major" form:"major" gorm:"column:major;comment:专业;type:varchar(10);size:10;"`
      PID  string `json:"PID" form:"PID" gorm:"column:PID;comment:身份证;type:char;"`
      Hash  string `json:"hash" form:"hash" gorm:"column:hash;comment:hash密码;type:char;"`
      Salt  string `json:"salt" form:"salt" gorm:"column:salt;comment:盐;type:char;"`
      CancelNums  int `json:"cancelNums" form:"cancelNums" gorm:"column:cancel_nums;comment:取消次数;type:int;size:10;"`

      Authority   SysAuthority `json:"authority" gorm:"foreignKey:AuthorityId;references:AuthorityId;comment:用户角色"`
      AuthorityId string       `json:"authorityId" gorm:"default:888;comment:用户角色ID"`
}


func (Stu) TableName() string {
  return "stu"
}


// 如果使用工作流功能 需要打开下方注释 并到initialize的workflow中进行注册 且必须指定TableName
// type StuWorkflow struct {
// 	// 工作流操作结构体
// 	WorkflowBase      `json:"wf"`
// 	Stu   `json:"business"`
// }

// func (Stu) TableName() string {
// 	return "stu"
// }

// 工作流注册代码

// initWorkflowModel内部注册
// model.WorkflowBusinessStruct["stu"] = func() model.GVA_Workflow {
//   return new(model.StuWorkflow)
// }

// initWorkflowTable内部注册
// model.WorkflowBusinessTable["stu"] = func() interface{} {
// 	return new(model.Stu)
// }
