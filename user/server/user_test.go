package server

import (
	"context"
	"proto/pb"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Init() pb.UserServiceClient {
	conn, err := grpc.NewClient("127.0.0.1:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	client := pb.NewUserServiceClient(conn)
	return client
}

func TestCreatUser(t *testing.T) {
	cl := Init()
	user := &pb.UserInfo{
		NickName: "李四",
		Password: "admin",
		Mobile:   "12345678911",
		Birthday: "1999-12-13",
		Email:    "hhh@qq.com",
		Gender:   1,
		Role:     2,
	}
	resp, err := cl.CreateUser(context.Background(), user)
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(resp)
}

func TestLogin(t *testing.T) {
	req := &pb.UserLogin{
		Mobile:   "12345678911",
		Password: "admin",
	}
	cl := Init()
	resp, err := cl.Login(context.Background(), req)
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(resp)
}

func TestGetUser(t *testing.T) {
	req := &pb.UserSearch{
		UserId: 9,
	}
	cl := Init()
	resp, err := cl.GetUser(context.Background(), req)
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(resp)
}

func TestGetUserList(t *testing.T) {
	req := &pb.UserListReq{}
	cl := Init()
	resp, err := cl.GetUserList(context.Background(), req)
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(resp)
}

func TestUpdateUser(t *testing.T) {
	req := &pb.UserInfo{
		UserId:   1,
		NickName: "我是管理员",
		Mobile:   "11111111111",
		Birthday: "1998-02-13",
		Email:    "admin@qq.com",
		Role:     2,
	}
	cl := Init()
	resp, err := cl.UpdateUser(context.Background(), req)
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(resp)
}
