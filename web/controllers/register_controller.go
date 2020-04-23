package controllers

import (
	"gcm-push/services"
	"github.com/kataras/iris/v12"
)

type RegisterController struct {
	Ctx             iris.Context
	RegisterService services.RegisterService
}

func (c *RegisterController) Get() {
	c.Ctx.JSON(ApiResource(true, nil, "註冊成功"))
}

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

func (c *RegisterController) Delete() {
	var (
		deviceId = c.Ctx.FormValue("device_id")
	)

	err := c.RegisterService.Delete(deviceId)

	if err != nil {
		_, _ = c.Ctx.JSON(ApiResource(false, nil, "刪除失敗"))
		return
	}

	_, _ = c.Ctx.JSON(ApiResource(true, nil, "刪除成功"))
}
