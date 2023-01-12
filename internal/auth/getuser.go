package auth

import (
	"net/http"
	"strings"

	"github.com/cory-evans/location-tracker/internal/mqttagent"
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase/core"
)

func CreateGetUserRoute(app core.App, path string, internalMQTTPassword string) echo.Route {
	return echo.Route{
		Method: http.MethodPost,
		Path:   path,
		Handler: func(c echo.Context) error {
			var data AuthCredentials
			c.Bind(&data)

			var authChecker AuthChecker

			clientType := strings.Split(data.ClientId, "_")[0]
			switch clientType {
			case "device":
				authChecker = NewDeviceAuthChecker(app)
			case mqttagent.INTERNAL_USERNAME_PREFIX:
				authChecker = NewInternalAuthChecker(internalMQTTPassword)
			default:
				return echo.NewHTTPError(http.StatusUnauthorized)
			}

			isValid, err := authChecker.Valid(data)
			if err != nil {
				return err
			}

			if !isValid {
				return echo.NewHTTPError(http.StatusUnauthorized)
			}

			return nil
		},
	}
}
