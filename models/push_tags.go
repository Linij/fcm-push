package models

import (
	"time"
)

type PushTag struct {
	ID        uint   `gorm:"primary_key"`
	DeviceId  string `gorm:"not null;default:'';comment('设备 id');type:varchar(191);unique_index;"`
	Tag       string `gorm:"type:varchar(64);default:'';comment('标签')"`
	CreatedAt time.Time
}
