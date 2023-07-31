package model

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	IsActive bool
	UserName string
	Password string
	Role     string
}
