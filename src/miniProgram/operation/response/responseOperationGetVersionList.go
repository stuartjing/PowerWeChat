package response

import (
	"github.com/stuartjing/PowerWeChat/v3/src/kernel/power"
	"github.com/stuartjing/PowerWeChat/v3/src/kernel/response"
)

type ResponseOperationGetVersionList struct {
	response.ResponseMiniProgram
	CVList []*power.HashMap `json:"cvlist"`
}
