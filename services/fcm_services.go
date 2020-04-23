package services

import (
	"context"
	"errors"
	"firebase.google.com/go/messaging"
	"fmt"
	"gcm-push/tools"
)

type FcmService struct{}

// 向指定的设备发送消息
func (f *FcmService) SendToToken(token string, data map[string]string, notification map[string]string) (string, error) {
	notify := &messaging.Notification{
		Title:    notification["title"],
		Body:     notification["body"],
		ImageURL: notification["image_url"],
	}

	message := &messaging.Message{
		Data:         data,
		Token:        token,
		Notification: notify,
	}

	response, err := tools.Client.Send(context.Background(), message)

	if err != nil {
		return "Firebase 推送错误", err
	}

	return response, err
}

// 批量推送
func (f *FcmService) SendBatch(tokens []string, data map[string]string, notification map[string]string) ([]string, error) {
	notify := &messaging.Notification{
		Title:    notification["title"],
		Body:     notification["body"],
		ImageURL: notification["image_url"],
	}

	// 不能超过 100 个
	message := &messaging.MulticastMessage{
		Data:         data,
		Tokens:       tokens,
		Notification: notify,
	}

	br, sendErr := tools.Client.SendMulticast(context.Background(), message)

	if sendErr != nil {
		return nil, sendErr
	}

	var failedTokens []string
	var successTokens []string

	if br.FailureCount > 0 {
		for idx, resp := range br.Responses {
			if !resp.Success {
				failedTokens = append(failedTokens, tokens[idx])
			}
		}
		fmt.Printf("List of tokens that caused failures: %v\n", failedTokens)
	}

	if br.SuccessCount > 0 {
		for _, resp := range br.Responses {
			if resp.Success {
				successTokens = append(successTokens, resp.MessageID)
			}
		}
	}

	return successTokens, sendErr
}

// 向主题进行推送
func (f *FcmService) SendTopic(topic string, data map[string]string, notification map[string]string) (string, error) {
	notify := &messaging.Notification{
		Title:    notification["title"],
		Body:     notification["body"],
		ImageURL: notification["image_url"],
	}

	message := &messaging.Message{
		Data:         data,
		Topic:        topic,
		Notification: notify,
	}

	fmt.Printf("ptr 的值为 : %x\n", tools.FSApp)
	fmt.Printf("ptr 的值为 : %x\n", message)

	response, sendErr := tools.Client.Send(context.Background(), message)

	if sendErr != nil {
		fmt.Printf("send to topic error: %v\n", sendErr)
		return "发送错误", sendErr
	}

	fmt.Println("Successfully sent message:", response)

	return response, sendErr
}

// 客户端订阅主题
func (f *FcmService) SubscribeToTopic(tokens []string, topic string) (string, error) {

	// 订阅的客户端一次不能超过 1000 个
	if len(tokens) > 1000 {
		return "单次订阅客户端不能超过 1000 个", errors.New("单次订阅客户端不能超过 1000 个")
	}

	response, err := tools.Client.SubscribeToTopic(context.Background(), tokens, topic)
	if err != nil {
		return "订阅发生错误", err
	}

	fmt.Println(response.SuccessCount)

	return "订阅成功", err
}

// 客户端取消订阅
func (f *FcmService) UnsubscribeToTopic(tokens []string, topic string) (string, error) {

	// 订阅的客户端一次不能超过 1000 个
	if len(tokens) > 1000 {
		return "单次取消订阅客户端不能超过 1000 个", errors.New("单次取消订阅客户端不能超过 1000 个")
	}

	response, err := tools.Client.UnsubscribeFromTopic(context.Background(), tokens, topic)
	if err != nil {
		return "取消订阅发生错误", err
	}

	fmt.Println(response.SuccessCount)

	return "取消订阅成功", err
}
