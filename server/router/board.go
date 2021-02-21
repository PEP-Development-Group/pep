package router

import (
	"gin-vue-admin/api/v1"
	"github.com/gin-gonic/gin"
)

func InitGlobalRouter(Router *gin.RouterGroup) {
	GlobalRouter := Router.Group("board")
	{
		GlobalRouter.POST("createRecord", v1.CreateRecord) // 新建留言板记录
		GlobalRouter.PUT("updateRecord", v1.UpdateBoard)   // 更新Boats
		GlobalRouter.DELETE("delRecord", v1.DeleteBoard)   // 删除留言
	}
}

func InitUserBoard(Router *gin.RouterGroup) {
	GlobalRouter := Router.Group("board")
	{
		GlobalRouter.GET("getRecord", v1.GetRecord)        // 获取记录
	}
}