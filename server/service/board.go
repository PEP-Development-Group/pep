package service

import (
	"errors"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

//@author: [sh1luo](https://github.com/sh1luo)
//@function: CreateBoard
//@description: 创建留言记录
//@param: boats model.Boats
//@return: err error

func CreateBoard(board model.Board) (err error) {
	//err = global.GVA_DB.Create(&board).Error
	//if err != nil {
	//	return err
	//}
	//future, _ := time.ParseInLocation("2006-01-02 15:04:05", "2040-07-07 09:00:00", time.Local)
	//return global.GVA_REDIS.ZAdd("board", redis.Z{Score: float64(future.Unix() - time.Now().Unix()), Member: board.Msg}).Err()
	global.GVA_REDIS.Set("board", board.Msg, 0)
	result, err := global.GVA_REDIS.Get("board").Result()
	if err != nil || result != board.Msg {
		return err
	}
	return nil
}

//@author: [sh1luo](https://github.com/sh1luo)
//@function: GetBoardRecords
//@description: 获取留言
//@param: ""
//@return: string

func GetBoardRecords() string {
	msg, _ := global.GVA_REDIS.Get("board").Result()
	return msg
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteBoats
//@description: 删除Boats记录
//@param: boats model.Boats
//@return: err error

func DeleteBoats() error {
	global.GVA_REDIS.Del("board")
	if global.GVA_REDIS.Get("board").String() != "" {
		return errors.New("删除失败")
	}
	return nil
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateBoats
//@description: 更新Boats记录
//@param: boats *model.Boats
//@return: err error

func UpdateBoats(board model.Board) (err error) {
	global.GVA_REDIS.Set("board", board.Msg, 0)
	result, err := global.GVA_REDIS.Get("board").Result()
	if err != nil || result != board.Msg {
		return errors.New("更新失败")
	}
	return nil
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetBoatsInfoList
//@description: 分页获取Boats记录
//@param: info request.BoatsSearch
//@return: err error, list interface{}, total int64

//func GetBoatsInfoList(info request.BoatsSearch) (err error, list interface{}, total int64) {
//	limit := info.PageSize
//	offset := info.PageSize * (info.Page - 1)
//	// 创建db
//	db := global.GVA_DB.Model(&model.Boats{})
//	var boatss []model.Boats
//	// 如果有条件搜索 下方会自动创建搜索语句
//	err = db.Count(&total).Error
//	err = db.Limit(limit).Offset(offset).Find(&boatss).Error
//	return err, boatss, total
//}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteBoatsByIds
//@description: 批量删除Boats记录
//@param: ids request.IdsReq
//@return: err error

func DeleteBoatsByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.Board{}, "id in ?", ids.Ids).Error
	return err
}
