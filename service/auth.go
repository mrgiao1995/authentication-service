package service

import (
	"authentication-service/config"
	model_gen "authentication-service/graph/model"
	"authentication-service/model"
	"authentication-service/repository"
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type AuthService struct {
	repo  repository.IUserRepo
	redis repository.IRedisRepo
}

type IAuthService interface {
	Login(ctx context.Context, userName string, password string) (*model_gen.LoginResponse, error)
}

func NewAuthService(redisConfig *config.RedisConfig, db *gorm.DB) IAuthService {
	return &AuthService{redis: repository.NewRedisRepo(redisConfig), repo: repository.NewUserRepo(db)}
}

func (s *AuthService) Login(ctx context.Context, userName string, password string) (*model_gen.LoginResponse, error) {
	user, err := s.repo.UserByUserName(ctx, userName)

	if err != nil {
		return nil, err
	}

	if user.Password != password {
		return nil, errors.New("user name or password invalid")
	}
	accessToken, err := s.createSignedToken(ctx, 1, user, os.Getenv("JWT_SECRET"))
	if err != nil {
		return nil, err
	}

	refreshToken, err := s.createSignedToken(ctx, 4, user, os.Getenv("JWT_SECRET"))
	if err != nil {
		return nil, err
	}

	return &model_gen.LoginResponse{
		AccessToken:  *accessToken,
		RefreshToken: *refreshToken,
	}, nil
}

func (s *AuthService) createSignedToken(ctx context.Context, duration int, user *model.User, secret string) (*string, error) {
	now := time.Now()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &model.UserClaim{
		ID:       user.ID.String(),
		Role:     user.Role,
		UserName: user.UserName,
		IsActive: user.IsActive,
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        user.ID.String(),
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(time.Duration(duration) * time.Hour))},
	})

	signedToken, _ := token.SignedString([]byte(secret))
	err := s.redis.Add(ctx, fmt.Sprintf("%s:access", user.ID), signedToken, now.Add(time.Duration(duration)*time.Hour))
	if err != nil {
		return nil, err
	}
	return &signedToken, nil
}
