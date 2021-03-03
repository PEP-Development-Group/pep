package router

import (
	"pep/api/v1"
	"github.com/gin-gonic/gin"
)

func InitGlobalRouter(Router *gin.RouterGroup) {
	GlobalRouter := Router.Group("board")
	{
		GlobalRouter.POST("createRecord", v1.CreateRecord) // 新建留言板记录	1
		GlobalRouter.PUT("updateRecord", v1.UpdateBoard)   // 更新Boats	1
		GlobalRouter.DELETE("delRecord", v1.DeleteBoard)   // 删除留言	1

		GlobalRouter.DELETE("delAll", v1.InitDB) // 初始化数据库到学期初	1
	}
}

func InitUserBoard(Router *gin.RouterGroup) {
	GlobalRouter := Router.Group("board")
	{
		GlobalRouter.GET("getRecord", v1.GetRecord) // 获取记录	1
	}
}
