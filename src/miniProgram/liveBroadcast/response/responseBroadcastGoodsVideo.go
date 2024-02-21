package response

import (
	"github.com/stuartjing/PowerWeChat/v3/src/kernel/response"
)

type ResponseBroadcastGoodsVideo struct {
	response.ResponseMiniProgram

	URL int `json:"url"`
}
