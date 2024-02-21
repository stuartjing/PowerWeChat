package response

import (
	"github.com/stuartjing/PowerWeChat/v3/src/kernel/power"
	response2 "github.com/stuartjing/PowerWeChat/v3/src/kernel/response"
)

type ResponseExpressAccountGetAll struct {
	*response2.ResponseMiniProgram

	Count int              `json:"count"`
	List  []*power.HashMap `json:"list"`
}
