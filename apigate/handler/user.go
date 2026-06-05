package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GetUserList(c *gin.Context) {
	zap.L().Info("get user list")
	c.JSON(http.StatusOK, gin.H{
		"msg": "这是用户列表",
	})
	return
}
