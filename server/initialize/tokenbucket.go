package initialize

import (
	"github.com/juju/ratelimit"
	"time"
)

func InitTokenBucket() *ratelimit.Bucket {
	// 瞬时最大可能1500人，每人每秒平均5r，系统每秒最多处理1000人，5000请求左右
	return ratelimit.NewBucketWithQuantum(time.Second, 15000, 5000)
}
