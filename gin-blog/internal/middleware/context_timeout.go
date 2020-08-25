package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"time"
)

func ContextTimeout(t time.Duration) func(c *gin.Context) {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), t) //通过 context.WithTimeout 设置当前超时时间
		defer cancel()

		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
