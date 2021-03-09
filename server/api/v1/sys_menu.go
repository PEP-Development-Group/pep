package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"pep/global"
	"pep/model"
	"pep/model/request"
	"pep/model/response"
	"pep/service"
	"pep/utils"
)

var (
	stu = response.SysMenusResponse{Menus: []model.SysMenu{
		{MenuId: "1"}}}
	stuMenu = []model.SysBaseMenu{
		{GVA_MODEL: global.GVA_MODEL{ID: 1}, Hidden: false, ParentId: "0", Path: "dashboard", Name: "dashboard", Component: "view/dashboard/index.vue", Sort: 0, Meta: model.Meta{KeepAlive: false, DefaultMenu: false, Title: "首页", Icon: "house", CloseTab: false}},
		{GVA_MODEL: global.GVA_MODEL{ID: 2}, Hidden: true, ParentId: "0", Path: "about", Name: "about", Component: "view/about/index.vue", Sort: 7, Meta: model.Meta{KeepAlive: false, DefaultMenu: false, Title: "关于我们", Icon: "info", CloseTab: false}},
		{GVA_MODEL: global.GVA_MODEL{ID: 8}, Hidden: false, ParentId: "0", Path: "person", Name: "person", Component: "view/person/person.vue", Sort: 4, Meta: model.Meta{KeepAlive: false, DefaultMenu: false, Title: "个人信息", Icon: "message-solid", CloseTab: false}},
		{GVA_MODEL: global.GVA_MODEL{ID: 33}, Hidden: false, ParentId: "0", Path: "enroll", Name: "enroll", Component: "view/stu/enroll.vue", Sort: 3, Meta: model.Meta{KeepAlive: false, DefaultMenu: false, Title: "学生选课", Icon: "s-flag", CloseTab: false}},
		{GVA_MODEL: global.GVA_MODEL{ID: 35}, Hidden: false, ParentId: "0", Path: "lessons", Name: "lessons", Component: "view/stu/lessons.vue", Sort: 3, Meta: model.Meta{KeepAlive: false, DefaultMenu: false, Title: "选课查询", Icon: "success", CloseTab: false}},
	}
)

// 只缓存学生和老师
var menuMap map[string]*response.SysMenusResponse

func init() {
	menuMap = make(map[string]*response.SysMenusResponse, 2)
}

// @Tags AuthorityMenu
// @Summary 获取用户动态路由
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body request.Empty true "空"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /menu/getMenu [post]
func GetMenu(c *gin.Context) {
	auID := getUserAuthorityId(c)

	// 管理员均从数据库查询
	//if auID == "888" {
	//	_, menus := service.GetMenuTree(auID)
	//	response.OkWithDetailed(response.SysMenusResponse{Menus: menus}, "获取成功", c)
	//	return
	//}
	//
	//// 如果该角色权限缓存还没过期，直接返回
	//if service.CheckAuthorityTime(auID) {
	//	fmt.Println("该权限id已存在，直接返回:", *menuMap[auID])
	//	response.OkWithDetailed(menuMap[auID], "获取成功", c)
	//	return
	//}

	if menus, ok :=service.CheckUserAuthorityExist(auID);ok{
		fmt.Println("该权限id已存在，直接返回:", *menus)
		response.OkWithDetailed(response.SysMenusResponse{Menus: *menus}, "获取成功", c)
		return
	}

	// 否则从数据库查询最新数据，重置权限ID时间，再覆盖menuMap对应的权限
	if err, menus := service.GetMenuTree(auID); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		//service.CacheAuthorityID(auID)
		//menuMap[auID] = &response.SysMenusResponse{Menus: menus}
		service.CacheAuthorityMenu(auID, &menus)
		response.OkWithDetailed(response.SysMenusResponse{Menus: menus}, "获取成功", c)
	}
}

// @Tags AuthorityMenu
// @Summary 获取用户动态路由
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body request.Empty true "空"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /menu/getBaseMenuTree [post]
func GetBaseMenuTree(c *gin.Context) {
	//if menus, ok := service.CheckBase("baseMenus"); ok {
	//	response.OkWithDetailed(response.SysBaseMenusResponse{Menus: *menus}, "获取成功", c)
	//	return
	//}

	if err, menus := service.GetBaseMenuTree(); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.SysBaseMenusResponse{Menus: menus}, "获取成功", c)
		// service.CacheAuthorityID("baseMenus", &menus)
	}
}

// @Tags AuthorityMenu
// @Summary 增加menu和角色关联关系
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.AddMenuAuthorityInfo true "角色ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"添加成功"}"
// @Router /menu/addMenuAuthority [post]
func AddMenuAuthority(c *gin.Context) {
	var authorityMenu request.AddMenuAuthorityInfo
	_ = c.ShouldBindJSON(&authorityMenu)
	if err := utils.Verify(authorityMenu, utils.AuthorityIdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := service.AddMenuAuthority(authorityMenu.Menus, authorityMenu.AuthorityId); err != nil {
		global.GVA_LOG.Error("添加失败!", zap.Any("err", err))
		response.FailWithMessage("添加失败", c)
	} else {
		response.OkWithMessage("添加成功", c)
	}
	service.DelMenu(authorityMenu.AuthorityId)
}

// @Tags AuthorityMenu
// @Summary 获取指定角色menu
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetAuthorityId true "角色ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /menu/GetMenuAuthority [post]
func GetMenuAuthority(c *gin.Context) {
	var param request.GetAuthorityId
	_ = c.ShouldBindJSON(&param)
	if err := utils.Verify(param, utils.AuthorityIdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err, menus := service.GetMenuAuthority(&param); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithDetailed(response.SysMenusResponse{Menus: menus}, "获取失败", c)
	} else {
		response.OkWithDetailed(gin.H{"menus": menus}, "获取成功", c)
	}
}

// @Tags Menu
// @Summary 新增菜单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysBaseMenu true "路由path, 父菜单ID, 路由name, 对应前端文件路径, 排序标记"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"添加成功"}"
// @Router /menu/addBaseMenu [post]
func AddBaseMenu(c *gin.Context) {
	var menu model.SysBaseMenu
	_ = c.ShouldBindJSON(&menu)
	if err := utils.Verify(menu, utils.MenuVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := utils.Verify(menu.Meta, utils.MenuMetaVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := service.AddBaseMenu(menu); err != nil {
		global.GVA_LOG.Error("添加失败!", zap.Any("err", err))

		response.FailWithMessage("添加失败", c)
	} else {
		response.OkWithMessage("添加成功", c)
	}
}

// @Tags Menu
// @Summary 删除菜单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "菜单id"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /menu/deleteBaseMenu [post]
func DeleteBaseMenu(c *gin.Context) {
	var menu request.GetById
	_ = c.ShouldBindJSON(&menu)
	if err := utils.Verify(menu, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := service.DeleteBaseMenu(menu.Id); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Menu
// @Summary 更新菜单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysBaseMenu true "路由path, 父菜单ID, 路由name, 对应前端文件路径, 排序标记"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /menu/updateBaseMenu [post]
func UpdateBaseMenu(c *gin.Context) {
	var menu model.SysBaseMenu
	_ = c.ShouldBindJSON(&menu)
	if err := utils.Verify(menu, utils.MenuVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := utils.Verify(menu.Meta, utils.MenuMetaVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := service.UpdateBaseMenu(menu); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags Menu
// @Summary 根据id获取菜单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "菜单id"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /menu/getBaseMenuById [post]
func GetBaseMenuById(c *gin.Context) {
	var idInfo request.GetById
	_ = c.ShouldBindJSON(&idInfo)
	if err := utils.Verify(idInfo, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err, menu := service.GetBaseMenuById(idInfo.Id); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.SysBaseMenuResponse{Menu: menu}, "获取成功", c)
	}
}

// @Tags Menu
// @Summary 分页获取基础menu列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "页码, 每页大小"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /menu/getMenuList [post]
func GetMenuList(c *gin.Context) {
	var pageInfo request.PageInfo
	_ = c.ShouldBindJSON(&pageInfo)
	if err := utils.Verify(pageInfo, utils.PageInfoVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err, menuList, total := service.GetInfoList(); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     menuList,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
