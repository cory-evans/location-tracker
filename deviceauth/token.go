package deviceauth

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/pocketbase/pocketbase/tools/security"
)

func NewDeviceToken(tokenId, deviceId string) (string, error) {
	return security.NewToken(
		jwt.MapClaims{"id": tokenId, "sub": deviceId, "type": "device"},
		"device-secure-token",
		int64(time.Hour*24*7*52),
	)
}
