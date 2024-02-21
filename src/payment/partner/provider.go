package partner

import (
	"github.com/stuartjing/PowerWeChat/v3/src/payment/kernel"
)

func RegisterProvider(app kernel.ApplicationPaymentInterface) (*Client, error) {

	return NewClient(&app)

}
