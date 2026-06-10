package logic

import (
	"apigate/svc"
	"apigate/types"
	"apigate/utils/jwt"
	"context"
	"fmt"
	"proto/pb"
)

type UserSrvLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserSrvLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserSrvLogic {
	return &UserSrvLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserSrvLogic) CreateUser(req types.UserInfo) (*types.Resp, error) {
	// 将请求结构体转换为 pb 结构体
	pbReq := &pb.UserInfo{
		NickName: req.NickName,
		Password: req.Password,
		Email:    req.Email,
		Mobile:   req.Mobile,
		Birthday: req.Birthday,
		Gender:   uint32(req.Gender),
		Role:     uint32(req.Role),
	}

	// 调用微服务请求
	pbResp, err := l.svcCtx.UserSrv.CreateUser(l.ctx, pbReq)
	if err != nil {
		return nil, err
	}

	// 将 pb 响应结构体转换为 http 响应结构体
	data := &types.UserInfo{
		UserID:   pbResp.UserId,
		NickName: pbResp.NickName,
		Mobile:   pbResp.Mobile,
		Email:    pbResp.Email,
		Birthday: pbResp.Birthday,
		Gender:   uint8(pbResp.Gender),
		Role:     uint8(pbResp.Role),
	}
	resp := &types.Resp{
		Code: 200,
		Msg:  "success",
		Data: data,
	}
	return resp, nil
}

// Login 用户登录逻辑
func (l *UserSrvLogic) Login(req types.LoginReq) (*types.Resp, error) {
	// types -> pb
	pbReq := &pb.UserLogin{
		Mobile:   req.Mobile,
		Password: req.Pasword,
	}

	// 调用rpc
	pbResp, err := l.svcCtx.UserSrv.Login(l.ctx, pbReq)
	if err != nil {
		return nil, err
	}

	// 登录成功后，生成jwt响应给客户端
	tokenString, err := jwt.GenToken(pbResp.UserId, pbResp.NickName)
	if err != nil {
		return nil, fmt.Errorf("生成token失败: %v", err)
	}
	// pb -> types

	resp := &types.Resp{
		Code: 200,
		Msg:  "登录成功",
		Data: map[string]any{
			"token": tokenString,
		},
	}
	return resp, nil
}

// GetUserByID 根据id获取用户信息
func (l *UserSrvLogic) GetUserByID(req types.GetUserByID) (*types.Resp, error) {
	// types -> pb
	pbReq := &pb.UserSearch{
		UserId: req.UserID,
	}

	rpcResp, err := l.svcCtx.UserSrv.GetUser(l.ctx, pbReq)
	if err != nil {
		return nil, err
	}

	data := &types.UserInfo{
		UserID:   rpcResp.UserId,
		NickName: rpcResp.NickName,
		Mobile:   rpcResp.Mobile,
		Email:    rpcResp.Email,
		Birthday: rpcResp.Birthday,
		Gender:   uint8(rpcResp.Gender),
		Role:     uint8(rpcResp.Role),
	}

	resp := &types.Resp{
		Code: 200,
		Msg:  "success",
		Data: data,
	}
	return resp, nil
}
