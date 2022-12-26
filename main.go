package main

import (
	"log"

	"github.com/cory-evans/location-tracker/devicelocation"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"
)

func main() {
	app := pocketbase.New()

	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.AddRoute(devicelocation.CreateStoreLocationRoute(e.App, "/api/location"))
		e.Router.AddRoute(devicelocation.CreateDeviceTokenRoute(app, "/api/device/:id/token"))
		return nil
	})

	migratecmd.MustRegister(app, app.RootCmd, &migratecmd.Options{
		Automigrate: false,
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
