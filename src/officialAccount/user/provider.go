package user

import (
	"github.com/stuartjing/PowerWeChat/v3/src/kernel"
	"github.com/stuartjing/PowerWeChat/v3/src/officialAccount/user/tag"
)

func RegisterProvider(app kernel.ApplicationInterface) (*Client, *tag.Client, error) {

	userClient, err := NewClient(app)
	if err != nil {
		return nil, nil, err
	}

	tagClient, err := tag.NewClient(app)
	if err != nil {
		return nil, nil, err
	}

	return userClient, tagClient, nil
}
