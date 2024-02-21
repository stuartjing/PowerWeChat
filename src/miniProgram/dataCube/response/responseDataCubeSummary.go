package response

import (
	"github.com/stuartjing/PowerWeChat/v3/src/kernel/power"
	"github.com/stuartjing/PowerWeChat/v3/src/kernel/response"
)

type ResponseDataCubeSummary struct {
	response.ResponseMiniProgram

	List []*power.HashMap `json:"list"`
}
