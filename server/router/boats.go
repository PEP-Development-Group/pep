package router

import (
	"gin-vue-admin/api/v1"
	"github.com/gin-gonic/gin"
)

func InitGlobalRouter(Router *gin.RouterGroup) {
	GlobalRouter := Router.Group("board")
	{
		GlobalRouter.POST("createBoardRecord", v1.CreateRecord)       // 新建留言板记录
		//GlobalRouter.DELETE("deleteBoats", v1.DeleteBoats)           // 删除Boats
		//GlobalRouter.DELETE("deleteBoatsByIds", v1.DeleteBoatsByIds) // 批量删除Boats
		//GlobalRouter.PUT("updateBoats", v1.UpdateBoats)              // 更新Boats
		//GlobalRouter.GET("findBoats", v1.FindBoats)                  // 根据ID获取Boats
		GlobalRouter.GET("getRecordList", v1.GetRecordsList)            // 获取Boats列表
	}
}
