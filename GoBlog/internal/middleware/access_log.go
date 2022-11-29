package middleware

import (
	"bytes"
	"github.com/darrenli6/blog-server/global"
	"github.com/darrenli6/blog-server/pkg/logger"
	"github.com/gin-gonic/gin"
	"time"
)

type AccessLogWrite struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w AccessLogWrite) Write(p []byte) (int, error) {
	if n, err := w.body.Write(p); err != nil {
		return n, err
	}
	return w.ResponseWriter.Write(p)
}

func AccessLog() gin.HandlerFunc {
	return func(c *gin.Context) {

		bodyWrite := &AccessLogWrite{
			body:           bytes.NewBufferString(""),
			ResponseWriter: c.Writer,
		}
		c.Writer = bodyWrite

		beginTime := time.Now().Unix()
		c.Next()
		endTime := time.Now().Unix()

		fields := logger.Fields{
			"request":  c.Request.PostForm.Encode(),
			"response": bodyWrite.body.String(),
		}

		s := "access log : method %s status_code %d ," + "begine_time %d ,end_time: %d "
		global.Logger.WithFields(fields).Infof(c, s, c.Request.Method, bodyWrite.Status(), beginTime, endTime)
	}
}
