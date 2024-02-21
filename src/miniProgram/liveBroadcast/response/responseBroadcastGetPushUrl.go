package response

import (
	"github.com/stuartjing/PowerWeChat/v3/src/kernel/response"
)

type ResponseBroadcastGetPushUrl struct {
	response.ResponseMiniProgram

	PushAddr string `json:"pushAddr"`
}
