package response

import (
	"github.com/stuartjing/PowerWeChat/v3/src/kernel/power"
	"github.com/stuartjing/PowerWeChat/v3/src/kernel/response"
)

type ResponseOperationGetJsErrSearch struct {
	response.ResponseMiniProgram
	results *power.HashMap `json:"results"`
	Total   int64          `json:"total"`
}
