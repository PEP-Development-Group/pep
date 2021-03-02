package initialize

import (
	"github.com/juju/ratelimit"
	"time"
)

func InitTokenBucket() *ratelimit.Bucket {
	// 瞬时最大可能2000人，系统每秒最多100人左右
	return ratelimit.NewBucketWithQuantum(time.Second, 10000, 500)
}
