package logic

import (
	"context"
	"errors"
	"proto/pb"
	"time"
	"user-service/common/util"
	"user-service/model"
	"user-service/svc"
)

type UserService struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserService(ctx context.Context, svcCtx *svc.ServiceContext) *UserService {
	return &UserService{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserService) CreateUserLogic(req *pb.UserInfo) (*pb.UserInfo, error) {
	if req.NickName == "" {
		return nil, errors.New("请输入用户名")
	}
	if req.Mobile == "" {
		return nil, errors.New("请输入手机号")
	}

	// 手机号唯一校验
	row := &model.User{
		Mobile: req.Mobile,
	}
	user, err := row.GetUser(l.svcCtx.DbConn)
	if err != nil {
		return nil, err
	}
	if user != nil {
		return nil, errors.New("手机号已被注册")
	}

	// 密码加密
	hashPwd, err := util.Encryption(req.Password)
	if err != nil {
		return nil, errors.New("服务器故障")
	}

	// 出生日期
	birthday, err := time.Parse("2006-01-02", req.Birthday)
	if err != nil {
		return nil, errors.New("出生日期格式不对，正确格式为：2006-01-02")
	}
	// 创建用户
	admin := &model.User{
		NickName: req.NickName,
		Password: hashPwd,
		Mobile:   req.Mobile,
		Email:    req.Email,
		Birthday: birthday,
		Gender:   uint8(req.Gender),
		Role:     uint8(req.Role),
	}

	if err := admin.CreateUser(l.svcCtx.DbConn); err != nil {
		return nil, err
	}

	return admin.ToProto(), nil
}

func (l *UserService) Login(req *pb.UserLogin) (*pb.UserInfo, error) {
	row := &model.User{
		Mobile: req.Mobile,
	}
	user, err := row.GetUser(l.svcCtx.DbConn)
	if err != nil {
		return nil, errors.New("internal server error")
	}
	if user == nil {
		return nil, errors.New("用户未注册")
	}

	if !util.IsCorrectPwd(user.Password, req.Password) {
		return nil, errors.New("用户名或密码错误")
	}
	return user.ToProto(), nil
}

func (l *UserService) GetUser(req *pb.UserSearch) (*pb.UserInfo, error) {

	row := &model.User{
		UserID: req.UserId,
		Mobile: req.Mobile,
	}

	user, err := row.GetUser(l.svcCtx.DbConn)
	if err != nil {
		return nil, err
	}
	return user.ToProto(), nil
}

func (l *UserService) GetUserList(req *pb.UserListReq) (*pb.UserList, error) {
	out := &pb.UserList{}

	row := &model.User{}
	users, total, err := row.GetUserList(req, l.svcCtx.DbConn)
	if err != nil {
		return nil, err
	}

	var data []*pb.UserInfo
	for _, user := range users {
		data = append(data, user.ToProto())
	}
	out.Data = data
	out.Total = total
	return out, nil
}

func (l *UserService) UpdateUser(req *pb.UserInfo) (*pb.UserInfo, error) {
	// 查询用户
	row := &model.User{
		UserID: req.UserId,
	}
	user, err := row.GetUser(l.svcCtx.DbConn)
	if err != nil {
		return nil, errors.New("internal server err")
	}
	if user == nil {
		return nil, nil
	}

	// 出生日期
	birthday, err := time.Parse("2006-01-02", req.Birthday)
	if err != nil {
		return nil, errors.New("出生日期格式不对，正确格式为：2006-01-02")
	}

	updateMap := make(map[string]any)
	updateMap["nick_name"] = req.NickName
	updateMap["birthday"] = birthday
	updateMap["gender"] = req.Gender

	if err := row.UpdateUser(updateMap, l.svcCtx.DbConn); err != nil {
		return nil, errors.New("internal server err")
	}

	user, err = row.GetUser(l.svcCtx.DbConn)
	if err != nil {
		return nil, errors.New("internal server err")
	}
	return user.ToProto(), nil
}
