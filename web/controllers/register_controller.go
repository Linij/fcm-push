package controllers

import (
	"github.com/kataras/iris/v12"
	"go-push/services"
)

type RegisterController struct {
	Ctx             iris.Context
	RegisterService services.RegisterService
}

func (c *RegisterController) Get() {
	c.Ctx.JSON(ApiResource(true, nil, "註冊成功"))
}

// 注册推送的信息
func (c *RegisterController) Post() {
	var (
		uuid     = c.Ctx.FormValue("uuid")
		deviceId = c.Ctx.FormValue("device_id")
		token    = c.Ctx.FormValue("token")
		platform = c.Ctx.FormValue("platform")
	)

	err := c.RegisterService.Register(uuid, deviceId, token, platform)

	if err != nil {
		_, _ = c.Ctx.JSON(ApiResource(false, nil, "註冊失敗"))
		return
	}

	_, _ = c.Ctx.JSON(ApiResource(true, nil, "註冊成功"))
}

// 删除 device_id 的推送 token
func (c *RegisterController) DeleteByDeviceId() {
	var (
		deviceId = c.Ctx.FormValue("device_id")
	)

	err := c.RegisterService.DeleteByDevicedId(deviceId)

	if err != nil {
		_, _ = c.Ctx.JSON(ApiResource(false, nil, "刪除失敗"))
		return
	}

	_, _ = c.Ctx.JSON(ApiResource(true, nil, "刪除成功"))
}

// 根据 uuid 删除推送 token
func (c *RegisterController) DeleteByUUID() {
	var (
		uuid = c.Ctx.FormValue("uuid")
	)

	err := c.RegisterService.DeleteByUUID(uuid)

	if err != nil {
		_, _ = c.Ctx.JSON(ApiResource(false, nil, "刪除失敗"))
		return
	}

	_, _ = c.Ctx.JSON(ApiResource(true, nil, "刪除成功"))
}
