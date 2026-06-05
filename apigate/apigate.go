package main

import (
	"apigate/initialize"
	"apigate/router"

	"go.uber.org/zap"
)

func main() {

	// 初始化日志记录器
	if err := initialize.Logger(); err != nil {
		panic("init logger err")
	}
	// 路由初始化
	r := router.Init()

	zap.L().Info("router start listen 127.0.0.1:8080")
	if err := r.Run("127.0.0.1:8080"); err != nil {
		zap.L().Error("router start listen err: %v", zap.Error(err))
		return
	}
}
