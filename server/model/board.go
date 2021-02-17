// 自动生成模板Boats
package model

// 如果含有time.Time 请自行import time包
type Board struct {
	Msg string `json:"msg" form:"msg"`
}

func (Board) TableName() string {
	return "board"
}
