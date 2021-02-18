package model

import (
	"gin-vue-admin/global"
	"github.com/satori/go.uuid"
)

type SysUser struct {
	global.GVA_MODEL
	UUID     uuid.UUID `json:"uuid" gorm:"comment:用户UUID"`
	Username string    `json:"username" gorm:"type:varchar(20);comment:学号/工号"`
	Password string    `json:"-" gorm:"type:varchar(32);comment:登录密码"`
	Name     string    `json:"name" gorm:"type:varchar(10);default:系统用户;comment:姓名"`

	Authority   SysAuthority `json:"authority" gorm:"foreignKey:AuthorityId;references:AuthorityId;comment:用户角色"`
	AuthorityId string       `json:"authorityId" gorm:"default:888;comment:用户角色ID"`

	CancelNums int    `json:"cancel_nums" gorm:"default:3;comment:取消次数"`
	Class      string `json:"class" gorm:"varchar(10);comment:班级"`

	HaveCredits  int `json:"have_credits" gorm:"comment:已修学时"`
	TotalCredits int `json:"total_credits" gorm:"comment:应修学时"`
}
