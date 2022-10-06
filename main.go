package main

import (
	"log"
	"net/http"

	"github.com/cory-evans/location-tracker/deviceauth"
	"github.com/cory-evans/location-tracker/devicelocation"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/models"
)

func main() {
	app := pocketbase.New()

	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.AddRoute(echo.Route{
			Method: http.MethodGet,
			Path:   "/api/device/:deviceId/token",
			Handler: func(c echo.Context) error {
				coll, err := app.Dao().FindCollectionByNameOrId("device")
				if err != nil {
					return echo.NewHTTPError(http.StatusBadRequest, err.Error())
				}

				id := c.PathParam("deviceId")

				device, err := app.Dao().FindRecordById(coll, id, nil)
				if err != nil {
					return echo.NewHTTPError(http.StatusBadRequest, err.Error())
				}

				deviceTokenColl, err := app.Dao().FindCollectionByNameOrId("device_tokens")
				if err != nil {
					return echo.NewHTTPError(http.StatusBadRequest, err.Error())
				}

				tokens, err := app.Dao().FindRecordsByExpr(deviceTokenColl, dbx.HashExp{"device": device.Id})
				if err != nil {
					return echo.NewHTTPError(http.StatusBadRequest, err.Error())
				}

				if len(tokens) > 0 {
					return echo.NewHTTPError(http.StatusBadRequest, "device token already exists")
				}

				tokenRecord := models.NewRecord(deviceTokenColl)
				tokenRecord.SetDataValue("device", device.Id)
				err = app.Dao().SaveRecord(tokenRecord)
				if err != nil {
					return echo.NewHTTPError(http.StatusBadRequest, err.Error())
				}

				token, err := deviceauth.NewDeviceToken(tokenRecord.Id, device.Id)
				if err != nil {
					return echo.NewHTTPError(http.StatusBadRequest, err.Error())
				}

				return c.String(200, token)
			},
			Middlewares: []echo.MiddlewareFunc{
				apis.RequireAdminOrUserAuth(),
			},
		})

		e.Router.Use(deviceauth.LoadDeviceContext(e.App))
		e.Router.AddRoute(devicelocation.CreateStoreLocationRoute(e.App, "/api/location"))

		return nil
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
