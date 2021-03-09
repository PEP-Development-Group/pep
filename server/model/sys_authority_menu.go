package model

type SysMenu struct {
	SysBaseMenu
	MenuId      string                 `json:"menuId" gorm:"comment:菜单ID"`
	AuthorityId string                 `json:"-" gorm:"comment:角色ID"`
	Children    []SysMenu              `json:"children,omitempty" gorm:"-"`
	Parameters  []SysBaseMenuParameter `json:"parameters,omitempty" gorm:"foreignKey:SysBaseMenuID;references:MenuId"`
}

func (s SysMenu) TableName() string {
	return "authority_menu"
}
