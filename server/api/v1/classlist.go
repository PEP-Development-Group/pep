package v1

import (
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	"gin-vue-admin/model/response"
	"gin-vue-admin/service"
	"github.com/gin-gonic/gin"
)

// @Tags ClassList
// @Summary 创建ClassList
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ClassList true "创建ClassList"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /classlist/createClassList [post]
func CreateClassList(c *gin.Context) {
	var classlist model.ClassList
	_ = c.ShouldBindJSON(&classlist)
	err := service.CreateClassList(classlist)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags ClassList
// @Summary 删除ClassList
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ClassList true "删除ClassList"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /classlist/deleteClassList [delete]
func DeleteClassList(c *gin.Context) {
	var classlist model.ClassList
	_ = c.ShouldBindJSON(&classlist)
	err := service.DeleteClassList(classlist)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags ClassList
// @Summary 批量删除ClassList
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ClassList"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /classlist/deleteClassListByIds [delete]
func DeleteClassListByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	err := service.DeleteClassListByIds(IDS)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags ClassList
// @Summary 更新ClassList
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ClassList true "更新ClassList"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /classlist/updateClassList [put]
func UpdateClassList(c *gin.Context) {
	var classlist model.ClassList
	_ = c.ShouldBindJSON(&classlist)
	err := service.UpdateClassList(&classlist)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("更新失败，%v", err), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags ClassList
// @Summary 用id查询ClassList
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ClassList true "用id查询ClassList"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /classlist/findClassList [get]
func FindClassList(c *gin.Context) {
	var classlist model.ClassList
	_ = c.ShouldBindQuery(&classlist)
	err, reclasslist := service.GetClassList(classlist.ID)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("查询失败，%v", err), c)
	} else {
		response.OkWithData(gin.H{"reclasslist": reclasslist}, c)
	}
}

// @Tags ClassList
// @Summary 获取ClassList列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.ClassListSearch true "获取ClassList列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /classlist/getClassListList [get]
func GetClassListList(c *gin.Context) {
	err, list := service.GetClassListInfoList()
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取数据失败，%v", err), c)
	} else {
		var ts []model.SysUser
		global.GVA_DB.Select("name","username").Where("authority_id = ?", "2").Find(&ts)
		var tl []response.TeacherRecord
		for _, t := range ts {
			tl = append(tl, response.TeacherRecord{Name: t.Name,TID: t.Username})
		}
		response.OkWithData(gin.H{
			"classlist": list,
			"teachers": tl,
		}, c)
	}
}
