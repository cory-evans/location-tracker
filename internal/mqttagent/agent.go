package mqttagent

import (
	"fmt"
	"log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/pocketbase/pocketbase/core"
)

const INTERNAL_USERNAME_PREFIX = "internal"

type MQTTAgent struct {
	app    core.App
	client mqtt.Client
	token  mqtt.Token
}

func NewMQTTAgent(app core.App, broker string, port int, password string) MQTTAgent {

	opts := mqtt.NewClientOptions()
	opts.AutoReconnect = true
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", broker, port))

	opts.SetClientID(INTERNAL_USERNAME_PREFIX + "_api_service")
	opts.SetUsername(INTERNAL_USERNAME_PREFIX + "_api_service")
	opts.SetPassword(password)

	opts.SetOrderMatters(false)

	opts.SetDefaultPublishHandler(func(c mqtt.Client, m mqtt.Message) {
		log.Printf("unhandled mqtt: %d - %s", m.MessageID(), m.Topic())
	})

	opts.OnConnect = func(c mqtt.Client) {
		topic := "location/#"
		t := c.Subscribe(topic, 0, CreateLocationPubHandler(app))
		go func() {
			_ = t.Wait()
			if t.Error() != nil {
				fmt.Printf("ERROR SUBSCRIBING: %s\n", t.Error())
			} else {
				fmt.Println("subscribed to: ", topic)
			}
		}()
	}

	client := mqtt.NewClient(opts)

	return MQTTAgent{
		app:    app,
		client: client,
	}
}

func (a *MQTTAgent) Start() {
	a.token = a.client.Connect()
}

func (a *MQTTAgent) Quit() {
	a.client.Disconnect(250)
}
