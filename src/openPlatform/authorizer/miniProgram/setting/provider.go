package setting

import "github.com/stuartjing/PowerWeChat/v3/src/kernel"

func RegisterProvider(app kernel.ApplicationInterface) (*Client, error) {

	code, err := NewClient(&app)
	if err != nil {
		return nil, err
	}

	return code, nil
}
