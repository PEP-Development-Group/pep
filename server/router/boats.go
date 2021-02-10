package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitBoatsRouter(Router *gin.RouterGroup) {
	BoatsRouter := Router.Group("boats").Use(middleware.OperationRecord())
	{
		BoatsRouter.POST("createBoats", v1.CreateBoats)             // 新建Boats
		BoatsRouter.DELETE("deleteBoats", v1.DeleteBoats)           // 删除Boats
		BoatsRouter.DELETE("deleteBoatsByIds", v1.DeleteBoatsByIds) // 批量删除Boats
		BoatsRouter.PUT("updateBoats", v1.UpdateBoats)              // 更新Boats
		BoatsRouter.GET("findBoats", v1.FindBoats)                  // 根据ID获取Boats
		BoatsRouter.GET("getBoatsList", v1.GetBoatsList)            // 获取Boats列表
	}
}
