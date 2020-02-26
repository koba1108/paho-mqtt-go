package subscriber

import (
	"cloud.google.com/go/bigquery"
	"cloud.google.com/go/firestore"
	"context"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/koba1108/paho-mqtt-go/internal/subscriber/external"
	"github.com/koba1108/paho-mqtt-go/internal/subscriber/models"
	"os"
	"os/signal"
	"strings"
)

var (
	mqttClient mqtt.Client
	// firestore
	fsClient *firestore.Client
	// bigquery
	bqClient   *bigquery.Client
	bqDataset  *bigquery.Dataset
	batteryTbl *bigquery.Table
	chargerTbl *bigquery.Table
	// チャンネル
	signalCh    = make(chan os.Signal, 1)
	batteryFsCh = make(chan mqtt.Message)
	batteryBqCh = make(chan mqtt.Message)
	chargerFsCh = make(chan mqtt.Message)
	chargerBqCh = make(chan mqtt.Message)
)

func Run() {
	ctx := context.Background()
	initFirestore(ctx)
	initBigQuery(ctx)
	initMqtt(ctx)
	startSubscribe(ctx)
}

func initFirestore(ctx context.Context) {
	fs, err := external.GetFirestore(ctx)
	if err != nil {
		panic(err)
	}
	fsClient = fs
}

func initBigQuery(ctx context.Context) {
	bqClient = external.GetBigquery(ctx)
	bqDataset = bqClient.Dataset(os.Getenv("BQ_DATASET_ID"))
	batteryTbl = bqDataset.Table(models.TableNameBattery)
	chargerTbl = bqDataset.Table(models.TableNameCharger)
}

func initMqtt(ctx context.Context) {
	client := external.NewMqttClient()
	client.Subscribe(models.TopicBattery, 1, onMessage)
	client.Subscribe(models.TopicCharger, 1, onMessage)
	mqttClient = client
}

func startSubscribe(ctx context.Context) {
	signal.Notify(signalCh, os.Interrupt)
	defer mqttClient.Disconnect(1000)
	for {
		select {
		// 現在地の表示用
		case msg := <-batteryFsCh:
			data := models.NewBattery(msg.Payload())
			fmt.Printf("batteryFsCh: %v \n", data)
			docRef := fsClient.Collection(models.CollectionNameBattery).Doc(data.TID)
			if _, err := docRef.Set(ctx, data); err != nil {
				_ = fmt.Errorf("Error at <-chargerFsCh docRef.Set: %s. data: %v ", err.Error(), data)
			}
		case msg := <-chargerFsCh:
			data := models.NewCharger(msg.Payload())
			fmt.Printf("chargerFsCh: %v \n", data)
			docRef := fsClient.Collection(models.CollectionNameCharger).Doc(data.CsID)
			if _, err := docRef.Set(ctx, data); err != nil {
				_ = fmt.Errorf("Error at <-chargerFsCh docRef.Set: %s. data: %v ", err.Error(), data)
			}
		// ログ蓄積用
		case msg := <-batteryBqCh:
			data := models.NewBattery(msg.Payload())
			fmt.Printf("batteryBqCh: %v \n", data)
			if err := batteryTbl.Inserter().Put(ctx, data); err != nil {
				_ = fmt.Errorf("Error at <-batteryBqCh batteryTbl.Put: %s. data: %v ", err.Error(), data)
			}
		case msg := <-chargerBqCh:
			data := models.NewCharger(msg.Payload())
			fmt.Printf("chargerBqCh: %v \n", data)
			if err := chargerTbl.Inserter().Put(ctx, data); err != nil {
				_ = fmt.Errorf("Error at <-chargerBqCh chargerTbl.Put: %s. data: %v ", err.Error(), data)
			}
		// CLIで止めた時用
		case <-signalCh:
			fmt.Printf("Interrupt detected.\n")
			return
		}
	}
}

func onMessage(_ mqtt.Client, msg mqtt.Message) {
	topicBase := strings.Split(msg.Topic(), "/")[0]
	switch topicBase {
	case models.TopicBaseBattery:
		batteryFsCh <- msg
		batteryBqCh <- msg
	case models.TopicBaseCharger:
		chargerFsCh <- msg
		chargerBqCh <- msg
	}
}
