package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateStu
//@description: 创建Stu记录
//@param: stu model.Stu
//@return: err error

func CreateStu(stu model.Stu) (err error) {
	err = global.GVA_DB.Create(&stu).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteStu
//@description: 删除Stu记录
//@param: stu model.Stu
//@return: err error

func DeleteStu(stu model.Stu) (err error) {
	err = global.GVA_DB.Delete(&stu).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteStuByIds
//@description: 批量删除Stu记录
//@param: ids request.IdsReq
//@return: err error

func DeleteStuByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.Stu{},"id in ?",ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateStu
//@description: 更新Stu记录
//@param: stu *model.Stu
//@return: err error

func UpdateStu(stu model.Stu) (err error) {
	err = global.GVA_DB.Save(&stu).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetStu
//@description: 根据id获取Stu记录
//@param: id uint
//@return: err error, stu model.Stu

func GetStu(id uint) (err error, stu model.Stu) {
	err = global.GVA_DB.Where("id = ?", id).First(&stu).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetStuInfoList
//@description: 分页获取Stu记录
//@param: info request.StuSearch
//@return: err error, list interface{}, total int64

func GetStuInfoList(info request.StuSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.Stu{})
    var stus []model.Stu
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&stus).Error
	return err, stus, total
}