package middleware

import (
	"github.com/gin-gonic/gin"
	"pep/constant"
	"pep/global"
	"time"
)

// 限制单IP单接口每秒请求速率
func RateLimitByIP() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP() + c.Request.URL.String()
		if in, _ := global.GVA_REDIS.Get("black-" + ip).Result(); in != "" {
			c.AbortWithStatusJSON(403, constant.TooMany)
			return
		}

		if cnt, _ := global.GVA_REDIS.Get(ip).Int(); cnt >= 15 {
			c.AbortWithStatusJSON(403, constant.TooMany)
			global.GVA_REDIS.Set("black-"+ip, "exist", 5*time.Minute)
		} else if cnt == 0 {
			global.GVA_REDIS.Set(ip, 0, 10*time.Second)
		}

		global.GVA_REDIS.Incr(ip)
		c.Next()
	}
}
