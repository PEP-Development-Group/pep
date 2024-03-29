package initialize

import (
	"github.com/gin-gonic/gin"
	_ "pep/docs"
	"pep/global"
	"pep/middleware"
	"pep/router"
)

// 初始化总路由

func Routers() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	var Router = gin.Default()

	// Router.StaticFS(global.GVA_CONFIG.Local.Path, http.Dir(global.GVA_CONFIG.Local.Path)) // 为用户头像和文件提供静态地址
	// Router.Use(middleware.LoadTls())  // 打开就能玩https了
	global.GVA_LOG.Info("use middleware logger")
	// 跨域
	Router.Use(middleware.Cors())
	global.GVA_LOG.Info("use middleware cors")
	// Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// global.GVA_LOG.Info("register swagger handler")
	// 方便统一添加路由组前缀 多服务器上线使用
	PublicGroup := Router.Group("")
	PublicGroup.Use(middleware.RateLimitByIP()).Use(middleware.RateLimitByTokenBucket())
	{
		router.InitBaseRouter(PublicGroup) // 登录注册获取留言
		router.InitUserBoard(PublicGroup)
	}
	PrivateGroup := Router.Group("")
	PrivateGroup.Use(middleware.RateLimitByIP()).Use(middleware.RateLimitByTokenBucket()).Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		router.InitApiRouter(PrivateGroup)  // 注册功能api路由
		router.InitJwtRouter(PrivateGroup)  // jwt相关路由
		router.InitUserRouter(PrivateGroup) // 注册用户路由
		router.InitMenuRouter(PrivateGroup) // 注册menu路由
		// router.InitEmailRouter(PrivateGroup)                 // 邮件相关路由
		router.InitSystemRouter(PrivateGroup) // system相关路由
		router.InitCasbinRouter(PrivateGroup) // 权限相关路由
		//router.InitCustomerRouter(PrivateGroup)  // 客户路由
		//router.InitAutoCodeRouter(PrivateGroup)  // 创建自动化代码
		router.InitAuthorityRouter(PrivateGroup) // 注册角色路由
		// router.InitSimpleUploaderRouter(PrivateGroup)        // 断点续传（插件版）
		// router.InitSysDictionaryRouter(PrivateGroup) // 字典管理
		// router.InitSysOperationRecordRouter(PrivateGroup)    // 操作记录
		// router.InitSysDictionaryDetailRouter(PrivateGroup)   // 字典详情管理
		router.InitFileUploadAndDownloadRouter(PrivateGroup) // 文件上传下载功能路由
		//router.InitWorkflowProcessRouter(PrivateGroup)       // 工作流相关接口

		router.InitClassRouter(PrivateGroup)     // 课程相关
		router.InitClassListRouter(PrivateGroup) // 课程固定信息相关，课程名-学时-上课地点

		router.InitGlobalRouter(PublicGroup) // 全局留言
	}
	global.GVA_LOG.Info("router register success")
	return Router
}
