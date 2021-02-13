package service

import (
	"errors"
	"fmt"
	"gin-vue-admin/constant"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	"gin-vue-admin/model/response"
	"gorm.io/gorm"
)

//@author: [sh1luo](https://github.com/sh1luo)
//@function: CreateClass
//@description: 创建选课记录
//@param: class request.SelectClass
//@return: err error

func SelectClass(sc request.SelectClass) (err error) {
	cls := model.Class{}
	err = global.GVA_DB.Select("selected", "total", "id").Where("id = ?", sc.Cid).First(&cls).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return constant.ErrClassNotExist
	}

	if cls.Selected >= cls.Total {
		return constant.ErrClassHasFull
	}

	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		err = tx.Model(&cls).Update("selected", cls.Selected+1).Error
		if err != nil {
			return err
		}
		scm := model.SelectClass{
			Cid:      sc.Cid,
			Username: sc.Username,
		}
		return tx.Create(&scm).Error
	})
}

//@author: [sh1luo](https://github.com/sh1luo)
//@function: DeleteSelect
//@description: 退课
//@param: class request.SelectClass
//@return: err error

func DeleteSelect(sc request.SelectClass) (err error) {
	user := model.SysUser{}

	err = global.GVA_DB.Select("CancelNums").Where("username = ?", sc.Username).First(&user).Error
	if err != nil {
		fmt.Print("err1")
		return err
	}

	if user.CancelNums >= global.GVA_CONFIG.System.CancelNums {
		return constant.ErrDelClassTooMany
	}

	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		err = tx.Model(&user).Where("username = ?", sc.Username).Update("CancelNums", user.CancelNums+1).Error
		if err != nil {
			fmt.Print("err2")
			return err
		}

		cls := model.Class{}
		db := tx.Select("selected").Where("id = ?", sc.Cid)
		_ = db.First(&cls).Error
		err = db.Update("selected", cls.Selected-1).Error
		if err != nil {
			return err
		}

		return tx.Where("class_id = ? AND username = ?", sc.Cid, sc.Username).Delete(&model.SelectClass{}).Error
	})
}

//@author: [sh1luo](https://github.com/sh1luo)
//@function: GetPersonalClasses
//@description: 获取个人选课列表
//@param: class request.SelectClass
//@return: err error

func GetPersonalClasses(rq request.ClassList) (cls []model.Class, err error) {
	// 获取所有选课记录的课程id
	var scm []model.SelectClass
	global.GVA_DB.Select("cid").Where("username = ?", rq.Username).Find(&scm)
	return
}

//@author: [piexlmax](https//github.com/piexlmax)
//@function: CreateClass
//@description: 创建Class记录
//@param: class model.Class
//@return: err error

func CreateClass(class model.Class) (err error) {
	err = global.GVA_DB.Create(&class).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteClass
//@description: 删除Class记录
//@param: class model.Class
//@return: err error

func DeleteClass(class model.Class) (err error) {
	err = global.GVA_DB.Delete(&class).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteClassByIds
//@description: 批量删除Class记录
//@param: ids request.IdsReq
//@return: err error

func DeleteClassByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.Class{}, "id in ?", ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateClass
//@description: 更新Class记录
//@param: class *model.Class
//@return: err error

func UpdateClass(class model.Class) (err error) {
	err = global.GVA_DB.Save(&class).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetClass
//@description: 根据id获取Class记录
//@param: id uint
//@return: err error, class model.Class

func GetClass(id uint) (err error, class model.Class) {
	err = global.GVA_DB.Where("id = ?", id).First(&class).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetClassInfoList
//@description: 分页获取Class记录
//@param: info request.ClassSearch
//@return: err error, list interface{}, total int64

func GetClassInfoList(info request.ClassSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.Class{})
	var classs []model.Class
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Ccredit != 0 {
		db = db.Where("`ccredit` = ?", info.Ccredit)
	}
	if info.Cname != "" {
		db = db.Where("`cname` LIKE ?", "%"+info.Cname+"%")
	}
	if info.Tname != "" {
		db = db.Where("`tname` LIKE ?", "%"+info.Tname+"%")
	}
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&classs).Error
	return err, classs, total
}

//func GetClassInfoListWithPerson(rq request.ClassList) (err error, list interface{}, total int) {
//	var cls []model.Class
//	global.GVA_DB.Model(&model.Class{}).Order("cname").Find(&cls)
//
//	l := len(cls)
//	m := make(map[uint]bool, l)
//
//	// 获取所有选课的课程id
//	if l > 0 {
//		var scs []model.SelectClass
//		global.GVA_DB.Select("cid").Where("username = ?", rq.Username).Find(&scs)
//		for _, sc :=  range scs {
//			m[sc.Cid] = true
//		}
//	}
//
//	// 同名分类
//	var classes response.ClassListResponse
//	curGroup := &response.Group{}
//	if l > 0 {
//		curGroup.ID = 0
//		curGroup.ClassName = cls[0].Cname
//		curGroup.Hours = cls[0].Ccredit
//	}
//	idx := 1
//	for i := 0; i < l; i++ {
//		if i > 0 && cls[i].Cname != cls[i-1].Cname {
//			classes.G = append(classes.G, *curGroup) // 加入之前的分组
//			curGroup = &response.Group{}             // 创建新分组
//			curGroup.ID = idx
//			idx++
//			curGroup.ClassName = cls[i].Cname
//			curGroup.Hours = cls[i].Ccredit
//		}
//		c := response.Course{
//			ID: cls[i].ID, TeacherName: cls[i].Tname, Time: cls[i].Time,
//			Desc: cls[i].Desc, ClassRoom: cls[i].Classroom, Max: cls[i].Total,
//			Now: cls[i].Selected,
//		}
//		if m[cls[i].ID] {
//			c.Selected = true
//		}
//
//		curGroup.List = append(curGroup.List, c)
//	}
//
//	return err, classes, len(classes.G)
//}

func GetClassInfoListWithPerson(rq request.ClassList) (err error, list interface{}, total int) {
	var cls []model.Class
	global.GVA_DB.Find(&cls)

	l := len(cls)
	m := make(map[uint]bool, l)

	var classes response.ClassListResponse
	if l > 0 {
		// 用于后续判断该用户是否选了该课程
		var scs []model.SelectClass
		global.GVA_DB.Select("class_id").Where("username = ?", rq.Username).Find(&scs)
		for _, sc := range scs {
			m[sc.Cid] = true
		}

		// 同名分类
		classes.G = make(map[string]*response.Group, 15)
		gidx := 0
		for _, c := range cls {
			if _, ok := classes.G[c.Cname]; !ok {
				classes.G[c.Cname] = &response.Group{}
				classes.G[c.Cname].ID = gidx
				gidx++
				classes.G[c.Cname].Hours = c.Ccredit
				//classes.G[c.Cname].ClassName = c.Cname
			}
			co := response.Course{
				ID:          c.ID,
				TeacherName: c.Tname,
				Time:        c.Time,
				Desc:        c.Desc,
				ClassRoom:   c.Classroom,
				Max:         c.Total,
				Now:         c.Selected,
			}
			if m[c.ID] {
				co.Selected = true
			}
			classes.G[c.Cname].List = append(classes.G[c.Cname].List, co)

		}
	}

	return err, classes.G, len(classes.G)
}
