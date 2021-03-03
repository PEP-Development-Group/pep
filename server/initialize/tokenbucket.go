package initialize

import (
	"github.com/juju/ratelimit"
	"time"
)

func InitTokenBucket() *ratelimit.Bucket {
	// 瞬时最大可能3000人，系统每秒最多1000人左右
	return ratelimit.NewBucketWithQuantum(time.Second, 15000, 5000)
}
