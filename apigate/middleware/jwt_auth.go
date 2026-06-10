package middleware

import (
	"apigate/utils/jwt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// JwtAuth jwt鉴权中间件，在客户端访问时验证token
func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "未认证",
			})
			c.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": 400,
				"msg":  "错误的认证信息",
			})
			c.Abort()
			return
		}

		claims, err := jwt.ParseToken(parts[1])
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "未认证",
			})
			c.Abort()
			return
		}

		zap.L().Sugar().Infof("id为 %d 的用户 %s 访问了本站", claims.UserID, claims.Username)
		c.Set("userID", claims.UserID)
		c.Set("username", claims.Username)
		c.Next()
	}
}
