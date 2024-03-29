package v1

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"pep/global"
	"pep/model"
	"pep/model/response"
	"pep/service"
)

// @Tags Board
// @Summary 创建Board记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Boats true "创建Board记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"添加成功"}"
// @Router /board/createBoardRecord [post]
func CreateRecord(c *gin.Context) {
	var board model.Board
	_ = c.ShouldBindJSON(&board)
	if err := service.CreateBoard(board); err != nil {
		global.GVA_LOG.Error("新增留言失败!", zap.Any("err", err))
		response.FailWithMessage("新增留言失败", c)
	} else {
		response.OkWithMessage("新增留言成功", c)
	}
}

// @Tags Boats
// @Summary 获取留言
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.BoatsSearch true "获取留言"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /board/GetRecord [get]
func GetRecord(c *gin.Context) {
	response.OkWithData(service.GetBoardRecords(), c)
}

// @Tags Boats
// @Summary 删除留言
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Boats true "删除留言"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /board/deleteBoard [delete]
func DeleteBoard(c *gin.Context) {
	err := service.DeleteBoats()
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Boats
// @Summary 更新board
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Boats true "更新board"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /boats/updateBoard [put]
func UpdateBoard(c *gin.Context) {
	var board model.Board
	_ = c.ShouldBindJSON(&board)
	if err := service.UpdateBoats(board); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags Boats
// @Summary 批量删除Boats
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Boats"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /boats/deleteBoatsByIds [delete]
//func DeleteBoatsByIds(c *gin.Context) {
//	var IDS request.IdsReq
//	_ = c.ShouldBindJSON(&IDS)
//	if err := service.DeleteBoatsByIds(IDS); err != nil {
//		global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
//		response.FailWithMessage("批量删除失败", c)
//	} else {
//		response.OkWithMessage("批量删除成功", c)
//	}
//}

// @Tags Boats
// @Summary 用id查询Boats
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Boats true "用id查询Boats"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /boats/findBoats [get]
//func FindBoats(c *gin.Context) {
//	var boats model.Boats
//	_ = c.ShouldBindQuery(&boats)
//	if err, reboats := service.GetBoats(boats.ID); err != nil {
//		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
//		response.FailWithMessage("查询失败", c)
//	} else {
//		response.OkWithData(gin.H{"reboats": reboats}, c)
//	}
//}
