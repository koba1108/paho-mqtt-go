package external

import (
	"cloud.google.com/go/bigquery"
	"context"
	"os"
)

var client *bigquery.Client

func InitBigquery(ctx context.Context) {
	c, err := bigquery.NewClient(ctx, os.Getenv("GCP_PROJECT_ID"))
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
