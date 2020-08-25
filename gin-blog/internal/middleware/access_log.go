package middleware

import (
	"bytes"
	"gin-blog/global"
	"gin-blog/pkg/logger"
	"github.com/gin-gonic/gin"
	"time"
)

//针对访问日志的结构体
type AccessLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

//实现Write方法，解决无法直接获取方法响应主题的问题
func (w AccessLogWriter) Write(p []byte) (int, error) {
	if n, err := w.body.Write(p); err != nil {
		return n, err
	}
	return w.ResponseWriter.Write(p)
}

//通过函数的body取到值
func AccessLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		bodyWriter := &AccessLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = bodyWriter

		beginTime := time.Now().Unix()
		c.Next()
		endTime := time.Now().Unix()

		fields := logger.Fields{
			"request":  c.Request.PostForm.Encode(),
			"response": bodyWriter.body.String(),
		}
		s := "access log: method: %s, status_code: %d, " +
			"begin_time: %d, end_time: %d"
		//把 *gin.Context 传入日志方法中
		global.Logger.WithFields(fields).Infof(c, s,
			c.Request.Method,
			bodyWriter.Status(),
			beginTime,
			endTime,
		)
	}
}
