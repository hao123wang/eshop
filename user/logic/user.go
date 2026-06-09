package logic

import (
	"context"
	"proto/pb"
	"time"
	"user-service/common/util"
	"user-service/model"
	"user-service/svc"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
		return nil, status.Error(codes.InvalidArgument, "请输入用户名")
	}
	if req.Mobile == "" {
		return nil, status.Error(codes.InvalidArgument, "请输入手机号")
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
		return nil, status.Error(codes.InvalidArgument, "手机号已被注册")
	}

	// 密码加密
	hashPwd, err := util.Encryption(req.Password)
	if err != nil {
		return nil, status.Error(codes.Internal, "服务器内部错误")
	}

	// 出生日期
	birthday, err := time.Parse("2006-01-02", req.Birthday)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "出生日期格式错误，正确格式为：yyyy-MM-dd")
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
		return nil, status.Error(codes.Internal, "用户创建失败")
	}

	return admin.ToProto(), nil
}

func (l *UserService) Login(req *pb.UserLogin) (*pb.UserInfo, error) {
	row := &model.User{
		Mobile: req.Mobile,
	}
	user, err := row.GetUser(l.svcCtx.DbConn)
	if err != nil {
		return nil, status.Error(codes.Internal, "服务器内部错误")
	}
	if user == nil {
		return nil, status.Error(codes.InvalidArgument, "用户名或密码错误")
	}

	if !util.IsCorrectPwd(user.Password, req.Password) {
		return nil, status.Error(codes.InvalidArgument, "用户名或密码错误")
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
		return nil, status.Error(codes.Internal, "服务器内部错误")
	}
	return user.ToProto(), nil
}

func (l *UserService) GetUserList(req *pb.UserListReq) (*pb.UserList, error) {
	out := &pb.UserList{}

	row := &model.User{}
	users, total, err := row.GetUserList(req, l.svcCtx.DbConn)
	if err != nil {
		return nil, status.Error(codes.Internal, "服务器内部错误")
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
		return nil, status.Error(codes.Internal, "服务器内部错误")
	}
	if user == nil {
		return nil, nil
	}

	// 出生日期
	birthday, err := time.Parse("2006-01-02", req.Birthday)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "出生日期格式不对，正确格式为：2006-01-02")
	}

	updateMap := make(map[string]any)
	updateMap["nick_name"] = req.NickName
	updateMap["birthday"] = birthday
	updateMap["gender"] = req.Gender

	if err := row.UpdateUser(updateMap, l.svcCtx.DbConn); err != nil {
		return nil, status.Error(codes.InvalidArgument, "internal server err")
	}

	user, err = row.GetUser(l.svcCtx.DbConn)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal server err")
	}
	return user.ToProto(), nil
}
