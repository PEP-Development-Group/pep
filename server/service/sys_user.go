package service

import (
	"errors"
	"pep/global"
	"pep/model"
	"pep/model/request"
	"pep/utils"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

//@author: [sh1luo](https://github.com/sh1luo)
//@function: Register
//@description: 用户注册
//@param: u model.SysUser
//@return: err error, userInter model.SysUser

var authority map[string]model.SysAuthority

func init() {
	authority = make(map[string]model.SysAuthority, 3)
	authority["888"] = model.SysAuthority{
		AuthorityId:   "888",
		AuthorityName: "管理员",
		ParentId:      "0",
		DefaultRouter: "dashboard",
	}
	authority["1"] = model.SysAuthority{
		AuthorityId:   "1",
		AuthorityName: "同学",
		ParentId:      "0",
		DefaultRouter: "dashboard",
	}
	authority["2"] = model.SysAuthority{
		AuthorityId:   "2",
		AuthorityName: "老师",
		ParentId:      "0",
		DefaultRouter: "dashboard",
	}
}

func Register(u model.SysUser) (err error, userInter model.SysUser) {
	var user model.SysUser
	if !errors.Is(global.GVA_DB.Where("username = ?", u.Username).First(&user).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
		return errors.New("用户名已注册"), userInter
	}
	u.Password = utils.MD5V([]byte(u.Password))
	u.UUID = uuid.NewV4()
	err = global.GVA_DB.Create(&u).Error
	return err, u
}

//@author: [sh1luo](https://github.com/sh1luo)
//@function: Login
//@description: 用户登录
//@param: u *model.SysUser
//@return: err error, userInter *model.SysUser

func Login(u *model.SysUser) (err error, userInter *model.SysUser) {
	var user model.SysUser
	up := utils.MD5V([]byte(u.Password))
	err = global.GVA_DB.Select("username","AuthorityId","password", "uuid", "name", "class", "have_credits", "total_credits", "selected_credits", "cancel_nums").First(&user, u.Username).Error
	if err != nil {
		return err, nil
	}
	if up != user.Password {
		return errors.New("密码错误"), nil
	}
	user.Authority = authority[user.AuthorityId]
	return err, &user
}

//@author: [sh1luo](https://github.com/sh1luo)
//@function: AddCancelNums
//@description: 增加一个用户的取消次数
//@param: ac request.AddCancelNums
//@return: err error

func AddCancelNums(ac request.AddCancelNums) (err error) {
	var u model.SysUser
	return global.GVA_DB.Select("cancel_nums").Where("username = ?", ac.Username).First(&u).Update("cancel_nums", u.CancelNums+ac.Cnt).Error
}

//@author: [sh1luo](https://github.com/sh1luo)
//@function: ChangePassword
//@description: 修改用户密码
//@param: u *model.SysUser, newPassword string
//@return: err error, userInter *model.SysUser

func ChangePassword(u *model.SysUser, newPassword string) (err error, userInter *model.SysUser) {
	var user model.SysUser
	u.Password = utils.MD5V([]byte(u.Password))
	err = global.GVA_DB.Where("username = ? AND password = ?", u.Username, u.Password).First(&user).Update("password", utils.MD5V([]byte(newPassword))).Error
	return err, u
}

//@author: [sh1luo](https://github.com/sh1luo)
//@function: GetUserInfoList
//@description: 分页获取数据
//@param: info request.PageInfo
//@return: err error, list interface{}, total int64

func GetUserInfoList(info request.PageInfo) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&model.SysUser{})
	var userList []model.SysUser
	db.Where("authority_id = ? ", info.Param).Count(&total)
	err = db.Limit(limit).Offset(offset).Preload("Authority").Select("username", "name", "cancel_nums", "class", "have_credits", "total_credits").Where("authority_id = ? ", info.Param).Find(&userList).Error
	return err, userList, total
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: SetUserAuthority
//@description: 设置一个用户的权限
//@param: uuid uuid.UUID, authorityId string
//@return: err error

func SetUserAuthority(uuid uuid.UUID, authorityId string) (err error) {
	err = global.GVA_DB.Where("uuid = ?", uuid).First(&model.SysUser{}).Update("authority_id", authorityId).Error
	return err
}

//@author: [sh1luo](https://github.com/sh1luo)
//@function: DeleteUser
//@description: 删除用户
//@param: id float64
//@return: err error

func DeleteUser(username int64) error {
	err := global.GVA_DB.Where("username = ?", username).Delete(&model.SysUser{}).Error
	return err
}

//@author: [sh1luo](https://github.com/sh1luo)
//@function: SetUserInfo
//@description: 设置用户信息
//@param: reqUser model.SysUser
//@return: err error, user model.SysUser

func SetUserInfo(reqUser model.SysUser) (err error, user model.SysUser) {
	if reqUser.Password != "" {
		reqUser.Password = utils.MD5V([]byte(reqUser.Password))
	}
	err = global.GVA_DB.Where("username = ?", reqUser.Username).Updates(&reqUser).Error
	return err, reqUser
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@function: FindUserById
//@description: 通过id获取用户信息
//@param: id int
//@return: err error, user *model.SysUser

func FindUserById(username int) (err error, user *model.SysUser) {
	var u model.SysUser
	err = global.GVA_DB.Where("`username` = ?", username).First(&u).Error
	return err, &u
}
