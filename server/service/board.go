package service

import (
	"errors"
	"pep/global"
	"pep/model"
)

//@author: [sh1luo](https://github.com/sh1luo)
//@function: CreateBoard
//@description: 创建留言记录
//@param: board model.Board
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

//@author: [sh1luo](https://github.com/sh1luo)
//@function: DeleteBoats
//@description: 删除Boats记录
//@param: ""
//@return: err error

func DeleteBoats() error {
	global.GVA_REDIS.Del("board")
	if global.GVA_REDIS.Get("board").String() != "" {
		return errors.New("删除失败")
	}
	return nil
}

//@author: [sh1luo](https://github.com/sh1luo)
//@function: UpdateBoard
//@description: 更新Board记录
//@param: board model.Board
//@return: err error

func UpdateBoats(board model.Board) (err error) {
	global.GVA_REDIS.Set("board", board.Msg, 0)
	result, err := global.GVA_REDIS.Get("board").Result()
	if err != nil || result != board.Msg {
		return errors.New("更新失败")
	}
	return nil
}
