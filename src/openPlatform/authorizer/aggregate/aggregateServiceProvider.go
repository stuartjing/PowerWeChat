package aggregate

import (
	"github.com/stuartjing/PowerWeChat/v3/src/kernel"
	"github.com/stuartjing/PowerWeChat/v3/src/openPlatform/authorizer/aggregate/account"
)

func RegisterProvider(app kernel.ApplicationInterface) (*account.Client, error) {

	//account, err := account.NewClient(app)
	//if err != nil {
	//	return nil, err
	//}
	//
	//return account, nil

	return nil, nil
}
