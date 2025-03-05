package user

import (
	"context"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserServiceImpl struct {
	DB *gorm.DB
}

func (s *UserServiceImpl) Register(ctx context.Context, req *RegisterReq) (*RegisterResp, error) {
	// 密码加密
	hashedPwd, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	user := User{
		Email:    req.Email,
		Password: string(hashedPwd),
	}
	if err := s.DB.Create(&user).Error; err != nil {
		return nil, err
	}
	return &RegisterResp{UserId: int32(user.ID)}, nil
}

func (s *UserServiceImpl) Login(ctx context.Context, req *LoginReq) (*LoginResp, error) {
	var user User
	if err := s.DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
		return nil, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, err
	}
	return &LoginResp{UserId: int32(user.ID)}, nil
}
