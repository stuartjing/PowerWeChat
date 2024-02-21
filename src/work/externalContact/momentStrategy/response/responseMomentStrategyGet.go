package response

import (
	"github.com/stuartjing/PowerWeChat/v3/src/kernel/power"
	"github.com/stuartjing/PowerWeChat/v3/src/kernel/response"
)

type ResponseMomentStrategyGet struct {
	response.ResponseWork

	Strategy *power.HashMap `json:"strategy"`
}
