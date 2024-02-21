package virtualPayment

import (
	"github.com/stuartjing/PowerWeChat/v3/src/kernel"
)

func RegisterProvider(app kernel.ApplicationInterface) (*Client, error) {
	client, err := NewClient(&app)
	if err != nil {
		return nil, err
	}

	return client, nil

}
