package middleware

import (
	"github.com/gin-gonic/gin"
	"pep/constant"
	"pep/global"
	"pep/model/response"
	"pep/service"
	"time"
)

// 限制单IP单接口每秒请求速率
func RateLimitByIP() gin.HandlerFunc {
	return func(c *gin.Context) {
		// token 在黑名单里直接返回
		token := c.Request.Header.Get("x-token")
		if token == "" || service.IsBlacklist(token) {
			response.FailWithDetailed(gin.H{"reload": true}, "登录凭证失效或不存在", c)
			c.Abort()
			return
		}
		c.Set("token", token)

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
