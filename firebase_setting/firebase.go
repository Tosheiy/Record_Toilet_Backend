package firebase_setting

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

var AuthClient *auth.Client

func Init() {
	ctx := context.Background()
	opt := option.WithCredentialsFile("./firebase_setting/fir-login-tutorial-9c6e2-firebase-adminsdk-wmhoj-c1fbb3fe31.json")
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v", err)
	}

	AuthClient, err = app.Auth(ctx)
	if err != nil {
		log.Fatalf("error getting Auth client: %v", err)
	}
}

func VerifyIDToken(idToken string) (*auth.Token, error) {
	// Firebase IDトークンの検証
	token, err := AuthClient.VerifyIDToken(context.Background(), idToken)
	if err != nil {
		// エラーが発生した場合、エラーを返す
		log.Printf("error verifying ID token: %v", err)
		return nil, err
	}

	return token, nil
}
