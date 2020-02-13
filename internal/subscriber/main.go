package subscriber

import (
	"context"
	"encoding/json"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/koba1108/paho-mqtt-go/internal/subscriber/external"
	"github.com/koba1108/paho-mqtt-go/internal/subscriber/models"
	"os"
	"os/signal"
	"strings"
)

var (
	msgCh     = make(chan mqtt.Message)
	signalCh  = make(chan os.Signal, 1)
	onMessage = func(_ mqtt.Client, msg mqtt.Message) { msgCh <- msg }
)

func Run() {
	ctx := context.Background()
	c := external.NewMqttClient()
	c.Subscribe(models.TopicBattery, 1, onMessage)
	c.Subscribe(models.TopicCharger, 1, onMessage)

	signal.Notify(signalCh, os.Interrupt)

	// modelに入れる？
	firebaseApp := external.GetFirebase(ctx)
	db, err := firebaseApp.Firestore(ctx)
	if err != nil {
		panic(err)
	}

	for {
		select {
		case msg := <-msgCh:
			topicBase := strings.Split(msg.Topic(), "/")[0]
			switch topicBase {
			case models.TopicBaseBattery:
				var battery models.Battery
				_ = json.Unmarshal(msg.Payload(), &battery)
				fmt.Printf("battery %v", battery)
				// todo: インドのタイムゾーン付ける
				if _, _, err := db.Collection(models.TopicBaseBattery).Add(ctx, battery); err != nil {
					// todo: add error log
				}
			case models.TopicBaseCharger:
				var charger models.Charger
				_ = json.Unmarshal(msg.Payload(), &charger)
				if _, _, err := db.Collection(models.TopicBaseCharger).Add(ctx, charger); err != nil {
					// todo: add error log
				}
			}
		case <-signalCh:
			fmt.Printf("Interrupt detected.\n")
			c.Disconnect(1000)
			return
		}
	}
}
