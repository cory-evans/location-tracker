package devicelocation

import (
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/models"
)

type Location struct {
	Lat       float32 `json:"lat"`
	Lon       float32 `json:"lon"`
	Speed     float32 `json:"speed"`
	Acc       float32 `json:"acc"`
	Timestamp struct {
		Year   int `json:"year"`
		Month  int `json:"month"`
		Day    int `json:"day"`
		Hour   int `json:"hour"`
		Minute int `json:"minute"`
		Second int `json:"second"`
	} `json:"timestamp"`
}

func (l *Location) GetTimestampAsISO() string {
	return fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02dZ", l.Timestamp.Year, l.Timestamp.Month, l.Timestamp.Day, l.Timestamp.Hour, l.Timestamp.Minute, l.Timestamp.Second)
}

func CreateStoreLocationRoute(app core.App, path string) echo.Route {
	return echo.Route{
		Method: http.MethodPost,
		Path:   path,
		Handler: func(c echo.Context) error {
			tokenId, ok := c.Request().Header["Authorization"]
			if !ok || len(tokenId) == 0 {
				return errors.New("token")
			}

			tokenColl, err := app.Dao().FindCollectionByNameOrId("device_tokens")
			if err != nil {
				return err
			}

			tokenRecord, err := app.Dao().FindRecordById(tokenColl, tokenId[0], nil)
			if err != nil {
				return err
			}

			key := tokenRecord.GetStringDataValue("key")
			deviceId := tokenRecord.GetStringDataValue("device")

			keyAsBytes := make([]byte, 16)
			_, err = base64.StdEncoding.Decode(keyAsBytes, []byte(key))
			if err != nil {
				return err
			}

			jsonBody, err := DecodeRequestBody(keyAsBytes, c.Request().Body)
			if err != nil {
				return err
			}

			coll, err := app.Dao().FindCollectionByNameOrId("locations")
			if err != nil {
				return err
			}

			r := models.NewRecord(coll)
			r.SetDataValue("device", deviceId)
			r.SetDataValue("lat", jsonBody.Lat)
			r.SetDataValue("lon", jsonBody.Lon)
			r.SetDataValue("acc", jsonBody.Acc)
			r.SetDataValue("speed", jsonBody.Speed)
			r.SetDataValue("timestamp", jsonBody.GetTimestampAsISO())

			err = app.Dao().SaveRecord(r)
			if err != nil {
				return err
			}

			return c.JSON(http.StatusOK, r)
		},
	}
}
