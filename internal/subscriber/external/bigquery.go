package external

import (
	"cloud.google.com/go/bigquery"
	"context"
	"google.golang.org/api/option"
	"os"
)

var client *bigquery.Client

func InitBigquery(ctx context.Context) {
	opt := option.WithCredentialsFile(os.Getenv("BIGQUERY_SECRET_PATH"))
	c, err := bigquery.NewClient(ctx, os.Getenv("GCP_PROJECT_ID"), opt)
	if err != nil {
		panic(err)
	}
	client = c
}

func GetBigquery(ctx context.Context) *bigquery.Client {
	if client == nil {
		InitBigquery(ctx)
	}
	return client
}
