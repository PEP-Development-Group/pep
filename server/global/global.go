package global

import (
	"gin-vue-admin/config"
	"github.com/go-redis/redis"
	"github.com/juju/ratelimit"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	GVA_DB     *gorm.DB
	GVA_REDIS  *redis.Client
	GVA_CONFIG config.Server
	GVA_VP     *viper.Viper
	//GVA_LOG    *oplogging.Logger
	GVA_LOG *zap.Logger

	GVA_BUCKET *ratelimit.Bucket
)
