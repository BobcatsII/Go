package middleware

import (
	"gin-blog/pkg/app"
	"gin-blog/pkg/errcode"
	"gin-blog/pkg/limiter"
	"github.com/gin-gonic/gin"
)

//将整体的限流器与对应的中间件逻辑串联起来
func RateLimiter(l limiter.LimiterIface) gin.HandlerFunc { //注意：入参是LimiterIface接口类型
	return func(c *gin.Context) {
		key := l.Key(c)
		if bucket, ok := l.GetBucket(key); ok {
			count := bucket.TakeAvailable(1) //TakeAvailable 会占用存储桶中立即可用的令牌数量，返回值为删除的令牌数；
			if count == 0 { //没有可用令牌，则返回0，说明已经超出配额
				response := app.NewResponse(c)
				response.ToErrorResponse(errcode.TooManyRequests) //返回0，则让客户减缓请求速度
				c.Abort()
				return
			}
		}

		c.Next()
	}
}
