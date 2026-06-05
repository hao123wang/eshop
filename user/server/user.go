// server 层， 通过声明结构体实现
package server

import (
	"context"
	"proto/pb"
	"user-service/logic"
	"user-service/svc"
)

type UserServiceServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedUserServiceServer
}

func NewUserServiceServer(svcCtx *svc.ServiceContext) *UserServiceServer {
	return &UserServiceServer{
		svcCtx: svcCtx,
	}
}

func (s *UserServiceServer) CreateUser(ctx context.Context, req *pb.UserInfo) (*pb.UserInfo, error) {
	l := logic.NewUserService(ctx, s.svcCtx)
	return l.CreateUserLogic(req)
}

func (s *UserServiceServer) Login(ctx context.Context, req *pb.UserLogin) (*pb.UserInfo, error) {
	l := logic.NewUserService(ctx, s.svcCtx)
	return l.Login(req)
}

func (s *UserServiceServer) GetUser(ctx context.Context, req *pb.UserSearch) (*pb.UserInfo, error) {
	l := logic.NewUserService(ctx, s.svcCtx)
	return l.GetUser(req)
}

func (s *UserServiceServer) GetUserList(ctx context.Context, req *pb.UserListReq) (*pb.UserList, error) {
	l := logic.NewUserService(ctx, s.svcCtx)
	return l.GetUserList(req)
}

func (s *UserServiceServer) UpdateUser(ctx context.Context, req *pb.UserInfo) (*pb.UserInfo, error) {
	l := logic.NewUserService(ctx, s.svcCtx)
	return l.UpdateUser(req)
}
