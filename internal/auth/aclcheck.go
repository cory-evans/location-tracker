package auth

import (
	"net/http"
	"strings"

	"github.com/cory-evans/location-tracker/internal/mqttagent"
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase/core"
)

const (
	MOSQ_ACL_NONE      = 0x00
	MOSQ_ACL_READ      = 0x01
	MOSQ_ACL_WRITE     = 0x02
	MOSQ_ACL_SUBSCRIBE = 0x04
)

type aclparams struct {
	Username string `json:"username"`
	ClientId string `json:"clientid"`
	Topic    string `json:"topic"`
	Acc      int32  `json:"acc"`
}

func CreateACLCheckRoute(app core.App, path string) echo.Route {
	return echo.Route{
		Method: http.MethodPost,
		Path:   path,
		Handler: func(c echo.Context) error {
			var data aclparams
			c.Bind(&data)

			// if internal return no error
			if strings.Split(data.Username, "_")[0] == mqttagent.INTERNAL_USERNAME_PREFIX {
				return nil
			}

			topicList := strings.Split(data.Topic, "/")
			if len(topicList) == 2 {
				if topicList[0] == "location" && topicList[1] == data.Username && data.Acc == MOSQ_ACL_WRITE {
					// devices can write to their location
					return nil
				}
			}

			return echo.NewHTTPError(http.StatusUnauthorized)
		},
	}
}
