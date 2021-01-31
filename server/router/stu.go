package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitStuRouter(Router *gin.RouterGroup) {
	StuRouter := Router.Group("stu").Use(middleware.OperationRecord())
	{
		StuRouter.POST("createStu", v1.CreateStu)   // 新建Stu
		StuRouter.DELETE("deleteStu", v1.DeleteStu) // 删除Stu
		StuRouter.DELETE("deleteStuByIds", v1.DeleteStuByIds) // 批量删除Stu
		StuRouter.PUT("updateStu", v1.UpdateStu)    // 更新Stu
		StuRouter.GET("findStu", v1.FindStu)        // 根据ID获取Stu
		StuRouter.GET("getStuList", v1.GetStuList)  // 获取Stu列表
	}
}
