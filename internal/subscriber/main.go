package subscriber

import (
	Bigquery "cloud.google.com/go/bigquery"
	Firestore "cloud.google.com/go/firestore"
	"context"
	"encoding/json"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/koba1108/paho-mqtt-go/internal/subscriber/external"
	"github.com/koba1108/paho-mqtt-go/internal/subscriber/models/bigquery"
	"github.com/koba1108/paho-mqtt-go/internal/subscriber/models/firestore"
	"github.com/koba1108/paho-mqtt-go/internal/subscriber/models/pubsub"
	"os"
	"os/signal"
	"strings"
)

var (
	// firestore
	fsClient *Firestore.Client
	// bigquery
	bqClient        *Bigquery.Client
	bqDataset       *Bigquery.Dataset
	batteryInserter *Bigquery.Inserter
	chargerInserter *Bigquery.Inserter
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
	client := external.NewMqttClient()
	client.Subscribe(pubsub.TopicBattery, 1, onMessage)
	client.Subscribe(pubsub.TopicCharger, 1, onMessage)
	startSubscribe(ctx, client)
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
	batteryInserter = bqDataset.Table(bigquery.TableNameBattery).Inserter()
	chargerInserter = bqDataset.Table(bigquery.TableNameCharger).Inserter()
}

func startSubscribe(ctx context.Context, c mqtt.Client) {
	signal.Notify(signalCh, os.Interrupt)
	for {
		select {
		// 現在地の表示用
		case msg := <-batteryFsCh:
			var battery firestore.Battery
			if err := json.Unmarshal(msg.Payload(), &battery); err == nil {
				if _, err := battery.Update(ctx, fsClient, battery.TID); err != nil {
					// todo: error log
				}
			}
			fmt.Printf("battery %v", battery)
		case msg := <-chargerFsCh:
			var charger firestore.Charger
			if err := json.Unmarshal(msg.Payload(), &charger); err != nil {
				if _, err := charger.Update(ctx, fsClient, charger.CsID); err != nil {
					// todo: error log
				}
			}
			fmt.Printf("charger %v", charger)
		// ログ蓄積用
		case msg := <-batteryBqCh:
			var battery bigquery.Battery
			if err := json.Unmarshal(msg.Payload(), &battery); err != nil {
				if err := batteryInserter.Put(ctx, battery); err != nil {
					// todo: error log
				}
			}
			fmt.Printf("battery %v", battery)
		case msg := <-chargerBqCh:
			var charger bigquery.Charger
			if err := json.Unmarshal(msg.Payload(), &charger); err != nil {
				if err := chargerInserter.Put(ctx, charger); err != nil {
					// todo: error log
				}
			}
		// CLIで止めた時用
		case <-signalCh:
			fmt.Printf("Interrupt detected.\n")
			c.Disconnect(1000)
			return
		}
	}
}

func onMessage(_ mqtt.Client, msg mqtt.Message) {
	topicBase := strings.Split(msg.Topic(), "/")[0]
	switch topicBase {
	case pubsub.TopicBaseBattery:
		batteryFsCh <- msg
		batteryBqCh <- msg
	case pubsub.TopicBaseCharger:
		chargerFsCh <- msg
		chargerBqCh <- msg
	}
}
