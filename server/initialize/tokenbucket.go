package initialize

import (
	"github.com/juju/ratelimit"
	"time"
)

func InitTokenBucket() *ratelimit.Bucket {
	// 瞬时最大可能500，系统每秒最多100
	return ratelimit.NewBucketWithQuantum(time.Second, 2500, 500)
}
