package devicelocation

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/cory-evans/location-tracker/deviceauth"
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
			device, _ := c.Get(deviceauth.ContextDeviceKey).(*models.Record)

			jsonBody := &Location{}
			err := json.NewDecoder(c.Request().Body).Decode(&jsonBody)
			if err != nil {
				return err
			}

			coll, err := app.Dao().FindCollectionByNameOrId("locations")
			if err != nil {
				return err
			}

			r := models.NewRecord(coll)
			r.SetDataValue("device", device.Id)
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
		Middlewares: []echo.MiddlewareFunc{
			deviceauth.RequireDeviceAuth(),
		},
	}
}
