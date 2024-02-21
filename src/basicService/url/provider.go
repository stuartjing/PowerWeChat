package url

import (
	"github.com/stuartjing/PowerWeChat/v3/src/kernel"
)

func RegisterProvider(app kernel.ApplicationInterface) (*Client, error) {

	return NewClient(&app)

}
