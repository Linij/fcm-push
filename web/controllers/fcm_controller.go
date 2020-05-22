package controllers

import (
	"encoding/json"
	"github.com/kataras/iris/v12"
	"go-push/services"
	"strings"
)

type FcmController struct {
	Ctx        iris.Context
	FcmService services.FcmService
}

func (c *FcmController) Get() {

}

func (c *FcmController) PostSend() {
	var (
		token        = c.Ctx.FormValue("token")
		deviceId     = c.Ctx.FormValue("device_id")
		data         = c.Ctx.FormValue("data")
		notification = c.Ctx.FormValue("notification")
	)

	var sendData map[string]string
	var sendNotification map[string]string

	_ = json.Unmarshal([]byte(data), &sendData)
	_ = json.Unmarshal([]byte(notification), &sendNotification)

	result, err := c.FcmService.SendByToken(token, sendData, sendNotification, deviceId)

	if err != nil {
		_, _ = c.Ctx.JSON(ApiResource(false, nil, err.Error()))
		return
	}

	_, _ = c.Ctx.JSON(ApiResource(true, nil, result))
}

func (c *FcmController) PostSendUuid() {
	var (
		uuid         = c.Ctx.FormValue("uuid")
		data         = c.Ctx.FormValue("data")
		notification = c.Ctx.FormValue("notification")
	)

	var sendData map[string]string
	var sendNotification map[string]string

	_ = json.Unmarshal([]byte(data), &sendData)
	_ = json.Unmarshal([]byte(notification), &sendNotification)

	result, err := c.FcmService.SendByUuid(uuid, sendData, sendNotification)

	if err != nil {
		_, _ = c.Ctx.JSON(ApiResource(false, nil, err.Error()))
		return
	}

	_, _ = c.Ctx.JSON(ApiResource(true, nil, result))
}

func (c *FcmController) PostSendBatch() {
	var (
		token        = c.Ctx.FormValue("token")
		data         = c.Ctx.FormValue("data")
		notification = c.Ctx.FormValue("notification")
	)

	var tokens []string

	tokens = strings.Split(token, ",")

	var sendData map[string]string
	var sendNotification map[string]string

	_ = json.Unmarshal([]byte(data), &sendData)
	_ = json.Unmarshal([]byte(notification), &sendNotification)

	result, err := c.FcmService.SendBatch(tokens, sendData, sendNotification)

	if err != nil {
		_, _ = c.Ctx.JSON(ApiResource(false, nil, err.Error()))
		return
	}

	_, _ = c.Ctx.JSON(ApiResource(true, result, "ok"))
}

func (c *FcmController) PostSendTopic() {
	var (
		topic        = c.Ctx.FormValue("topic")
		data         = c.Ctx.FormValue("data")
		notification = c.Ctx.FormValue("notification")
	)

	var sendData map[string]string
	var sendNotification map[string]string

	_ = json.Unmarshal([]byte(data), &sendData)
	_ = json.Unmarshal([]byte(notification), &sendNotification)

	result, err := c.FcmService.SendTopic(topic, sendData, sendNotification)

	if err != nil {
		_, _ = c.Ctx.JSON(ApiResource(false, nil, err.Error()))
		return
	}
	_, _ = c.Ctx.JSON(ApiResource(true, nil, result))
}

func (c *FcmController) PostSubscribeToTopic() {
	var (
		topic = c.Ctx.FormValue("topic")
		token = c.Ctx.FormValue("token")
	)
	var tokens []string

	tokens = strings.Split(token, ",")

	_, err := c.FcmService.SubscribeToTopic(tokens, topic)

	if err != nil {
		_, _ = c.Ctx.JSON(ApiResource(false, nil, err.Error()))
		return
	}
	_, _ = c.Ctx.JSON(ApiResource(true, nil, "ok"))
}

func (c *FcmController) PostUnsubscribeToTopic() {
	var (
		topic = c.Ctx.FormValue("topic")
		token = c.Ctx.FormValue("token")
	)
	var tokens []string

	tokens = strings.Split(token, ",")

	_, err := c.FcmService.UnsubscribeToTopic(tokens, topic)

	if err != nil {
		_, _ = c.Ctx.JSON(ApiResource(false, nil, err.Error()))
		return
	}
	_, _ = c.Ctx.JSON(ApiResource(true, nil, "ok"))
}
