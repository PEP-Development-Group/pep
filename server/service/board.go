package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	"github.com/go-redis/redis"
	"time"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateBoard
//@description: 创建Boats记录
//@param: boats model.Boats
//@return: err error

func CreateBoard(board model.Board) (err error) {
	err = global.GVA_DB.Create(&board).Error
	if err != nil {
		return err
	}
	future, _ := time.ParseInLocation("2006-01-02 15:04:05", "2040-07-07 09:00:00", time.Local)
	return global.GVA_REDIS.ZAdd("board", redis.Z{Score: float64(future.Unix() - time.Now().Unix()), Member: board.Msg}).Err()
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetBoardRecords
//@description: 获取留言列表
//@param: ""
//@return: err error, boats model.Boats

func GetBoardRecords() (resp []string, err error) {
	resp, err = global.GVA_REDIS.ZRange("board", 0,-1).Result()
	return
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
//@function: DeleteBoats
//@description: 删除Boats记录
//@param: boats model.Boats
//@return: err error

func DeleteBoats(boats model.Board) (err error) {
	err = global.GVA_DB.Delete(&boats).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteBoatsByIds
//@description: 批量删除Boats记录
//@param: ids request.IdsReq
//@return: err error

func DeleteBoatsByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.Board{}, "id in ?", ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateBoats
//@description: 更新Boats记录
//@param: boats *model.Boats
//@return: err error

func UpdateBoats(boats model.Board) (err error) {
	err = global.GVA_DB.Save(&boats).Error
	return err
}
