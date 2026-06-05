package model

import (
	"errors"
	"proto/pb"
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	CreatedAt time.Time      `gorm:"created_at"`
	UpdatedAt time.Time      `gorm:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"deleted_at"`
}

type User struct {
	UserID   uint32    `gorm:"column:user_id;primary_key"`
	NickName string    `gorm:"column:nick_name"`
	Password string    `gorm:"column:password"`
	Mobile   string    `gorm:"column:mobile"`
	Email    string    `gorm:"column:email"`
	Birthday time.Time `gorm:"birthday"`
	Gender   uint8     `gorm:"gender"`
	Role     uint8     `gorm:"role"`
	BaseModel
}

func (*User) TableName() string {
	return "user"
}

func (row *User) CreateUser(db *gorm.DB) error {
	if db == nil {
		return errors.New("empty connection")
	}
	if err := db.Create(row).Error; err != nil {
		return err
	}
	return nil
}

func (row *User) GetUser(db *gorm.DB) (*User, error) {
	if db == nil {
		return nil, errors.New("empty connection")
	}
	db = db.Model(&User{})
	if row.UserID != 0 {
		db = db.Where("user_id = ?", row.UserID)
	}
	if row.Mobile != "" {
		db = db.Where("mobile = ?", row.Mobile)
	}
	var user User
	if err := db.First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (row *User) GetUserList(req *pb.UserListReq, db *gorm.DB) ([]*User, int64, error) {
	if db == nil {
		return nil, 0, errors.New("empty connection")
	}

	db = db.Model(&User{})

	var pageNo int
	var pageSize int
	if req.PageNo == 0 {
		pageNo = 1
	} else {
		pageNo = int(req.PageNo)
	}

	if req.PageSize == 0 {
		pageSize = 10
	} else {
		pageSize = int(req.PageSize)
	}

	var users []*User
	var total int64
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if err := db.Offset((pageNo - 1) * pageSize).Limit(pageSize).Find(&users).Error; err != nil {
		return nil, 0, err
	}
	return users, total, nil
}

func (row *User) UpdateUser(updateMap map[string]any, db *gorm.DB) error {
	if db == nil {
		return errors.New("empty connection")
	}
	if row.UserID == 0 {
		return nil
	}
	if err := db.Model(&User{}).Where("user_id = ?", row.UserID).Updates(updateMap).Error; err != nil {
		return err
	}
	return nil
}

func (row *User) ToProto() *pb.UserInfo {
	if row == nil {
		return nil
	}

	return &pb.UserInfo{
		UserId:   row.UserID,
		NickName: row.NickName,
		Mobile:   row.Mobile,
		Email:    row.Email,
		Gender:   uint32(row.Gender),
		Role:     uint32(row.Role),
		Birthday: row.Birthday.Format("2006-01-02"),
	}
}
