package response

import (
	"github.com/stuartjing/PowerWeChat/v3/src/kernel/power"
	"github.com/stuartjing/PowerWeChat/v3/src/kernel/response"
)

type ResponseMomentStrategyGetRange struct {
	response.ResponseWork

	Range      []*power.HashMap `json:"range"`
	NextCursor string           `json:"next_cursor"`
}
