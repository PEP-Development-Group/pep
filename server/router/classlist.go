package router

import (
	"gin-vue-admin/api/v1"
	"github.com/gin-gonic/gin"
)

func InitClassListRouter(Router *gin.RouterGroup) {
	ClassListRouter := Router.Group("classlist")
	{
		ClassListRouter.POST("createClassList", v1.CreateClassList)             // 新建ClassList
		ClassListRouter.DELETE("deleteClassList", v1.DeleteClassList)           // 删除ClassList
		ClassListRouter.DELETE("deleteClassListByIds", v1.DeleteClassListByIds) // 批量删除ClassList
		ClassListRouter.PUT("updateClassList", v1.UpdateClassList)              // 更新ClassList
		ClassListRouter.GET("findClassList", v1.FindClassList)                  // 根据ID获取ClassList
		ClassListRouter.GET("getClassListList", v1.GetClassListList)            // 获取ClassList列表
	}
}
