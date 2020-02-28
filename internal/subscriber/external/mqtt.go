package external

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"io/ioutil"
	"os"
	"time"
)

func NewMqttClient() mqtt.Client {
	clientId := os.Getenv("CLIENT_ID") + ":" + time.Now().String()
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("%s:%s", os.Getenv("BROKER_URL"), os.Getenv("BROKER_PORT")))
	opts.SetClientID(clientId)
	opts.SetTLSConfig(NewTLSConfig())
	opts.SetAutoReconnect(true)
	opts.SetOnConnectHandler(func(client mqtt.Client) {
		o := client.OptionsReader()
		fmt.Printf("onConnect. %s\n", o.ClientID())
	})
	opts.SetConnectionLostHandler(func(client mqtt.Client, err error) {
		fmt.Printf("onConnectionLost. %s\n", err.Error())
	})

	c := mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	return c
}

func NewTLSConfig() *tls.Config {
	certPool := x509.NewCertPool()
	pemCerts, err := ioutil.ReadFile(os.Getenv("AWS_ROOT_CA_FILE_PATH"))
	if err == nil {
		certPool.AppendCertsFromPEM(pemCerts)
	}
	cert, err := tls.LoadX509KeyPair(os.Getenv("CERT_FILE_PATH"), os.Getenv("PRIVATE_KEY_PATH"))
	if err != nil {
		panic(err)
	}
	cert.Leaf, err = x509.ParseCertificate(cert.Certificate[0])
	if err != nil {
		panic(err)
	}
	return &tls.Config{
		RootCAs:            certPool,
		InsecureSkipVerify: true,
		Certificates:       []tls.Certificate{cert},
	}
}
