package devicelocation

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/cory-evans/pocketbase-app/deviceauth"
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/models"
)

func CreateStoreLocationRoute(app core.App, path string) echo.Route {
	return echo.Route{
		Method: http.MethodPost,
		Path:   path,
		Handler: func(c echo.Context) error {
			device, _ := c.Get(deviceauth.ContextDeviceKey).(*models.Record)

			jsonBody := make(map[string]interface{})
			err := json.NewDecoder(c.Request().Body).Decode(&jsonBody)
			if err != nil {
				return err
			}

			lat, latok := jsonBody["lat"]
			lon, lonok := jsonBody["lon"]
			if !latok || !lonok {
				return errors.New("cannot get lat/lon")
			}

			coll, err := app.Dao().FindCollectionByNameOrId("locations")
			if err != nil {
				return err
			}

			r := models.NewRecord(coll)
			r.SetDataValue("device", device.Id)
			r.SetDataValue("lat", lat)
			r.SetDataValue("lon", lon)

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
