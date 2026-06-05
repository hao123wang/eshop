package svc

import (
	"proto/pb"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceContext struct {
	UserSrv pb.UserServiceClient
}

func NewServiceContext() *ServiceContext {
	svc := &ServiceContext{}

	conn, err := grpc.NewClient("127.0.0.1:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		zap.L().Error("create new grpc client err: %v", zap.Error(err))
		return nil
	}
	// 用户微服务 grpc 客户端
	svc.UserSrv = pb.NewUserServiceClient(conn)

	return svc
}
