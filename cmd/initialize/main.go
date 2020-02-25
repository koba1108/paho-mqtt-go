package main

import (
	"cloud.google.com/go/bigquery"
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/koba1108/paho-mqtt-go/internal/subscriber/external"
	"github.com/koba1108/paho-mqtt-go/internal/subscriber/models"
	"os"
)

var tables = map[string]interface{}{
	models.TableNameBattery: models.Battery{},
	models.TableNameCharger: models.Charger{},
}

func main() {
	fmt.Printf("Start create tasks.\n")
	if os.Getenv("APP_ENV") != "production" {
		if err := godotenv.Load(".env"); err != nil {
			panic(err)
		}
	}
	ctx := context.Background()
	datasetId := os.Getenv("BQ_DATASET_ID")

	for tableId, model := range tables {
		createBqTableIfNotExists(ctx, datasetId, tableId, model)
	}
	fmt.Printf("End create tasks.\n")
}

func createBqTableIfNotExists(ctx context.Context, datasetId string, tableId string, model interface{}) {
	bqClient := external.GetBigquery(ctx)
	bqDataset := bqClient.Dataset(datasetId)
	table := bqDataset.Table(tableId)
	if _, err := table.Metadata(ctx); err != nil {
		fmt.Printf("%s table is not exists.\n", tableId)
		schema, err := bigquery.InferSchema(model)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Start table:%s create.\n", tableId)
		if err := table.Create(ctx, &bigquery.TableMetadata{Schema: schema}); err != nil {
			panic(err)
		}
		fmt.Printf("End table:%s create.\n", tableId)
	} else {
		fmt.Printf("%s table is already exists.\n", tableId)
	}
}
