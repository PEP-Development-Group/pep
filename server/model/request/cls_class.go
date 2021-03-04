package request

import (
	"pep/model"
)

type ClassSearch struct {
	model.Class
	PageInfo
}

type UsernameRequest struct {
	Username int64 `json:"username" form:"username"`
}

type TeacherRequest struct {
	UsernameRequest
	Cid   uint `json:"cid"`
	Grade uint `json:"grade"`
}

type SelectClass struct {
	UsernameRequest `json:"ur" form:"ur"`
	Cid             uint `json:"cid" form:"cid"`
}
