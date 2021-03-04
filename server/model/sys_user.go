package model

import (
	"github.com/satori/go.uuid"
)

type SysUser struct {
	//global.GVA_MODEL
	UUID     uuid.UUID `json:"uuid,omitempty" gorm:"comment:用户UUID"`
	Username int64     `json:"username,omitempty" gorm:"type:bigint(20);comment:学号/工号"`
	Password string    `json:"password,omitempty" gorm:"type:varchar(32);comment:登录密码"`
	Name     string    `json:"name,omitempty" gorm:"type:varchar(10);default:系统用户;comment:姓名"`

	Authority   SysAuthority `json:"authority" gorm:"foreignKey:AuthorityId;references:AuthorityId;comment:用户角色"`
	AuthorityId string       `json:"authorityId" gorm:"default:1;comment:用户角色ID"`

	CancelNums int    `json:"cancel_nums" gorm:"default:3;comment:取消次数"`
	Class      string `json:"class,omitempty" gorm:"varchar(10);comment:班级"`

	SelectedCredits int `json:"selected_credits" gorm:"default:0;comment:已选学时"`
	HaveCredits     int `json:"have_credits" gorm:"comment:已修学时"`
	TotalCredits    int `json:"total_credits" gorm:"comment:应修学时"`
}
