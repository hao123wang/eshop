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

	api.GET("/captcha", handler.Captcha(svcCtx)) // 获取验证码
	api.POST("/login", handler.Login(svcCtx))    // 用户登录

	// 用户路由组
	userGroup := api.Group("/users")
	userGroup.Use(middleware.JwtAuth())
	{
		userGroup.POST("/", handler.CreateUser(svcCtx))         // 创建用户
		userGroup.POST("/getUser", handler.GetUserByID(svcCtx)) // 获取用户信息
	}

	return r
}
