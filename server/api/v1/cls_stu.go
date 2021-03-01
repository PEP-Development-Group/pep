package v1

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model/request"
	"gin-vue-admin/model/response"
	"gin-vue-admin/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 获取学生学分，取消次数信息
// @Router /class/getUserCreditInfo [get]
func GetUserCreditInfo(c *gin.Context) {
	username := getUsername(c)
	if have, cancel, err := service.GetUserCreditsInfo(username); err != nil {
		global.GVA_LOG.Error("获取学生信息失败!", zap.Any("err", err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithData(gin.H{
			"have_credits": have,
			"cancel_nums":  cancel,
		}, c)
	}
}

// @Tags Class
// @Summary 选课
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.SelectClass true "选课"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"选课成功"}"
// @Router /class/selectClass [post]
func SelectClass(c *gin.Context) {
	var sc request.SelectClass
	sc.Username = getUsername(c)
	_ = c.ShouldBindJSON(&sc)
	if err := service.SelectClass(sc); err != nil {
		global.GVA_LOG.Error("选课失败!", zap.Any("err", err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage("选课成功", c)
	}
}

// @Tags Class
// @Summary 退选Class
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.SelectClass true "退选Class"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"退课成功"}"
// @Router /class/deleteSelect [delete]
func DeleteSelect(c *gin.Context) {
	var sc request.SelectClass
	sc.Username = getUsername(c)
	_ = c.ShouldBindJSON(&sc)
	if err := service.DeleteSelect(sc); err != nil {
		global.GVA_LOG.Error("退选失败!", zap.Any("err", err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage("退选成功", c)
	}
}

// @Tags Class
// @Summary 学生选课课表（树状api）
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.UsernameRequest true "学生选课课表（树状api）"
// @Success 200 {string} string "{"success":true,"data":{""},"msg":"获取成功"}"
// @Router /class/GetStuClassList [get]
func GetStuClassList(c *gin.Context) {
	var u request.UsernameRequest
	u.Username = getUsername(c)
	if err, list, total := service.GetStuClassList(u); err != nil {
		global.GVA_LOG.Error("获取失败", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithData(gin.H{"courses": list, "total": total}, c)
	}
}

// @Tags Class
// @Summary 获取个人已选课程(课表，首页用)
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.UsernameRequest true "获取个人已选课程(课表，首页用)"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /class/getMyClass [get]
func GetPersonalClasses(c *gin.Context) {
	var u request.UsernameRequest
	u.Username = getUsername(c)
	if list, total, err := service.GetPersonalClasses(u); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithData(gin.H{"list": list, "total": total}, c)
	}
}
