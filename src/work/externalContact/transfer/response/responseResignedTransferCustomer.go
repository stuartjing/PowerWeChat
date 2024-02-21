package response

import (
	"github.com/stuartjing/PowerWeChat/v3/src/kernel/power"
	"github.com/stuartjing/PowerWeChat/v3/src/kernel/response"
)

type ResponseResignedTransferCustomer struct {
	response.ResponseWork

	Customer []*power.HashMap `json:"customer"`
}
