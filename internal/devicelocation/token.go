package devicelocation

import (
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/forms"
	"github.com/pocketbase/pocketbase/models"

	"github.com/cory-evans/location-tracker/internal/util"
)

func CreateDeviceTokenRoute(app core.App, path string) echo.Route {
	return echo.Route{
		Method: http.MethodGet,
		Path:   path,
		Handler: func(c echo.Context) error {
			coll, err := app.Dao().FindCollectionByNameOrId("device_tokens")
			if err != nil {
				return err
			}

			deviceId := c.PathParam("id")

			authRecord, _ := c.Get(apis.ContextAuthRecordKey).(*models.Record)
			if authRecord == nil {
				return apis.NewForbiddenError("not authorized", nil)
			}

			device, err := app.Dao().FindRecordById("device", deviceId)
			if err != nil || device.GetString("owner") != authRecord.Id {
				return apis.NewForbiddenError("you do not own this device", nil)
			}

			record, err := app.Dao().FindFirstRecordByData("device_tokens", "device", deviceId)
			if err != nil {
				record = models.NewRecord(coll)
			}

			token, err := util.GenerateRandomStringURLSafe(64)
			if err != nil {
				return err
			}

			form := forms.NewRecordUpsert(app, record)
			form.LoadData(map[string]any{
				"device": deviceId,
				"token":  token,
			})

			if err := form.Submit(); err != nil {
				return err
			}

			return c.String(http.StatusOK, record.GetString("token"))
		},
		Middlewares: []echo.MiddlewareFunc{
			apis.ActivityLogger(app),
			apis.RequireRecordAuth(),
		},
	}
}
