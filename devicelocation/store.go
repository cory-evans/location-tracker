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

			tokenRecord, err := app.Dao().FindRecordById("device_tokens", tokenId[0], nil)
			if err != nil {
				return err
			}

			key := tokenRecord.GetString("key")
			deviceId := tokenRecord.GetString("device")

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
			r.Set("device", deviceId)
			r.Set("lat", jsonBody.Lat)
			r.Set("lon", jsonBody.Lon)
			r.Set("acc", jsonBody.Acc)
			r.Set("speed", jsonBody.Speed)
			r.Set("timestamp", jsonBody.GetTimestampAsISO())

			err = app.Dao().SaveRecord(r)
			if err != nil {
				return err
			}

			return c.JSON(http.StatusOK, r)
		},
	}
}
