package auth

import (
	"github.com/stuartjing/PowerWeChat/v3/src/kernel"
)

func RegisterProvider(app kernel.ApplicationInterface) (*AccessToken, error) {

	return NewAccessToken(&app)

}

func RegisterAuthProvider(app kernel.ApplicationInterface) (*Client, error) {

	return NewClient(&app)
}
