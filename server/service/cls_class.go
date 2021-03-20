package service

import (
	"database/sql"
	"errors"
	"gorm.io/gorm"
	"pep/constant"
	"pep/global"
	"pep/model"
	"pep/model/request"
	"pep/model/response"
	"strconv"
	"strings"
	"time"
)

func GetUserCreditsInfo(username int64) (selected int, have int, cancel int, err error) {
	var u model.SysUser
	err = global.GVA_DB.Select("have_credits", "cancel_nums", "selected_credits").Where("username = ?", username).First(&u).Error
	return u.SelectedCredits, u.HaveCredits, u.CancelNums, err
}

//@author: [sh1luo](https://github.com/sh1luo)
//@function: CreateClass
//@description: 创建选课记录
//@param: class request.SelectClass
//@return: err error

func SelectClass(sc request.SelectClass) (err error) {
	// 课程人数已满不能再选
	cls := model.Class{}
	err = global.GVA_DB.Select("id", "cname", "selected", "total", "ccredit").Where("id = ?", sc.Cid).First(&cls).Error
	if err == gorm.ErrRecordNotFound {
		return constant.ErrClassNotExist
	}
	//if cls.Selected >= cls.Total {
	//	return constant.ErrClassHasFull
	//}

	// 已修够或者超选不可继续选课
	u := model.SysUser{}
	global.GVA_DB.Select("selected_credits", "have_credits", "total_credits").Where("username = ?", sc.Username).First(&u)
	if u.HaveCredits >= u.TotalCredits || u.SelectedCredits+cls.Ccredit > u.TotalCredits {
		return constant.ErrSelectFull
	}

	// 同名课程只能选择一个
	var cname string
	var rows *sql.Rows
	rows, err = global.GVA_DB.Raw("select distinct cname from cls_class where id IN (select class_id from user_classes where username = ? AND deleted_at IS NULL) group by cname", sc.Username).Rows()
	if err != nil {
		return constant.InternalErr
	}
	defer rows.Close()
	for rows.Next() {
		_ = rows.Scan(&cname)
		if cname == cls.Cname {
			return constant.ErrClassNameSame
		}
	}

	// 每周只能选 global.GVA_CONFIG.System.WeeklyEnrollLimit 次课
	cntMap := make(map[uint8]uint8, 25)  // key 是周数，value是该周选课次数
	desc := ""
	rows, err = global.GVA_DB.Raw("select c.desc from cls_class c, user_classes sc where username = ? AND c.id=sc.class_id", sc.Username).Rows()
	if err != nil {
		return constant.InternalErr
	}
	defer rows.Close()
	for rows.Next() {
		_ = rows.Scan(&desc)
		ts := strings.Split(desc, "-")
		week, err := strconv.Atoi(ts[0])
		if err != nil {
			return constant.ErrClassDesc
		}
		if cntMap[uint8(week)]+1 == global.GVA_CONFIG.System.WeeklyEnrollLimit {
			return constant.ErrSelectClassTooManyWeekly
		}
		cntMap[uint8(week)]++
	}

	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		err = tx.Model(&model.SysUser{}).Where("username = ?", sc.Username).UpdateColumn("selected_credits", u.SelectedCredits+cls.Ccredit).Error
		if err != nil {
			return err
		}

		err = tx.Model(&model.Class{}).Where("id = ? and selected < ?", cls.ID, cls.Total).Update("selected", cls.Selected+1).Error
		if err != nil {
			return err
		}
		scm := model.SelectClass{}
		// 这里不需要再判断是否选过了，已经在 同名课程只能选一次 中判断过
		//err = tx.Where("class_id = ? AND username = ?", sc.Cid, sc.Username).First(&scm).Error
		//if err == gorm.ErrRecordNotFound {
		//	return constant.ErrClassHasSelected
		//}
		scm.Cid = sc.Cid
		scm.Username = sc.Username
		scm.Grade = 102
		return tx.Create(&scm).Error
	})
}

//@author: [sh1luo](https://github.com/sh1luo)
//@function: DeleteSelect
//@description: 退课
//@param: class request.SelectClass
//@return: err error

func DeleteSelect(sc request.SelectClass) (err error) {
	// 退课次数为0不允许再退课
	user := model.SysUser{}
	global.GVA_DB.Select("CancelNums", "SelectedCredits").Where("username = ?", sc.Username).First(&user)
	if user.CancelNums == 0 {
		return constant.ErrDelClassTooMany
	}

	srs := model.SelectClass{}
	err = global.GVA_DB.Select("grade").Where("class_id = ? AND username = ?", sc.Cid, sc.Username).First(&srs).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return constant.ErrClassHasNotSelected
	}
	// 成绩有变动不允许再退课
	if 	srs.Grade != 102 {
		return constant.ErrDelClassAfterGrade
	}

	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		err = global.GVA_DB.Where("class_id = ? AND username = ?", sc.Cid, sc.Username).Unscoped().Delete(&model.SelectClass{}).Error
		if err != nil {
			return constant.ErrDelClass
		}

		err = tx.Model(&user).Where("username = ?", sc.Username).Update("CancelNums", user.CancelNums-1).Error
		if err != nil {
			return err
		}

		// TODO:去掉order by
		c := model.Class{}
		tmptx2 := tx.Select("selected", "desc", "ccredit").First(&c, sc.Cid)

		// desc split by "-", such as "1-1-1"
		ts := strings.Split(c.Desc, "-")
		week, err := strconv.Atoi(ts[0])
		d, err := strconv.Atoi(ts[1])
		if err != nil {
			return err
		}

		// 上课当天及以后不允许退课
		t, _ := time.ParseInLocation("2006-01-02", global.GVA_CONFIG.System.FirstDay, time.Local)
		if time.Now().After(t.AddDate(0, 0, week*7+d)) {
			return constant.ErrDelClassAfterClass
		}

		err = tmptx2.Update("selected", c.Selected-1).Error
		if err != nil {
			return constant.ErrDelClass
		}

		err = tx.Model(&model.SysUser{}).Where("username = ?", sc.Username).UpdateColumn("selected_credits", user.SelectedCredits-c.Ccredit).Error
		if err != nil {
			return constant.ErrDelClass
		}

		return nil
	})
}

//@author: [sh1luo](https://github.com/sh1luo)
//@function: GetPersonalClasses
//@description: 获取个人选课列表
//@param: class request.SelectClass
//@return: err error

func GetPersonalClasses(rq request.UsernameRequest) (resp response.PersonalClassResponse, total int, err error) {
	var scm []model.SelectClass
	global.GVA_DB.Select("class_id", "grade").Where("username = ?", rq.Username).Find(&scm)
	if len(scm) == 0 {
		return
	}

	var ids []uint
	m := make(map[uint]uint)
	for _, sc := range scm {
		ids = append(ids, sc.Cid)
		m[sc.Cid] = sc.Grade
	}

	var cls []model.Class
	global.GVA_DB.Select("id", "cname", "ccredit", "tname", "desc", "classroom").Find(&cls, ids)

	for _, c := range cls {
		if g, ok := m[c.ID]; ok {
			resp.Crs = append(resp.Crs, response.ClassRecord{
				ID:        c.ID,
				Cname:     c.Cname,
				Hours:     c.Ccredit,
				Tname:     c.Tname,
				Desc:      c.Desc,
				Classroom: c.Classroom,
				Grade:     g,
			})
		}
	}

	total = len(resp.Crs)
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
	// 直接从db中删除
	err = global.GVA_DB.Unscoped().Delete(&class).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteClassByIds
//@description: 批量删除Class记录
//@param: ids request.IdsReq
//@return: err error

func DeleteClassByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Unscoped().Delete(&[]model.Class{}, "id in ?", ids.Ids).Error
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

func GetStuClassList(rq request.UsernameRequest) (err error, list interface{}, total int) {
	// TODO:SQL调优?
	var clsAll []model.Class
	global.GVA_DB.Select("id", "ccredit", "cname", "tname", "desc", "classroom", "total", "selected", "etime", "stime").Find(&clsAll)
	var cls []model.Class
	for _, class := range clsAll {
		if class.Stime.Before(time.Now()) && class.Etime.After(time.Now()) {
			cls = append(cls, class)
		}
	}
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
			}
			co := response.Course{
				ID:          c.ID,
				TeacherName: c.Tname,
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

		// 对已修过的课程进行标识
		var cname string
		var rows *sql.Rows
		// 找出用户选过的所有课程的课程名
		rows, err = global.GVA_DB.Raw("select distinct cname from cls_class where id IN (select class_id from user_classes where username = ? AND deleted_at IS NULL) group by cname", rq.Username).Rows()
		if err != nil {
			return constant.InternalErr, nil, 0
		}
		defer rows.Close()
		for rows.Next() {
			_ = rows.Scan(&cname)
			if _, ok := classes.G[cname]; ok {
				classes.G[cname].Learned = true
			}
		}
	}

	return err, classes.G, len(classes.G)
}

func GetTeacherClassList(rq request.UsernameRequest) (err error, list interface{}, total int) {
	var t model.SysUser
	err = global.GVA_DB.Select("name").First(&t, "username = ?", rq.Username).Error
	if err != nil {
		return
	}

	var cls []model.Class
	err = global.GVA_DB.Select("id", "cname", "ccredit", "desc", "selected", "classroom").Find(&cls, "tname = ?", t.Name).Error
	if err != nil {
		return
	}

	var resp response.TeacherClassResponse
	for _, c := range cls {
		resp.Tcrs = append(resp.Tcrs, response.TeacherClassRecord{
			Cid:       c.ID,
			Cname:     c.Cname,
			Ccredit:   c.Ccredit,
			Desc:      c.Desc,
			Selected:  c.Selected,
			Classroom: c.Classroom,
		})
	}

	return nil, resp, len(resp.Tcrs)
}

func GetTeacherAClassStuList(cid uint) (err error, list interface{}, total int) {
	var scs []model.SelectClass
	err = global.GVA_DB.Select("username", "grade").Find(&scs, "class_id = ?", cid).Error
	if err != nil {
		return
	}

	var names []int64
	m := make(map[int64]uint)
	for _, s := range scs {
		names = append(names, s.Username)
		m[s.Username] = s.Grade
	}

	var users []model.SysUser
	err = global.GVA_DB.Select("username", "name", "class").Find(&users, "username IN ?", names).Error
	if err != nil {
		return
	}

	var resp response.TeacherClassStuResponse
	for _, u := range users {
		resp.Tcsrs = append(resp.Tcsrs, response.TeacherClassStuRecord{
			Username: u.Username,
			Name:     u.Name,
			Class:    u.Class,
			Grade:    m[u.Username],
		})
	}

	return nil, resp.Tcsrs, len(resp.Tcsrs)
}

func SetStuGrade(rq request.TeacherRequest) (err error) {
	var sc model.SelectClass
	db := global.GVA_DB.Select("grade").Where("class_id = ? AND username = ?", rq.Cid, rq.Username).First(&sc)
	if db.Error != nil {
		return constant.InternalErr
	}

	// 成绩>=60，才算已修学时
	if sc.Grade >= 60 {
		var s model.SysUser
		db := global.GVA_DB.Model(&model.SysUser{}).Select("have_credits").Where("username = ?", rq.Username)
		err = db.First(&s).Error
		if err != nil {
			return err
		}
		var c model.Class
		err = global.GVA_DB.Select("ccredit").Where("id = ?", rq.Cid).First(&c).Error
		if err != nil {
			return err
		}
		err = db.Update("have_credits", s.HaveCredits+c.Ccredit).Error
		if err != nil {
			return err
		}
	}
	return db.Update("grade", rq.Grade).Error
}
