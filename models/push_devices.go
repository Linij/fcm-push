package models

import (
	"time"
)

type PushDevice struct {
	ID           uint   `gorm:"primary_key"`
	Uuid         string `gorm:"not null;default:'';comment('uuid');type:varchar(128);index;"`
	DeviceId     string `gorm:"not null;default:'';comment('设备 id');type:varchar(191);index;"`
	Token        string `gorm:"not null;comment('推送 token');type:varchar(191);index;"`
	Platform     string `gorm:"type:enum('pc','mobile','android','ios');default:'android';comment('平台')"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	LastPushedAt *time.Time
}
