package main

import (
	"log"

	"github.com/cory-evans/location-tracker/internal/auth"
	"github.com/cory-evans/location-tracker/internal/devicelocation"
	"github.com/cory-evans/location-tracker/internal/mqttagent"
	"github.com/cory-evans/location-tracker/internal/util"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"
)

func main() {
	app := pocketbase.New()
	interalMQTTPassword, err := util.GenerateRandomString(256)
	if err != nil {
		panic(err)
	}

	agent := mqttagent.NewMQTTAgent(app, "mosquitto", 1883, interalMQTTPassword)

	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.AddRoute(devicelocation.CreateDeviceTokenRoute(app, "/api/device/:id/token"))
		e.Router.AddRoute(auth.CreateGetUserRoute(app, "/api/mqtt/user", interalMQTTPassword))
		e.Router.AddRoute(auth.CreateACLCheckRoute(app, "/api/mqtt/acl"))
		return nil
	})

	app.OnAfterBootstrap().Add(func(e *core.BootstrapEvent) error {
		agent.Start()
		return nil
	})

	migratecmd.MustRegister(app, app.RootCmd, &migratecmd.Options{
		Automigrate: false,
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
		agent.Quit()
	}
}
