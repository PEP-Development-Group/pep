package v1

import (
	"gin-vue-admin/constant"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	"gin-vue-admin/model/response"
	"gin-vue-admin/service"
	"gin-vue-admin/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// @Tags Class
// @Summary 选课
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.SelectClass true "选课"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"选课成功"}"
// @Router /class/selectClass [post]
func SelectClass(c *gin.Context) {
	var class request.SelectClass
	_ = c.ShouldBindJSON(&class)
	if err := utils.Verify(class, utils.StudentVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := service.SelectClass(class); err != nil {
		global.GVA_LOG.Error("选课失败!", zap.Any("err", err))
		if err == constant.ErrClassNotExist {
			response.FailWithMessage("课程不存在", c)
		} else {
			response.FailWithMessage("选课人数已满", c)
		}
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
	var class request.SelectClass
	_ = c.ShouldBindJSON(&class)
	if err := utils.Verify(class, utils.StudentVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := service.DeleteSelect(class); err != nil {
		global.GVA_LOG.Error("退选失败!", zap.Any("err", err))
		if err == constant.ErrDelClassTooMany {
			response.FailWithMessage("退课次数太多", c)
		} else {
			response.FailWithMessage("未知原因，退课失败", c)
		}
	} else {
		response.OkWithMessage("退选成功", c)
	}
}

// @Tags Class
// @Summary 获取个人已选课程
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Class true "获取ClassList"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /class/getMyClass [get]
func GetPersonalClasses(c *gin.Context) {
	var class request.ClassList
	_ = c.ShouldBindJSON(&class)
	if err := utils.Verify(class, utils.StudentVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if _, err := service.GetPersonalClasses(class); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithMessage("获取成功", c)
	}
}

// @Tags Class
// @Summary 创建Class
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Class true "创建Class"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /class/createClass [post]
func CreateClass(c *gin.Context) {
	var class model.Class
	_ = c.ShouldBindJSON(&class)
	if err := service.CreateClass(class); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags Class
// @Summary 删除Class
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Class true "删除Class"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /class/deleteClass [delete]
func DeleteClass(c *gin.Context) {
	var class model.Class
	_ = c.ShouldBindJSON(&class)
	if err := service.DeleteClass(class); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Class
// @Summary 批量删除Class
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Class"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /class/deleteClassByIds [delete]
func DeleteClassByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteClassByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags Class
// @Summary 更新Class
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Class true "更新Class"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /class/updateClass [put]
func UpdateClass(c *gin.Context) {
	var class model.Class
	_ = c.ShouldBindJSON(&class)
	if err := service.UpdateClass(class); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags Class
// @Summary 用id查询Class
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Class true "用id查询Class"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /class/findClass [get]
func FindClass(c *gin.Context) {
	var class model.Class
	_ = c.ShouldBindQuery(&class)
	if err, reclass := service.GetClass(class.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reclass": reclass}, c)
	}
}

// @Tags Class
// @Summary 分页获取Class列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.ClassSearch true "分页获取Class列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /class/getClassList [get]
func GetClassList(c *gin.Context) {
	var pageInfo request.ClassSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetClassInfoList(pageInfo); err != nil {
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

// @Tags Class
// @Summary 分页获取Class列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.ClassSearch true "分页获取Class列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /class/GetClassListWithPerson [get]
func GetClassListWithPerson(c *gin.Context) {
	var class request.ClassList
	_ = c.ShouldBindJSON(&class)
	if err := utils.Verify(class, utils.StudentVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err, list, total := service.GetClassInfoListWithPerson(class); err != nil {
		global.GVA_LOG.Error("获取失败", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithData(gin.H{"courses": list,"total":total}, c)
	}
}