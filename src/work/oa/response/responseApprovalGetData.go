package response

import (
	"github.com/stuartjing/PowerWeChat/v3/src/kernel/power"
	"github.com/stuartjing/PowerWeChat/v3/src/kernel/response"
)

type ResponseApprovalGetData struct {
	response.ResponseWork
	Count     int            `json:"count"`
	Total     int            `json:"total"`
	NextSPNum int            `json:"next_spnum"`
	Data      *power.HashMap `json:"data"`
}
