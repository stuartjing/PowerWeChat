package response

import (
	"github.com/stuartjing/PowerWeChat/v3/src/kernel/power"
	"github.com/stuartjing/PowerWeChat/v3/src/kernel/response"
)

type ResponseCustomerStrategyGet struct {
	response.ResponseWork

	Strategy *power.HashMap `json:"momentStrategy"`
}
