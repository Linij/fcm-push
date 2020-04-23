package models

import (
	"encoding/json"
	"time"
)

type PushJob struct {
	ID           uint            `gorm:"primary_key"`
	Title        string          `gorm:"type:varchar(191);default:'';comment('标题')"`
	Body         json.RawMessage `sql:"type:json;comment('发送消息')"`
	PushObject   string          `gorm:"not null;default:'';comment('推送类型');type:varchar(191);"`
	PushPayload  string          `gorm:"not null;default:'';comment('推送tag或者推送device_id，默认是全部');type:varchar(191);"`
	ExcutePushAt time.Time
	CreatedAt    time.Time
}
