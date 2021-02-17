package router

import (
	"gin-vue-admin/api/v1"
	"github.com/gin-gonic/gin"
)

func InitGlobalRouter(Router *gin.RouterGroup) {
	GlobalRouter := Router.Group("board")
	{
		GlobalRouter.POST("createRecord", v1.CreateRecord) // 新建留言板记录
		//GlobalRouter.DELETE("deleteBoats", v1.DeleteBoats)           // 删除Boats
		//GlobalRouter.DELETE("deleteBoatsByIds", v1.DeleteBoatsByIds) // 批量删除Boats
		GlobalRouter.PUT("updateRecord", v1.UpdateBoard) // 更新Boats
		//GlobalRouter.GET("findBoats", v1.FindBoats)                  // 根据ID获取Boats
		GlobalRouter.GET("getRecord", v1.GetRecord)   // 获取记录
		GlobalRouter.GET("delRecord", v1.DeleteBoard) // 删除留言
	}
}
