package response

import (
	"github.com/stuartjing/PowerWeChat/v3/src/kernel/power"
	"github.com/stuartjing/PowerWeChat/v3/src/kernel/response"
)

type ResponseOperationGetSceneList struct {
	response.ResponseMiniProgram
	Scene []*power.HashMap `json:"scene"`
}
