package model

import (
	"gin-vue-admin/global"
	"github.com/satori/go.uuid"
)

type SysUser struct {
	global.GVA_MODEL
	UUID        uuid.UUID    `json:"uuid" gorm:"comment:用户UUID"`
	Username    string       `json:"username" gorm:"comment:用户登录名"`
	Password    string       `json:"-"  gorm:"comment:用户登录密码"`
	Name        string       `json:"name" gorm:"default:系统用户;comment:用户昵称" `
	// HeaderImg   string       `json:"headerImg" gorm:"default:http://qmplusimg.henrongyi.top/head.png;comment:用户头像"`
	Authority   SysAuthority `json:"authority" gorm:"foreignKey:AuthorityId;references:AuthorityId;comment:用户角色"`
	AuthorityId string       `json:"authorityId" gorm:"default:888;comment:用户角色ID"`

	College string `json:"college"`
	Major   string `json:"major"`
	PID     string `json:"pid"`
}
