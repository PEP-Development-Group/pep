package middleware

import (
	"github.com/gin-gonic/gin"
	"pep/constant"
	"pep/global"
	"pep/model/response"
)

// 全局令牌桶限流
func RateLimitByTokenBucket() gin.HandlerFunc {
	return func(c *gin.Context) {
		if remain := global.GVA_BUCKET.TakeAvailable(1); remain == 0 {
			response.FailWithMessage(constant.TooMany.Error(), c)
			c.Abort()
		}
		c.Next()
	}
}
