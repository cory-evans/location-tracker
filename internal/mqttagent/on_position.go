package mqttagent

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/cory-evans/location-tracker/internal/dtos"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/models"
)

func CreateLocationPubHandler(app core.App) mqtt.MessageHandler {
	return func(c mqtt.Client, m mqtt.Message) {
		topicList := strings.Split(m.Topic(), "/")

		if len(topicList) != 2 {
			return
		}

		deviceId := topicList[1]

		var jsonBody dtos.Location

		err := json.Unmarshal(m.Payload(), &jsonBody)
		if err != nil {
			log.Println(err)
			return
		}

		coll, err := app.Dao().FindCollectionByNameOrId("locations")
		if err != nil {
			log.Println(err)
			return
		}

		r := models.NewRecord(coll)
		r.Set("device", deviceId)
		r.Set("lat", jsonBody.Lat)
		r.Set("lon", jsonBody.Lon)
		r.Set("acc", jsonBody.Acc)
		r.Set("speed", jsonBody.Speed)
		r.Set("timestamp", jsonBody.Timestamp.GetTimestampAsISO())

		err = app.Dao().SaveRecord(r)
		if err != nil {
			log.Println(err)
			return
		}
	}
}
