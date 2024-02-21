package response

import (
	"github.com/stuartjing/PowerWeChat/v3/src/kernel/power"
	"github.com/stuartjing/PowerWeChat/v3/src/kernel/response"
)

type ResponsePluginGetDevApplyList struct {
	response.ResponseMiniProgram
	ApplyList []*power.HashMap `json:"apply_list"`
}
