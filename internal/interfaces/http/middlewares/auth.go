package middlewares

import (
	"github.com/gin-gonic/gin"
	"log"
	"nekoacm-client/internal/interfaces/http/vo"
	"nekoacm-client/pkg/config"
	"net/http"
)

// Token验证中间件
func TokenAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		conf := config.Conf.Server
		token := c.GetHeader("Authorization")

		// 验证Token
		if token != "Bearer "+conf.Token {
			log.Println("token错误，拒绝请求")
			c.JSON(http.StatusUnauthorized, vo.RespError("token错误，拒绝请求", nil))
			c.Abort()
			return
		}

		// 放行
		c.Next()
	}
}
