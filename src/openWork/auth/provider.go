package auth

import (
	"github.com/stuartjing/PowerWeChat/v3/src/kernel"
)

func RegisterProvider(app kernel.ApplicationInterface) (*AccessToken, error) {

	accessToken, err := NewAccessToken(&app)

	return accessToken, err

}
