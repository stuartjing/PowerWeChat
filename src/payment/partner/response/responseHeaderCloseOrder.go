package response

import (
	"github.com/stuartjing/PowerWeChat/v3/src/kernel/response"
)

type ResponseHeaderCloseOrdr struct {
	response.ResponsePayment

	Status string `header:"status"`
}
