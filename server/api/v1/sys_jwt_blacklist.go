package v1

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"pep/global"
	"pep/model/response"
	"pep/service"
)

// @Tags Jwt
// @Summary jwt加入黑名单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"拉黑成功"}"
// @Router /jwt/jsonInBlacklist [post]
func JsonInBlacklist(c *gin.Context) {
	token := c.Request.Header.Get("x-token")
	if err := service.CreateJsonBlackListRecord(token); err != nil {
		global.GVA_LOG.Error("jwt作废失败!", zap.Any("err", err))
		response.FailWithMessage("注销失败", c)
	} else {
		response.OkWithMessage("注销成功", c)
	}
}
