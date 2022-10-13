package main

import (
	"log"

	"github.com/cory-evans/location-tracker/devicelocation"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)

func main() {
	app := pocketbase.New()

	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.AddRoute(devicelocation.CreateStoreLocationRoute(e.App, "/api/location"))

		return nil
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
