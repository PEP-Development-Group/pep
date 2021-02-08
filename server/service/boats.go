package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateBoats
//@description: 创建Boats记录
//@param: boats model.Boats
//@return: err error

func CreateBoats(boats model.Boats) (err error) {
	err = global.GVA_DB.Create(&boats).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteBoats
//@description: 删除Boats记录
//@param: boats model.Boats
//@return: err error

func DeleteBoats(boats model.Boats) (err error) {
	err = global.GVA_DB.Delete(&boats).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteBoatsByIds
//@description: 批量删除Boats记录
//@param: ids request.IdsReq
//@return: err error

func DeleteBoatsByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.Boats{},"id in ?",ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateBoats
//@description: 更新Boats记录
//@param: boats *model.Boats
//@return: err error

func UpdateBoats(boats model.Boats) (err error) {
	err = global.GVA_DB.Save(&boats).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetBoats
//@description: 根据id获取Boats记录
//@param: id uint
//@return: err error, boats model.Boats

func GetBoats(id uint) (err error, boats model.Boats) {
	err = global.GVA_DB.Where("id = ?", id).First(&boats).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetBoatsInfoList
//@description: 分页获取Boats记录
//@param: info request.BoatsSearch
//@return: err error, list interface{}, total int64

func GetBoatsInfoList(info request.BoatsSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.Boats{})
    var boatss []model.Boats
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&boatss).Error
	return err, boatss, total
}