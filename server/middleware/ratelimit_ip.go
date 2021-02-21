package middleware

import (
	"gin-vue-admin/constant"
	"gin-vue-admin/global"
	"github.com/gin-gonic/gin"
	"time"
)

// 限制单IP单接口每秒请求速率
func RateLimitByIP() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP() + c.Request.URL.String()
		if in, _ := global.GVA_REDIS.Get("black-"+ip).Result(); in != "" {
			c.AbortWithStatusJSON(403, constant.TooMany)
			return
		}

		if cnt, _ := global.GVA_REDIS.Get(ip).Int(); cnt >= 10 {
			c.AbortWithStatusJSON(403, constant.TooMany)
			global.GVA_REDIS.Set("black-"+ip, "exist", 5*time.Second)
		} else if cnt == 0 {
			global.GVA_REDIS.Set(ip, 0, 10*time.Second)
		}

		global.GVA_REDIS.Incr(ip)
		c.Next()
	}
}
