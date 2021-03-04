package information

import (
	"github.com/gookit/color"
	"gorm.io/gorm"
	"pep/global"
	"pep/model"
)

var Admin = new(admin)

type admin struct{}

var admins = []model.SysUser{
	//{GVA_MODEL: global.GVA_MODEL{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()}, UUID: uuid.NewV4(), Username: 88888888, Password: "e10adc3949ba59abbe56e057f20f883e", Name: "管理员", AuthorityId: "888"},
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: sys_users 表数据初始化
func (a *admin) Init() error {
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if tx.Where("id IN ?", []int{1, 2}).Find(&[]model.SysUser{}).RowsAffected == 2 {
			color.Danger.Println("\n[Mysql] --> sys_users 表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&admins).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		color.Info.Println("\n[Mysql] --> sys_users 表初始数据成功!")
		return nil
	})
}
