package router

import (
	"gin-vue-admin/api/v1"
	"github.com/gin-gonic/gin"
)

func InitClassRouter(Router *gin.RouterGroup) {
	ClassRouter := Router.Group("class")
	{
		ClassRouter.POST("createClass", v1.CreateClass)             // 新建Class
		ClassRouter.DELETE("deleteClass", v1.DeleteClass)           // 删除Class
		ClassRouter.DELETE("deleteClassByIds", v1.DeleteClassByIds) // 批量删除Class
		ClassRouter.PUT("updateClass", v1.UpdateClass)              // 更新Class
		ClassRouter.GET("findClass", v1.FindClass)                  // 根据ID获取Class
		ClassRouter.PATCH("addUserCancelNums", v1.AddUserCancelNums)             // 增加某人取消次数

		ClassRouter.GET("getClassList", v1.GetStuClassList)          // 学生选课Class列表（树状api）
		ClassRouter.GET("getClass", v1.GetClassList)                 // Class管理列表
		ClassRouter.GET("getPersonalClasses", v1.GetPersonalClasses) // 学生获取个人课表

		ClassRouter.GET("getTeacherClassList", v1.GetTeacherClassList)         // 教师获取个人课表
		ClassRouter.GET("getTeacherAClassStuList", v1.GetTeacherAClassStuList) // 教师某课所有学生情况
		ClassRouter.PATCH("setStuGrade", v1.SetStuGrade)                       // 修改学生成绩（只能一次）

		ClassRouter.POST("selectClass", v1.SelectClass)     // 学生选课
		ClassRouter.DELETE("deleteSelect", v1.DeleteSelect) // 学生退选
	}
}
