package deviceauth

import (
	"errors"
	"strings"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/models"
	"github.com/pocketbase/pocketbase/tools/rest"
	"github.com/pocketbase/pocketbase/tools/security"
)

const (
	ContextDeviceKey string = "device"
)

func findDeviceByToken(app core.App, token string) (*models.Record, error) {
	claims, err := security.ParseUnverifiedJWT(token)
	if err != nil {
		return nil, err
	}

	// check required claims
	id, _ := claims["id"].(string)
	if id == "" {
		return nil, errors.New("missing or invalid token claims")
	}

	coll, err := app.Dao().FindCollectionByNameOrId("device")
	if err != nil {
		return nil, err
	}

	device, err := app.Dao().FindRecordById(coll, id, nil)
	if err != nil {
		return nil, err
	}

	if device == nil {
		return nil, errors.New("cannot find device")
	}

	return device, nil
}

func RequireDeviceAuth() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			device := c.Get(ContextDeviceKey)

			if device == nil {
				return rest.NewUnauthorizedError("The request requires valid device authorization token to be set.", nil)
			}

			return next(c)
		}
	}
}

func LoadDeviceContext(app core.App) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			token := c.Request().Header.Get("Authorization")

			if strings.HasPrefix(token, "Device ") {
				device, err := findDeviceByToken(app, token[7:])

				if err == nil && device != nil {
					c.Set(ContextDeviceKey, device)
				}
			}

			return next(c)
		}
	}
}
