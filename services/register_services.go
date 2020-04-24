package services

import (
	ds "go-push/datasources"
	"go-push/models"
	"time"
)

type RegisterService struct{}

// 註冊要推送的設備
func (r *RegisterService) Register(uuid string, deviceId string, token string, platform string) (err error) {
	deviceModel := &models.PushDevice{}
	ds.DB.Where("device_id = ? AND uuid = ?", deviceId, uuid).First(deviceModel)

	if deviceModel.ID == 0 {
		deviceModel.Uuid = uuid
		deviceModel.DeviceId = deviceId
		deviceModel.Token = token
		deviceModel.Platform = platform
		err = ds.DB.Create(deviceModel).Error
	} else {
		err = ds.DB.Model(deviceModel).Update(models.PushDevice{UpdatedAt: time.Now(), Uuid: uuid, Token: token, Platform: platform}).Error
	}

	return
}

// 刪除推送設備
func (r *RegisterService) Delete(deviceId string) (err error) {
	err = ds.DB.Delete(models.PushDevice{}, "device_id = ?", deviceId).Error

	return
}
