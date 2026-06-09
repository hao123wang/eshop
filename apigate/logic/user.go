package logic

import (
	"apigate/svc"
	"apigate/types"
	"context"
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

func (l *UserSrvLogic) CreateUser(req types.UserInfo) (*types.UserInfo, error) {
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
	httpResp := &types.UserInfo{
		UserID:   pbResp.UserId,
		NickName: pbResp.NickName,
		Mobile:   pbResp.Mobile,
		Email:    pbResp.Email,
		Birthday: pbResp.Birthday,
		Gender:   uint8(pbResp.Gender),
		Role:     uint8(pbResp.Role),
	}
	return httpResp, nil
}
