package initialize

import (
	"github.com/juju/ratelimit"
	"time"
)

func InitTokenBucket() {
	ratelimit.NewBucket(time.Second,0)

}