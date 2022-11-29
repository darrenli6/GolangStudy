package middleware

import (
	"fmt"
	"github.com/darrenli6/blog-server/global"
	"github.com/darrenli6/blog-server/pkg/app"
	"github.com/darrenli6/blog-server/pkg/email"

	"github.com/darrenli6/blog-server/pkg/errcode"
	"github.com/gin-gonic/gin"
	"time"
)

func Recovery() gin.HandlerFunc {
	defailMailer := email.NewEmail(&email.SMTPInfo{
		Host:     global.EmailSetting.Host,
		Port:     global.EmailSetting.Port,
		IsSSL:    global.EmailSetting.IsSSL,
		UserName: global.EmailSetting.UserName,
		Password: global.EmailSetting.Password,
		From:     global.EmailSetting.From,
	})

	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				s := "panic recover err %v"
				global.Logger.WithCallersFrames().Errorf(s, err)

				err := defailMailer.SendMail(
					global.EmailSetting.To,
					fmt.Sprintf("异常抛出 发生时间 %d", time.Now().Unix()),
					fmt.Sprintf("错误信息 ： %v", err),
				)
				if err != nil {
					global.Logger.Infof(c, "mail.sendmail err %v", err)
				}

				app.NewResponse(c).ToErrorResponse(errcode.ServerError)
				c.Abort()
			}
		}()

		c.Next()
	}
}
