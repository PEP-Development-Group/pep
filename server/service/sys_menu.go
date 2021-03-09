package service

import (
	jsoniter "github.com/json-iterator/go"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"pep/global"
	"pep/model"
	"pep/model/request"
	"strconv"
	"time"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

//@author: [sh1luo](https://github.com/sh1luo)
//@function: getMenuTreeMap
//@description: 获取路由总树map
//@param: authorityId string
//@return: err error, treeMap map[string][]model.SysMenu

func getMenuTreeMap(authorityId string) (err error, treeMap map[string][]model.SysMenu) {
	var allMenus []model.SysMenu
	treeMap = make(map[string][]model.SysMenu)
	err = global.GVA_DB.Where("authority_id = ?", authorityId).Order("sort").Preload("Parameters").Find(&allMenus).Error

	//if res, _ := global.GVA_REDIS.Get(authorityId).Result(); res != "" {
	//	bs := []byte(res)
	//	_ = json.Unmarshal(bs, &allMenus)
	//}
	//
	//if len(allMenus) == 0 {
	//	err = global.GVA_DB.Where("authority_id = ?", authorityId).Order("sort").Preload("Parameters").Find(&allMenus).Error
	//	bs, _ := json.Marshal(&allMenus)
	//	// 缓存30天，中途如果有变动就删除掉重新缓存
	//	global.GVA_REDIS.Set(authorityId, bs, time.Hour*24*30)
	//}

	for _, v := range allMenus {
		treeMap[v.ParentId] = append(treeMap[v.ParentId], v)
	}
	return err, treeMap
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetMenuTree
//@description: 获取动态菜单树
//@param: authorityId string
//@return: err error, menus []model.SysMenu

func GetMenuTree(authorityId string) (err error, menus []model.SysMenu) {
	err, menuTree := getMenuTreeMap(authorityId)
	menus = menuTree["0"]
	for i := 0; i < len(menus); i++ {
		err = getChildrenList(&menus[i], menuTree)
	}
	return err, menus
}

// 检查权限ID是否过期
func CheckAuthorityTime(authorityId string) bool {
	duration, _ := global.GVA_REDIS.TTL(authorityId).Result()
	if duration.Seconds() > 0 {
		return true
	}
	return false
}

func CheckUserAuthorityExist(authorityId string) (*[]model.SysMenu, bool) {
	var menus = make([]model.SysMenu, 15)
	res, _ := global.GVA_REDIS.Get(authorityId).Result()
	if res != "" {
		_ = json.Unmarshal([]byte(res), &menus)
		return &menus, true
	}

	return nil, false
}

func CheckBase(baseMenus string) (*[]model.SysBaseMenu, bool) {
	var menus []model.SysBaseMenu
	if res, _ := global.GVA_REDIS.Get(baseMenus).Result(); res != "" {
		_ = json.Unmarshal([]byte(res), &menus)
		return &menus, true
	}

	return nil, false
}

func CacheAuthorityMenu(auID  string, menu *[]model.SysMenu) {
	bs, _ := json.Marshal(menu)
	_,err := global.GVA_REDIS.Set(auID, bs, time.Hour*1).Result()
	if err != nil {
		fmt.Println("缓存菜单数据失败:", auID)
	}
	fmt.Println("缓存菜单数据成功", auID)
}

func CacheAuthorityID(auID string) {
	// 随便设置一个值缓存权限id一小时，在这一个小时之内对应权限角色都不需要查询数据库
	// 会存在一个小时的时间差，即管理员更新了某个角色的权限，该角色最晚一个小时后才能获得最新菜单
	_, err := global.GVA_REDIS.Set(auID, 1, time.Hour*1).Result()
	if err != nil {
		fmt.Println("cache authorityID err:", err)
		return
	}
	fmt.Println("cache menu successful！id:", auID)
}

func DelMenu(auID string) {
	err := global.GVA_REDIS.Del(auID).Err()
	if err != nil {
		fmt.Println("del menus err:", err)
	}
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: getChildrenList
//@description: 获取子菜单
//@param: menu *model.SysMenu, treeMap map[string][]model.SysMenu
//@return: err error

func getChildrenList(menu *model.SysMenu, treeMap map[string][]model.SysMenu) (err error) {
	menu.Children = treeMap[menu.MenuId]
	for i := 0; i < len(menu.Children); i++ {
		err = getChildrenList(&menu.Children[i], treeMap)
	}
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetInfoList
//@description: 获取路由分页
//@return: err error, list interface{}, total int64

func GetInfoList() (err error, list interface{}, total int64) {
	var menuList []model.SysBaseMenu
	err, treeMap := getBaseMenuTreeMap()
	menuList = treeMap["0"]
	for i := 0; i < len(menuList); i++ {
		err = getBaseChildrenList(&menuList[i], treeMap)
	}
	return err, menuList, total
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: getBaseChildrenList
//@description: 获取菜单的子菜单
//@param: menu *model.SysBaseMenu, treeMap map[string][]model.SysBaseMenu
//@return: err error

func getBaseChildrenList(menu *model.SysBaseMenu, treeMap map[string][]model.SysBaseMenu) (err error) {
	menu.Children = treeMap[strconv.Itoa(int(menu.ID))]
	for i := 0; i < len(menu.Children); i++ {
		err = getBaseChildrenList(&menu.Children[i], treeMap)
	}
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: AddBaseMenu
//@description: 添加基础路由
//@param: menu model.SysBaseMenu
//@return: err error

func AddBaseMenu(menu model.SysBaseMenu) error {
	if !errors.Is(global.GVA_DB.Where("name = ?", menu.Name).First(&model.SysBaseMenu{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在重复name，请修改name")
	}
	return global.GVA_DB.Create(&menu).Error
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: getBaseMenuTreeMap
//@description: 获取路由总树map
//@return: err error, treeMap map[string][]model.SysBaseMenu

func getBaseMenuTreeMap() (err error, treeMap map[string][]model.SysBaseMenu) {
	var allMenus []model.SysBaseMenu
	treeMap = make(map[string][]model.SysBaseMenu)
	err = global.GVA_DB.Order("sort").Preload("Parameters").Find(&allMenus).Error
	for _, v := range allMenus {
		treeMap[v.ParentId] = append(treeMap[v.ParentId], v)
	}
	return err, treeMap
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetBaseMenuTree
//@description: 获取基础路由树
//@return: err error, menus []model.SysBaseMenu

func GetBaseMenuTree() (err error, menus []model.SysBaseMenu) {
	err, treeMap := getBaseMenuTreeMap()
	menus = treeMap["0"]
	for i := 0; i < len(menus); i++ {
		err = getBaseChildrenList(&menus[i], treeMap)
	}
	return err, menus
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: AddMenuAuthority
//@description: 为角色增加menu树
//@param: menus []model.SysBaseMenu, authorityId string
//@return: err error

func AddMenuAuthority(menus []model.SysBaseMenu, authorityId string) (err error) {
	var auth model.SysAuthority
	auth.AuthorityId = authorityId
	auth.SysBaseMenus = menus
	err = SetMenuAuthority(&auth)
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetMenuAuthority
//@description: 查看当前角色树
//@param: info *request.GetAuthorityId
//@return: err error, menus []model.SysMenu

func GetMenuAuthority(info *request.GetAuthorityId) (err error, menus []model.SysMenu) {
	err = global.GVA_DB.Where("authority_id = ? ", info.AuthorityId).Order("sort").Find(&menus).Error
	//sql := "SELECT authority_menu.keep_alive,authority_menu.default_menu,authority_menu.created_at,authority_menu.updated_at,authority_menu.deleted_at,authority_menu.menu_level,authority_menu.parent_id,authority_menu.path,authority_menu.`name`,authority_menu.hidden,authority_menu.component,authority_menu.title,authority_menu.icon,authority_menu.sort,authority_menu.menu_id,authority_menu.authority_id FROM authority_menu WHERE authority_menu.authority_id = ? ORDER BY authority_menu.sort ASC"
	//err = global.GVA_DB.Raw(sql, authorityId).Scan(&menus).Error
	return err, menus
}
