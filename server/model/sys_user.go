package model

import (
	"gin-vue-admin/global"
	"github.com/satori/go.uuid"
)

type SysUser struct {
	global.GVA_MODEL
	UUID     uuid.UUID `json:"uuid" gorm:"comment:用户UUID"`
	Username string    `json:"username" gorm:"comment:学号/工号"`
	Password string    `json:"-" gorm:"comment:用户登录密码"`
	Name     string    `json:"name" gorm:"default:管理员;comment:真名" `
	// HeaderImg   string       `json:"headerImg" gorm:"default:http://qmplusimg.henrongyi.top/head.png;comment:用户头像"`
	Authority   SysAuthority `json:"authority" gorm:"foreignKey:AuthorityId;references:AuthorityId;comment:用户角色"`
	AuthorityId string       `json:"authorityId" gorm:"default:888;comment:用户角色ID"`

	CancelNums int    `json:"cancel_nums" gorm:"default:0;comment:取消次数"`
	College    string `json:"college"`
	Major      string `json:"major"`
	PID        string `json:"pid"`
}
