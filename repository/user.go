package repository

import (
	"authentication-service/model"
	"context"

	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

type IUserRepo interface {
	UserByUserName(ctx context.Context, userName string) (*model.User, error)
}

func NewUserRepo(db *gorm.DB) IUserRepo {
	return &UserRepo{db: db}
}

func (r *UserRepo) UserByUserName(ctx context.Context, userName string) (*model.User, error) {
	user := &model.User{
		UserName: userName,
	}
	err := r.db.Model(&model.User{}).First(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
