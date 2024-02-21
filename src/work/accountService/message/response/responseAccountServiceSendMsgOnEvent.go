package response

import (
	"github.com/stuartjing/PowerWeChat/v3/src/kernel/response"
)

type ResponseAccountServiceSendMsgOnEvent struct {
	response.ResponseWork

	MsgID string `json:"msgid"`
}
