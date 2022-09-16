package main

import (
	"log"
	"net/http"

	"github.com/cory-evans/pocketbase-app/deviceauth"
	"github.com/cory-evans/pocketbase-app/devicelocation"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
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
					return err
				}

				id := c.PathParam("deviceId")

				device, err := app.Dao().FindRecordById(coll, id, nil)
				if err != nil {
					return err
				}

				token, err := deviceauth.NewDeviceToken(device.Id)
				if err != nil {
					return err
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
