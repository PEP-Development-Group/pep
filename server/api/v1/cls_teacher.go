package v1

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model/request"
	"gin-vue-admin/model/response"
	"gin-vue-admin/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// @Tags Class
// @Summary 教师获取自己的上课列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.UsernameRequest true "教师获取自己的上课列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /class/GetTeacherClassList [get]
func GetTeacherClassList(c *gin.Context) {
	var class request.UsernameRequest
	_ = c.ShouldBindJSON(&class)

	if err, list, total := service.GetTeacherClassList(class); err != nil {
		global.GVA_LOG.Error("获取失败", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithData(gin.H{"courses": list, "total": total}, c)
	}
}

// @Tags Class
// @Summary 教师查看自己某节课的学生列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.TeacherRequest true "教师查看自己某节课的学生列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /class/GetTeacherAClassStuList [get]
func GetTeacherAClassStuList(c *gin.Context) {
	var class request.TeacherRequest
	_ = c.ShouldBindJSON(&class)

	if err, list, total := service.GetTeacherAClassStuList(class); err != nil {
		global.GVA_LOG.Error("获取失败", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithData(gin.H{"students": list, "total": total}, c)
	}
}

// @Tags Class
// @Summary 教师上某学生某课成绩
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.TeacherRequest true "教师上某学生某课成绩"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"提交成功"}"
// @Router /class/SetStuGrade [patch]
func SetStuGrade(c *gin.Context) {
	// grade, username ,cid
	var class request.TeacherRequest
	_ = c.ShouldBindJSON(&class)

	if err := service.SetStuGrade(class); err != nil {
		global.GVA_LOG.Error("提交失败", zap.Any("err", err))
		response.FailWithMessage("提交失败", c)
	} else {
		response.OkWithMessage("提交成功", c)
	}
}
