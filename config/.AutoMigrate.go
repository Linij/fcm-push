package main

/**
 *  僅用於初始化安裝 table 使用
 */

import (
    "fmt"
	"go-push/datasources"
	"go-push/models"
)

func main() {
	datasources.DB.DropTableIfExists(&models.PushDevice{}, &models.PushJob{}, &models.PushMessage{}, &models.PushTag{})
	fmt.Println("刪除原有表成功")
	datasources.DB.CreateTable(&models.PushDevice{}, &models.PushJob{}, &models.PushMessage{}, &models.PushTag{})
	fmt.Println("創建新表成功")
}
