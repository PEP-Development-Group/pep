package request

import (
	"gin-vue-admin/model"
)

type ClassSearch struct {
	model.Class
	PageInfo
}

type ClassListSearch struct {

}

type SelectClass struct {
	Username string `json:"username" form:"username"`
	Cid      uint    `json:"cid" form:"cid"`
}
