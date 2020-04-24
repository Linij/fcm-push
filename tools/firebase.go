package tools

import (
	"context"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"fmt"
	"google.golang.org/api/option"
	"sync"
)

var (
	FSApp      = NewFirebaseApp()
	Client     = NewFirebaseAppClient(FSApp)
	onceApp    sync.Once
	onceClient sync.Once
)

/**
 * 返回 *messaging.Client
 * @method New
 */
func NewFirebaseAppClient(app *firebase.App) *messaging.Client {
	var client *messaging.Client
	var err error

	onceClient.Do(func() {
		client, err = app.Messaging(context.Background())
	})

	if err != nil {
		fmt.Printf("Error getting Messaging client: %v\n", err)
		panic(err)
	}

	return client
}

/**
 * 返回 *firebase.App
 * @method New
 */
func NewFirebaseApp() *firebase.App {
	var app *firebase.App
	var err error

	onceApp.Do(func() {
		opt := option.WithCredentialsFile("./config/serviceAccountKey.json")
		app, err = firebase.NewApp(context.Background(), nil, opt)
	})

	if err != nil {
		fmt.Printf("Error getting Messaging client: %v\n", err)
		panic(err)
	}
	return app
}
