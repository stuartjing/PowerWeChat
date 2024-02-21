package response

import (
	"github.com/stuartjing/PowerWeChat/v3/src/kernel/power"
	"github.com/stuartjing/PowerWeChat/v3/src/kernel/response"
)

type ResponseServicerAdd struct {
	response.ResponseWork

	ResultList []*power.HashMap `json:"result_list"`
}
