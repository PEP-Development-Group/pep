// 自动生成模板Boats
package model

import (
	"gin-vue-admin/global"
)

// 如果含有time.Time 请自行import time包
type Board struct {
	global.GVA_MODEL
	Msg string `json:"msg" form:"msg" gorm:"column:msg;comment:留言消息;type:varchar(60)"`
}

func (Board) TableName() string {
	return "board"
}
