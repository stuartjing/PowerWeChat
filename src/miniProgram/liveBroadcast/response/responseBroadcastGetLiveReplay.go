package response

import (
	"github.com/stuartjing/PowerWeChat/v3/src/kernel/power"
	"github.com/stuartjing/PowerWeChat/v3/src/kernel/response"
)

type ResponseBroadcastGetLiveReplay struct {
	response.ResponseMiniProgram

	Total      int              `json:"total"`
	LiveReplay []*power.HashMap `json:"live_replay"`
}
