package auth

import (
	"errors"

	"github.com/pocketbase/pocketbase/core"
)

type DeviceAuthChecker struct {
	app core.App
}

func NewDeviceAuthChecker(app core.App) *DeviceAuthChecker {
	return &DeviceAuthChecker{
		app: app,
	}
}

func (c *DeviceAuthChecker) Valid(creds AuthCredentials) (bool, error) {
	record, err := c.app.Dao().FindFirstRecordByData("device_tokens", "device", creds.Username)
	if err != nil {
		return false, err
	}

	token := record.GetString("token")

	if creds.Password != token {
		return false, errors.New("incorrect password")
	}

	return true, nil
}
