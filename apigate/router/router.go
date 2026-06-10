package router

import (
	"apigate/handler"
	"apigate/middleware"
	"apigate/svc"

	"github.com/gin-gonic/gin"
)

func Init(svcCtx *svc.ServiceContext) *gin.Engine {
	r := gin.Default()

	api := r.Group("/v1")

	// 用户登录
	api.POST("/login", handler.Login(svcCtx))

	// 用户路由组
	userGroup := api.Group("/users")
	userGroup.Use(middleware.JwtAuth())
	{
		userGroup.POST("/", handler.CreateUser(svcCtx))
		userGroup.POST("/getUser", handler.GetUserByID(svcCtx))
	}

	return r
}
