package response

import (
	"github.com/stuartjing/PowerWeChat/v3/src/kernel/response"
)

type ResponseSubscribeMessageAddTemplate struct {
	response.ResponseMiniProgram
	PriTmplID string `json:"priTmplId"`
}
