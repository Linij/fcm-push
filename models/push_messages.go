package models

import (
	"encoding/json"
	"time"
)

type PushMessage struct {
	ID        uint            `gorm:"primary_key"`
	Body      json.RawMessage `sql:"type:json"`
	DeviceId  string          `gorm:"not null;default:'';comment('设备 id');type:varchar(191);index;"`
	Status    string          `gorm:"type:varchar(32);default:'pending';comment('发送状态')"`
	CreatedAt time.Time
}
