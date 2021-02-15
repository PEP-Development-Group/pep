package request

import (
	"gin-vue-admin/model"
)

type ClassSearch struct {
	model.Class
	PageInfo
}

type UsernameRequest struct {
	Username string `json:"username" form:"username"`
}

type TeacherRequest struct {
	Username string `json:"username"`
	Tname    string `json:"tname"`
	Cid      uint   `json:"cid"`
	Grade    uint   `json:"grade"`
}

type SelectClass struct {
	Username string `json:"username" form:"username"`
	Cid      uint   `json:"cid" form:"cid"`
}
