package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"user-service/server"
	"user-service/svc"

	"proto/pb"

	"google.golang.org/grpc"
)

func main() {

	var ip string
	var port int

	flag.StringVar(&ip, "ip", "127.0.0.1", "server ip")
	flag.IntVar(&port, "port", 8080, "server port")
	flag.Parse()

	// 声明依赖
	c := svc.NewServiceContext()
	server := server.NewUserServiceServer(c)
	grpcSrv := grpc.NewServer()
	// 注册 grpc 服务
	pb.RegisterUserServiceServer(grpcSrv, server)

	addr := fmt.Sprintf("%s:%d", ip, port)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	go func(addr string) {
		fmt.Printf("服务已启动，监听地址：%s", addr)
		if err := grpcSrv.Serve(listener); err != nil {
			panic(err)
		}
	}(addr)

	// 优雅退出
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	fmt.Println("程序优雅退出")
}
