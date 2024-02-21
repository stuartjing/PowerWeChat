package response

import (
	"github.com/stuartjing/PowerWeChat/v3/src/kernel/power"
	"github.com/stuartjing/PowerWeChat/v3/src/kernel/response"
)

type ResponseBroadcastGetRoleList struct {
	response.ResponseMiniProgram

	Total int              `json:"total,omitempty"`
	List  []*power.HashMap `json:"list,omitempty"`
}
