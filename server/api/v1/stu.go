package v1

import (
	"gin-vue-admin/global"
    "gin-vue-admin/model"
    "gin-vue-admin/model/request"
    "gin-vue-admin/model/response"
    "gin-vue-admin/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

// @Tags Stu
// @Summary 创建Stu
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Stu true "创建Stu"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /stu/createStu [post]
func CreateStu(c *gin.Context) {
	var stu model.Stu
	_ = c.ShouldBindJSON(&stu)
	if err := service.CreateStu(stu); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags Stu
// @Summary 删除Stu
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Stu true "删除Stu"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /stu/deleteStu [delete]
func DeleteStu(c *gin.Context) {
	var stu model.Stu
	_ = c.ShouldBindJSON(&stu)
	if err := service.DeleteStu(stu); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Stu
// @Summary 批量删除Stu
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Stu"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /stu/deleteStuByIds [delete]
func DeleteStuByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteStuByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags Stu
// @Summary 更新Stu
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Stu true "更新Stu"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /stu/updateStu [put]
func UpdateStu(c *gin.Context) {
	var stu model.Stu
	_ = c.ShouldBindJSON(&stu)
	if err := service.UpdateStu(stu); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags Stu
// @Summary 用id查询Stu
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Stu true "用id查询Stu"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /stu/findStu [get]
func FindStu(c *gin.Context) {
	var stu model.Stu
	_ = c.ShouldBindQuery(&stu)
	if err, restu := service.GetStu(stu.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"restu": restu}, c)
	}
}

// @Tags Stu
// @Summary 分页获取Stu列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.StuSearch true "分页获取Stu列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /stu/getStuList [get]
func GetStuList(c *gin.Context) {
	var pageInfo request.StuSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetStuInfoList(pageInfo); err != nil {
	    global.GVA_LOG.Error("获取失败", zap.Any("err", err))
        response.FailWithMessage("获取失败", c)
    } else {
        response.OkWithDetailed(response.PageResult{
            List:     list,
            Total:    total,
            Page:     pageInfo.Page,
            PageSize: pageInfo.PageSize,
        }, "获取成功", c)
    }
}
