package router

import (
	"apigate/handler"

	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	r := gin.Default()

	api := r.Group("/v1")

	// 用户路由组
	userGroup := api.Group("/users")
	{
		userGroup.GET("/", handler.GetUserList)
	}

	return r
}
