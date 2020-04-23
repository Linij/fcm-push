package main

import (
	"fmt"
	"gcm-push/datasources"
	"gcm-push/models"
)

func run() {
	datasources.DB.DropTableIfExists(&models.PushDevice{}, &models.PushJob{}, &models.PushMessage{}, &models.PushTag{})
	fmt.Println("刪除原有表成功")
	datasources.DB.CreateTable(&models.PushDevice{}, &models.PushJob{}, &models.PushMessage{}, &models.PushTag{})
	fmt.Println("創建新表成功")
}
