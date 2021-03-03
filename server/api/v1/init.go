package v1

import (
	"pep/global"
	"pep/model"
	"pep/model/response"
	"github.com/gin-gonic/gin"
)

func InitDB(c *gin.Context) {
	err := global.GVA_DB.Unscoped().Where("authority_id = ?", 1).Delete(&model.SysUser{}).Error
	err = global.GVA_DB.Exec("DELETE FROM cls_class").Error
	err = global.GVA_DB.Exec("DELETE FROM user_classes").Error
	if err != nil {
		response.FailWithMessage("清空某表时失败，请手动删除剩余数据", c)
	} else {
		response.OkWithMessage("删除完毕", c)
	}
}
