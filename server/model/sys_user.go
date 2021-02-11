package model

import (
	"gin-vue-admin/global"
	"github.com/satori/go.uuid"
)

type SysUser struct {
	global.GVA_MODEL
	UUID     uuid.UUID `json:"uuid" gorm:"comment:用户UUID"`
	Username string    `json:"username" gorm:"comment:学号/工号"`
	Password string    `json:"-" gorm:"type:varchar;comment:登录密码"`
	Name     string    `json:"name" gorm:"type:varchar;size:5;default:系统用户;comment:姓名"`

	Authority   SysAuthority `json:"authority" gorm:"foreignKey:AuthorityId;references:AuthorityId;comment:用户角色"`
	AuthorityId string       `json:"authorityId" gorm:"default:888;comment:用户角色ID"`

	CancelNums int    `json:"cancel_nums" gorm:"default:0;comment:取消次数"`
	Class      string `json:"class"`
	PID        string `json:"pid" gorm:"column:pid;type:char;comment:身份证号"`

	HaveCredits  int `json:"have_credits" gorm:"comment:已修学时"`
	TotalCredits int `json:"total_credits" gorm:"comment:应修学时"`

	Classes []*Class `gorm:"many2many:user_classes;foreignKey:Username;joinForeignKey:username;"`
}
