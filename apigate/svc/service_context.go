package svc

import (
	"fmt"
	"os"
	"proto/pb"

	"github.com/go-redis/redis"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceContext struct {
	Redis   *redis.Client
	UserSrv pb.UserServiceClient
}

func NewServiceContext() *ServiceContext {
	// 一、 建立redis连接
	// 1. 加载.env文件
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	// 2. 从环境变量中加载配置
	ip := os.Getenv("RDB_HOST")
	port := os.Getenv("RDB_PORT")
	rdbCli := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", ip, port),
		Password: "",
		DB:       0,
	})

	// 3. 验证redis连接是否成功
	_, err := rdbCli.Ping().Result()
	if err != nil {
		panic("rdb connect fail")
	}

	// 二、 连接rpc服务
	userSrvconn, err := grpc.NewClient("127.0.0.1:8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		zap.L().Error("create new grpc client err: %v", zap.Error(err))
		return nil
	}

	return &ServiceContext{
		Redis:   rdbCli,
		UserSrv: pb.NewUserServiceClient(userSrvconn),
	}
}
