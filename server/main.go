package main

import (
	"gin-vue-admin/core"
	"gin-vue-admin/global"
	"gin-vue-admin/initialize"
)

// @title Swagger Example API
// @version 0.0.1
// @description This is a sample Server pets
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name x-token
// @BasePath /
func main() {
	// TODO:feature: 管理员留言消息
	// TODO:optimize: 建立索引
	// TODO:只能查询本人信息，验证x-token
	// TODO:username改为int64

	global.GVA_VP = core.Viper()          // 初始化Viper
	global.GVA_LOG = core.Zap()           // 初始化zap日志库
	global.GVA_DB = initialize.Gorm()     // gorm连接数据库
	initialize.MysqlTables(global.GVA_DB) // 初始化表
	// 程序结束前关闭数据库链接
	db, _ := global.GVA_DB.DB()
	defer db.Close()

	core.RunWindowsServer()
}
