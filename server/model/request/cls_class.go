package request

import (
	"gin-vue-admin/model"
)

type ClassSearch struct {
	model.Class
	PageInfo
}

type SelectClass struct {
	Username string `json:"username" form:"username"`
	Cid      int    `json:"cid" form:"cid"`
}
