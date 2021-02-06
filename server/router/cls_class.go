package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitClassRouter(Router *gin.RouterGroup) {
	ClassRouter := Router.Group("class").Use(middleware.OperationRecord())
	{
		ClassRouter.POST("createClass", v1.CreateClass)   // 新建Class
		ClassRouter.DELETE("deleteClass", v1.DeleteClass) // 删除Class
		ClassRouter.DELETE("deleteClassByIds", v1.DeleteClassByIds) // 批量删除Class
		ClassRouter.PUT("updateClass", v1.UpdateClass)    // 更新Class
		ClassRouter.GET("findClass", v1.FindClass)        // 根据ID获取Class
		ClassRouter.GET("getClassList", v1.GetClassList)  // 获取Class列表
	}
}
