package entity

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id         uuid.UUID `gorm:"type:uuid;"`
	Name       string
	Last       string
	CreateDate time.Time
	CreateBy   string
	UpdateDate time.Time
	UpdateBy   string
	IsActive   bool
}
