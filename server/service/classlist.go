package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// @title    CreateClassList
// @description   create a ClassList
// @param     classlist               model.ClassList
// @auth                     （2020/04/05  20:22）
// @return    err             error

func CreateClassList(classlist model.ClassList) (err error) {
	err = global.GVA_DB.Create(&classlist).Error
	return err
}

// @title    DeleteClassList
// @description   delete a ClassList
// @auth                     （2020/04/05  20:22）
// @param     classlist               model.ClassList
// @return                    error

func DeleteClassList(classlist model.ClassList) (err error) {
	return global.GVA_DB.Where("id = ?", classlist.ID).Unscoped().Delete(&model.ClassList{}).Error
}

// @title    DeleteClassListByIds
// @description   delete ClassLists
// @auth                     （2020/04/05  20:22）
// @param     classlist               model.ClassList
// @return                    error

func DeleteClassListByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.ClassList{},"id in (?)",ids.Ids).Error
	return err
}

// @title    UpdateClassList
// @description   update a ClassList
// @param     classlist          *model.ClassList
// @auth                     （2020/04/05  20:22）
// @return                    error

func UpdateClassList(classlist *model.ClassList) (err error) {
	err = global.GVA_DB.Save(classlist).Error
	return err
}

// @title    GetClassList
// @description   get the info of a ClassList
// @auth                     （2020/04/05  20:22）
// @param     id              uint
// @return                    error
// @return    ClassList        ClassList

func GetClassList(id uint) (err error, classlist model.ClassList) {
	err = global.GVA_DB.Where("id = ?", id).First(&classlist).Error
	return
}

// @title    GetClassListInfoList
// @description   get ClassList list by pagination, 分页获取用户列表
// @auth                     （2020/04/05  20:22）
// @param     info            PageInfo
// @return                    error

func GetClassListInfoList() (err error, list interface{}) {
    var classlists []model.ClassList
	err = global.GVA_DB.Select("id","cname","ccredit","classroom").Find(&classlists).Error
	return err, classlists
}