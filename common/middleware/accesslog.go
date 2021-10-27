package middleware

import (
	"bytes"
	"fmt"

	"github.com/gin-gonic/gin"
)

type AccessLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w AccessLogWriter) Write(p []byte) (int, error) {
	if n, err := w.body.Write(p); err != nil {
		return n, err
	}
	return w.ResponseWriter.Write(p)
}

func AccessLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		bodyWriter := &AccessLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = bodyWriter

		// beginTime := time.Now().Unix()
		c.Next()
		// endTime := time.Now().Unix()
		fmt.Println(c.Request.PostForm.Encode())
		// fmt.Println(bodyWriter.body.String())

		// fields := logger.Fields{
		// 	"request":  c.Request.PostForm.Encode(),
		// 	"response": bodyWriter.body.String(),
		// }
		// global.Log.WithFields(fields).Infof("access log: method: %s, status_code: %d, begin_time: %d, end_time: %d",
		// 	c.Request.Method,
		// 	bodyWriter.Status(),
		// 	beginTime,
		// 	endTime,
		// )
	}
}
