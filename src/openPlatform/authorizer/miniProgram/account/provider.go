package account

import (
	"github.com/stuartjing/PowerWeChat/v3/src/kernel"
)

func RegisterProvider(app kernel.ApplicationInterface) (*Client, error) {

	account, err := NewClient(app)
	if err != nil {
		return nil, err
	}

	return account, nil
}
