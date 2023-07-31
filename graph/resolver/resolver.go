package resolver

import (
	"authentication-service/config"
	"authentication-service/service"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	authService service.IAuthService
}

func NewResolver(config *config.AppConfigs, logger *logrus.Logger, db *gorm.DB) *Resolver {
	return &Resolver{authService: service.NewAuthService(config.RedisConfig, db)}
}
