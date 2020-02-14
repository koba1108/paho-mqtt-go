package external

import (
	"cloud.google.com/go/firestore"
	"context"
	firebase "firebase.google.com/go"
	"fmt"
	"google.golang.org/api/option"
	"os"
)

var firebaseApp *firebase.App

func InitFirebase(ctx context.Context) {
	opt := option.WithCredentialsFile(os.Getenv("FIREBASE_SECRET_PATH"))
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		fmt.Println(err)
	}
	firebaseApp = app
}

func GetFirebase(ctx context.Context) *firebase.App {
	if firebaseApp == nil {
		InitFirebase(ctx)
	}
	return firebaseApp
}

func GetFirestore(ctx context.Context) (*firestore.Client, error) {
	if firebaseApp == nil {
		InitFirebase(ctx)
	}
	return firebaseApp.Firestore(ctx)
}
