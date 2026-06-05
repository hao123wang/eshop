package logic

import (
	"apigate/types"
)

type UserSrvLogic struct {
}

func NewUserSrvLogic() *UserSrvLogic {
	return &UserSrvLogic{}
}

func (l *UserSrvLogic) CreateUser(req types.UserInfo) (*types.UserInfoResp, error) {
	return nil, nil
}
