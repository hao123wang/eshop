package router

import (
	"apigate/handler"
	"apigate/svc"

	"github.com/gin-gonic/gin"
)

func Init(svcCtx *svc.ServiceContext) *gin.Engine {
	r := gin.Default()

	api := r.Group("/v1")

	// 用户路由组
	userGroup := api.Group("/users")
	{
		userGroup.POST("/", handler.CreateUser(svcCtx))
	}

	return r
}
