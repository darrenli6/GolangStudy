package middleware

import (
	"github.com/darrenli6/blog-server/pkg/app"
	"github.com/darrenli6/blog-server/pkg/errcode"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			token  string
			encode = errcode.Success
		)

		if s, exist := c.GetQuery("token"); exist {
			token = s
		} else {
			token = c.GetHeader("token")
		}

		if token == "" {
			encode = errcode.InvalidParams
		} else {
			_, err := app.ParseToken(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					encode = errcode.UnauthorizedTokenTimeout
				default:
					encode = errcode.UnauthorizedTokenError
				}
			}
		}

		if encode != errcode.Success {
			response := app.NewResponse(c)
			response.ToErrorResponse(encode)
			c.Abort()
			return
		}

		c.Next()

	}
}
