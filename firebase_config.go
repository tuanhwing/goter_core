package gotercore

import (
	"context"
	"os"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

//Initialize firebase app
func NewFirebaseApp() *firebase.App {
	opt := option.WithCredentialsFile(os.Getenv("FIREBASE_ADMIN_FILE"))
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		panic(err)
	}

	return app
}

//Initialize access auth firebase service
func NewFirebaseAuth(app *firebase.App) *auth.Client {
	auth, err := app.Auth(context.Background())
	if err != nil {
		panic(err)
	}
	return auth
}
