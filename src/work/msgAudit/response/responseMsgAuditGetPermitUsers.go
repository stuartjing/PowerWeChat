package response

import (
	"github.com/stuartjing/PowerWeChat/v3/src/kernel/response"
)

type ResponseMsgAuditGetPermitUsers struct {
	response.ResponseWork
	IDs []string `json:"ids"`
}
